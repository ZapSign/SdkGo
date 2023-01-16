package utils

import (
	"encoding/json"
	"fmt"
)

func ConvertObjectToJSON(dockTemplate interface{}) string {
	jsonDoc, err := json.Marshal(dockTemplate)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	return string(jsonDoc)
}

func GetApiToken() string {
	return "f1fddb1d-5d5e-4747-8e50-8ceee841f2ee991b2861-0f6a-4601-97fa-63359b3bbcd2"
}

func GetApiRoute() string {
	return "http://localhost:3001/api/v1/"
}

func GetSignerToken() string {
	return "c5d60a4e-4ee6-4ba6-b934-7e78ce78d778"
}

func GetDocToken() string {
	return "11c380f7-4e0c-4283-b8ae-4ae6da0232f4"
}

// Use when you want to test a document token for deletion
func GetDocTokenThatWillBeDeleted() string {
	return "8c9e3294-8bde-4bdb-bea2-41da2fac0a2c"
}

// Use when you want to test a signer token for deletion
func GetSignerTokenThatWillBeDeleted() string {
	return "a4d5a4e5-2c08-4f42-9cfc-633fd8d390c5"
}

// Use when you have only one signer on the document and try to delete
func GetLastSignerTokenThatWillBeDeleted() string {
	return "c5d60a4e-4ee6-4ba6-b934-7e78ce78d778"
}

func getBearerToken() string {
	return "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiZXhwIjoxNjczNjM1MDE5LCJqdGkiOiI5NTBkNjRhMDhiOTk0ODAxOTBkOGI4MjJlMzcxNjllMCIsInVzZXJfaWQiOjF9.t2sGGyufcaDgRphgWb74ODIBC6nYKzHlzJzdTi-VsVE"
}
