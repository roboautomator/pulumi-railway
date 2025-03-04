package provider

import (
    "context"
    "fmt"
	"encoding/json"
)

type Service struct{}

type ServiceCreateInput struct {
	ProjectId string `json:"projectId"`
	EnvironmentId string `json:"environmentId"`
	Name string `json:"name"`
	Source *ServiceSource `json:"source,omitempty"`
	Icon string `json:"icon,omitempty"`
	Variables map[string]string `json:"variables,omitempty"`
}

type ServiceSource struct {
	Image string `json:"image"`
	// Repo string `pulumi:"repo,optional"`
}

type ServiceArgs struct {
	EnvironmentId string `pulumi:"environmentId"`
	ProjectId string `pulumi:"projectId"`
	ApiToken string `pulumi:"apiToken"`
	Source *ServiceSource `pulumi:"source,optional"`
	Icon string `pulumi:"icon,optional"`
	Variables map[string]string `pulumi:"variables,optional"`
}

type ServiceState struct {
	ServiceArgs
	ServiceId string `pulumi:"serviceId"`
	Result string `pulumi:"result"`
}

type serviceCreateResponseData struct {
    Data struct {
        ServiceCreate struct {
            ID string `json:"id"`
        } `json:"serviceCreate"`
    } `json:"data"`
}

func (Service) Create(ctx context.Context, name string, input ServiceArgs, preview bool) (string, ServiceState, error) {
	
	state := ServiceState{ServiceArgs: input}

	if preview {
		return name, state, nil
	}

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
			Name: name,
			Source: input.Source,
			Icon: input.Icon,
			Variables: input.Variables,
		},
	}

	serviceCreateResponse := makeGraphQLRequest(serviceCreateQuery, serviceCreateVariables, input.ApiToken)
	fmt.Println("Service Create Response:", serviceCreateResponse)

    var response serviceCreateResponseData
    err := json.Unmarshal([]byte(serviceCreateResponse), &response)
    if err != nil {
        return "", state, err
    }

	state.ServiceId = response.Data.ServiceCreate.ID
	state.Result = serviceCreateResponse

	// // unlink service from source to prevent auto deploys
	// serviceUnlinkQuery := `
	// mutation serviceDisconnect($id: String!) {
	// 	serviceDisconnect(id: $id)
	// }`
	// serviceUnlinkVariables := map[string]interface{}{
	// 	"id": state.ServiceId,
	// }

	// serviceUnlinkResponse := makeGraphQLRequest(serviceUnlinkQuery, serviceUnlinkVariables, input.ApiToken)
	// fmt.Println("Service Unlink Response:", serviceUnlinkResponse)
	
	// state.Result = serviceUnlinkResponse

	return name, state, nil
}

func (Service) Delete(ctx context.Context, name string, input ServiceState) error {

	serviceDeleteQuery := `
	mutation serviceDelete($id: String!) {
		serviceDelete(id: $id)
	}`
	serviceDeleteVariables := map[string]interface{}{
		"id": input.ServiceId,
	}

	serviceDeleteResponse := makeGraphQLRequest(serviceDeleteQuery, serviceDeleteVariables, input.ApiToken)
	fmt.Println("Service Delete Response:", serviceDeleteResponse)

	return nil
}