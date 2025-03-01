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

type Project struct {}

type ProjectArgs struct {
	ApiToken string `pulumi:"apiToken"`
}

type ProjectState struct {
	ProjectArgs
	Result string `pulumi:"result"`
	ProjectId string `pulumi:"projectId"`
}

func (Project) Create(ctx context.Context, name string, input ProjectArgs, preview bool) (string, ProjectState, error) {
	state := ProjectState{ProjectArgs: input}
	if preview {
		return name, state, nil
	}

	url := "https://api.railway.app/graphql/v2"
	payload := map[string]interface{}{
		"query": fmt.Sprintf(`
			mutation {
			    projectCreate(input: { name: "%s" }) {
					id
				}
			}
		`, name),
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

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", state, err
	}

	state.ProjectId = response["data"].(map[string]interface{})["projectCreate"].(map[string]interface{})["id"].(string)
	state.Result = "Project created"

	return name, state, nil
}

func (Project) Delete(ctx context.Context, name string, input ProjectState) error {

	log.Printf("Deleting project %s", input.ProjectId)

	url := "https://api.railway.app/graphql/v2"
	payload := map[string]interface{}{
		"query": fmt.Sprintf(`
			mutation {
			    projectDelete(id: "%s")
			}
		`, input.ProjectId),
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

	input.Result = "Project deleted"

	return nil
}