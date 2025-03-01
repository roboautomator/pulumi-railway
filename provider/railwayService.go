package provider

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
	"log"
    "net/http"
    "io/ioutil"
)

type Service struct{}

// type EnvironmentConfig struct {
// 	Services map[string]map[string]interface{} `json:"services"`
// }

// type ServiceCreateResponse struct {
// 	Data struct {
// 		ServiceCreate struct {
// 			ID   string `json:"id"`
// 			Name string `json:"name"`
// 		} `json:"serviceCreate"`
// 	} `json:"data"`
// }


type ServiceCreateInput struct {
	ProjectId string `json:"projectId"`
	EnvironmentId string `json:"environmentId"`
	Name string `json:"name"`
}

type ServiceArgs struct {
	EnvironmentId string `pulumi:"environmentId"`
	ProjectId string `pulumi:"projectId"`
	ApiToken string `pulumi:"apiToken"`
	Name string `pulumi:"name"`
}

type ServiceState struct {
	ServiceArgs
	ServiceId string `pulumi:"serviceId"`
	Result string `pulumi:"result"`
}

func (Service) Create(ctx context.Context, name string, input ServiceArgs, preview bool) (string, ServiceState, error) {
	
	state := ServiceState{ServiceArgs: input}

	serviceCreateQuery := `
	mutation serviceCreate($input: ServiceCreateInput!) {
		serviceCreate(input: $input) {
			id
		}
	}`
	serviceCreateVariables := map[string]interface{}{
		"input": ServiceCreateInput{
			ProjectId: input.ProjectId,
			EnvironmentId: input.EnvironmentId,
			Name: input.Name,
		},
	}

	serviceCreateResponse := makeGraphQLRequest(serviceCreateQuery, serviceCreateVariables, input.ApiToken)
	fmt.Println("Service Create Response:", serviceCreateResponse)

	state.Result = serviceCreateResponse

	return name, state, nil
}

func (Service) Delete(ctx context.Context, name string, input ServiceState) error {

	log.Printf("Deleting service with ID: %s", input.ServiceId)


	url := "https://api.railway.app/graphql/v2"
	payload := map[string]interface{}{
		"query": fmt.Sprintf(`
			mutation {
			    serviceDelete(id: "%s")
			}
		`, input.ServiceId),
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error creating request:", err)
		return err
	}

	req.Header.Set("Authorization", "Bearer " + input.ApiToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making request:", err)
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return err
	}

	fmt.Println(string(body))

	input.Result = "Service created"

	return nil
}