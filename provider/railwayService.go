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

type ServiceArgs struct {
	EnvironmentId string `pulumi:"environmentId"`
	ProjectId string `pulumi:"projectId"`
	ApiToken string `pulumi:"apiToken"`
}

type ServiceState struct {
	ServiceArgs
	ServiceId string `pulumi:"serviceId"`
	Result string `pulumi:"result"`
}

func (Service) Create(ctx context.Context, name string, input ServiceArgs, preview bool) (string, ServiceState, error) {
	state := ServiceState{ServiceArgs: input}
	if preview {
		return name, state, nil
	}

	url := "https://api.railway.app/graphql/v2"
	payload := map[string]interface{}{
		"query": fmt.Sprintf(`
			mutation {
			    serviceCreate(input: { name: "%s", projectId: "%s" }) {
                    id
            	}
			}
		`, name, input.ProjectId),
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", state, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", state, err
	}

	req.Header.Set("Authorization", "Bearer " + input.ApiToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", state, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", state, err
	}

	fmt.Println(string(body))

	// Extract the service ID from the response body
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", state, err
	}

	if data, ok := response["data"].(map[string]interface{}); ok {
		if serviceCreate, ok := data["serviceCreate"].(map[string]interface{}); ok {
			if id, ok := serviceCreate["id"].(string); ok {
				state.ServiceId = id
			} else {
				return "", state, fmt.Errorf("serviceCreate.id not found in response: %v", serviceCreate)
			}
		} else {
			return "", state, fmt.Errorf("serviceCreate missing in response: %v", data)
		}
	} else {
		return "", state, fmt.Errorf("Invalid API response: %v", response)
	}


	log.Printf("Stored ServiceId: %s", state.ServiceId)

	state.Result = "Service created"
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