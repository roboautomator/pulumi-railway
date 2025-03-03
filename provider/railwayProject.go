package provider

import (
    "context"
	"encoding/json"
    "fmt"
	"log"
)

type Project struct {}

type Repo struct {
	Branch string `pulumi:"branch"`
	FullRepoName string `pulumi:"fullRepoName"`
}

type Runtime string
const (
	RuntimeLegacy Runtime = "LEGACY"
	RuntimeUnspecified Runtime = "UNSPECIFIED"
	RuntimeV2 Runtime = "V2"
)

type ProjectArgs struct {
	ApiToken 				string `pulumi:"apiToken"`
	DefaultEnvironmentName 	string `pulumi:"defaultEnvironmentName,optional"`
	Description 			string `pulumi:"description,optional"`
	IsPublic 				bool `pulumi:"isPublic,optional"`
	PrDeploys 				bool `pulumi:"prDeploys,optional"`
	Runtime 				Runtime `pulumi:"runtime,optional"`
	// Plugins 				*string `pulumi:"plugins,optional"`
	// TeamId 				string `pulumi:"teamId,optional"`
	// Repo 				*Repo `pulumi:"repo,optional"`
}

type ProjectCreateInput struct {
	Name 					string `json:"name"`
	DefaultEnvironmentName 	string `json:"defaultEnvironmentName"`
	Description 			string `json:"description"`
	IsPublic 				bool `json:"isPublic"`
	PrDeploys 				bool `json:"prDeploys"`
	Runtime 				Runtime `json:"runtime"`
	// Plugins 				*string `pulumi:"plugins,optional"`
	// TeamId 				string `json:"teamId"`
	// Repo 				*Repo `pulumi:"repo,optional"`
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
			Name:                   name,
			DefaultEnvironmentName: getOrDefault(input.DefaultEnvironmentName, "production"),
			Description:            getOrDefault(input.Description, "Pulumi Generated Railway Project"),
			IsPublic:               getOrDefault(input.IsPublic, false),
			PrDeploys:              getOrDefault(input.PrDeploys, false),
			Runtime:                getOrDefault(input.Runtime, RuntimeV2),
			// TeamId:                 getOrDefault(input.TeamId, "default-team-id").(string),
			// Plugins:                getOrDefault(input.Plugins, ""),
			// Repo:                   getOrDefault(input.Repo, &Repo{}).(Repo),
		},
	}

	log.Println("Project Create Variables:", projectCreateVariables)
	

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