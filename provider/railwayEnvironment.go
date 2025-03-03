package provider

import (
    "context"
    "encoding/json"
    "fmt"
)

type Environment struct {}

type EnvironmentArgs struct {
	ApiToken string `pulumi:"apiToken"`
	ProjectId string `pulumi:"projectId"`
	SkipInitialDeploys bool `pulumi:"skipInitialDeploys,optional"`
	// SourceEnvironmentId string `pulumi:"sourceEnvironmentId,optional"`
	StageInitialChanges bool `pulumi:"stageInitialChanges,optional"`
}

type EnvironmentCreateInput struct {
	Name string `json:"name"`
	ProjectId string `json:"projectId"`
	SkipInitialDeploys bool `json:"skipInitialDeploys"`
	// SourceEnvironmentId string `json:"sourceEnvironmentId"`
	StageInitialChanges bool `json:"stageInitialChanges"`
}

type EnvironmentState struct {
	EnvironmentArgs
	Result string `pulumi:"result"`
	EnvironmentId string `pulumi:"environmentId"`
}

type environmentCreateResponseData struct {
	Data struct {
		EnvironmentCreate struct {
			ID string `json:"id"`
		} `json:"environmentCreate"`
	} `json:"data"`
}

func (Environment) Create(ctx context.Context, name string, input EnvironmentArgs, preview bool) (string, EnvironmentState, error) {

	state := EnvironmentState{EnvironmentArgs: input}

	if preview {
		return name, state, nil
	}

	environmentCreateQuery := `
	mutation environmentCreate($input: EnvironmentCreateInput!) {
		environmentCreate(input: $input) {
			id
		}
	}`
	environmentCreateVariables := map[string]interface{}{
		"input": EnvironmentCreateInput{
			Name: name,
			ProjectId: input.ProjectId,
			SkipInitialDeploys: getOrDefault(input.SkipInitialDeploys, false),
			// SourceEnvironmentId: getOrDefault(input.SourceEnvironmentId, ""),
			StageInitialChanges: getOrDefault(input.StageInitialChanges, false),
		},
	}

	environmentCreateResponse := makeGraphQLRequest(environmentCreateQuery, environmentCreateVariables, input.ApiToken)
	fmt.Println("Environment Create Response:", environmentCreateResponse)

	var response environmentCreateResponseData
	err := json.Unmarshal([]byte(environmentCreateResponse), &response)
	if err != nil {
		return "", state, err
	}

	state.Result = environmentCreateResponse
	state.EnvironmentId = response.Data.EnvironmentCreate.ID

	return name, state, nil
}

func (Environment) Delete(ctx context.Context, name string, input EnvironmentState) error {

	projectDeleteQuery := `
	mutation environmentDelete($id: String!) {
			environmentDelete(id: $id)
	}`
	projectDeleteVariables := map[string]interface{}{
		"id": input.EnvironmentId,
	}

	environmentDeleteResponse := makeGraphQLRequest(projectDeleteQuery, projectDeleteVariables, input.ApiToken)
	fmt.Println("Environment Delete Response:", environmentDeleteResponse)

	return nil
}