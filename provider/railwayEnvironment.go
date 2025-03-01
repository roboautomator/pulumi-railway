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

type Environment struct {}

type EnvironmentArgs struct {
	ApiToken string `pulumi:"apiToken"`
	ProjectId string `pulumi:"projectId"`
}


type EnvironmentState struct {
	EnvironmentArgs
	Result string `pulumi:"result"`
	EnvironmentId string `pulumi:"environmentId"`
}

func (Environment) Create(ctx context.Context, name string, input EnvironmentArgs, preview bool) (string, EnvironmentState, error) {
	state := EnvironmentState{EnvironmentArgs: input}
	if preview {
		return name, state, nil
	}

	url := "https://api.railway.app/graphql/v2"
	payload := map[string]interface{}{
		"query": fmt.Sprintf(`
			mutation {
				environmentCreate(input: { name: "%s", projectId: "%s" }) {
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

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", state, err
	}

	state.EnvironmentId = response["data"].(map[string]interface{})["environmentCreate"].(map[string]interface{})["id"].(string)
	state.Result = "Environment created successfully"

	return name, state, nil
}

func (Environment) Delete(ctx context.Context, name string, input EnvironmentState) error {
	url := "https://api.railway.app/graphql/v2"
	payload := map[string]interface{}{
		"query": fmt.Sprintf(`
			mutation {
				environmentDelete(input: { id: "%s" }) {
					success
				}
			}
		`, input.EnvironmentId),
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer " + input.ApiToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
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