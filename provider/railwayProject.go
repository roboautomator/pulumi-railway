package provider

import (
    "context"
	"encoding/json"
    "fmt"
)

type Project struct {}

type ProjectCreateInput struct {
	Name string `json:"name"`
}

type ProjectArgs struct {
	ApiToken string `pulumi:"apiToken"`
}

type ProjectState struct {
	ProjectArgs
	Result string `pulumi:"result"`
	ProjectId string `pulumi:"projectId"`
}

type projectCreateResponseData struct {
    Data struct {
        ProjectCreate struct {
            ID string `json:"id"`
        } `json:"projectCreate"`
    } `json:"data"`
}

func (Project) Create(ctx context.Context, name string, input ProjectArgs, preview bool) (string, ProjectState, error) {

	state := ProjectState{ProjectArgs: input}

	if preview {
		return name, state, nil
	}

	projectCreateQuery := `
	mutation projectCreate($input: ProjectCreateInput!) {
		projectCreate(input: $input) {
			id
		}
	}`
	projectCreateVariables := map[string]interface{}{
		"input": ProjectCreateInput{
			Name: name,
		},
	}

	projectCreateResponse := makeGraphQLRequest(projectCreateQuery, projectCreateVariables, input.ApiToken)
	fmt.Println("Project Create Response:", projectCreateResponse)

    var response projectCreateResponseData
    err := json.Unmarshal([]byte(projectCreateResponse), &response)
    if err != nil {
        return "", state, err
    }

    state.Result = projectCreateResponse
    state.ProjectId = response.Data.ProjectCreate.ID

	return name, state, nil
}

func (Project) Delete(ctx context.Context, name string, input ProjectState) error {

	projectDeleteQuery := `
	mutation projectDelete($id: String!) {
		projectDelete(id: $id)
	}`
	projectDeleteVariables := map[string]interface{}{
		"id": input.ProjectId,
	}

	projectDeleteResponse := makeGraphQLRequest(projectDeleteQuery, projectDeleteVariables, input.ApiToken)
	fmt.Println("Project Delete Response: %v", projectDeleteResponse)

	return nil
}