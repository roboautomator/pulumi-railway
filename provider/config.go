package provider

import (
    "bytes"
    "encoding/json"
	"log"
    "net/http"
    "io/ioutil"
)

const RailwayAPIURL = "https://api.railway.app/graphql/v2"

type GraphQLRequest struct {
	Query         string      `json:"query"`
	Variables     interface{} `json:"variables"`
}

// Helper function to make GraphQL requests
func makeGraphQLRequest(query string, variables interface{}, apiToken string) string {
	// Prepare the GraphQL request
	requestBody := GraphQLRequest{
		Query:     query,
		Variables: variables,
	}

	log.Println("Making Request to Railway API with Query: %v", requestBody)

	// Convert request body to JSON
	reqBody, err := json.Marshal(requestBody)
	if err != nil {
		log.Println("Error marshalling request body: %v", err)
	}

	// Create the HTTP POST request
	req, err := http.NewRequest("POST", RailwayAPIURL, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Println("Error creating request: %v", req)
	}

	log.Println("Making Request to Railway API with Query: %v", query)

	// Set headers for the request
	req.Header.Set("Authorization", "Bearer " + apiToken)
	req.Header.Set("Content-Type", "application/json")

	// Execute the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making request: %v", err)
	}
	defer resp.Body.Close()

	// Read and return the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body: %v", err)
	}

	return string(body)
}

func getOrDefault[T comparable](value T, defaultValue T) T {
    var zeroValue T
    if value == zeroValue {
        return defaultValue
    }
    return value
}
