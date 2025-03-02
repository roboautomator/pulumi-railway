package provider

import (
    "context"
    "fmt"
)

type Service struct{}

type ServiceCreateInput struct {
	ProjectId string `json:"projectId"`
	EnvironmentId string `json:"environmentId"`
	Name string `json:"name"`
}

type ServiceDeleteInput struct {
	ServiceId string `json:"id"`
}

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
		},
	}

	serviceCreateResponse := makeGraphQLRequest(serviceCreateQuery, serviceCreateVariables, input.ApiToken)
	fmt.Println("Service Create Response:", serviceCreateResponse)

	state.Result = serviceCreateResponse

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