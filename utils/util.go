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
	return "add8fb64-827c-49a5-a489-48959a5f9bbfece49a7b-02b0-4d82-b84b-e7f0c83aa46a"
}

func GetApiRoute() string {
	return "http://localhost:3001/api/v1/"
}

func GetSignerToken() string {
	return "8ce2bcf4-731d-41e5-9c4c-7fe0bdba8a5b"
}

func GetDocToken() string {
	return "3f51f0af-ddca-478c-a762-0c9d24148bbe"
}

func getBearerToken() string {
	return "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiZXhwIjoxNjczNDU1NzE0LCJqdGkiOiIwZjEyMmI3ZDZlZjg0N2QzYWUzNDE1MWUxM2UyOTM1YSIsInVzZXJfaWQiOjF9.xhUn0Ak0ZWJ-H5gWW_ZFMVUCjkcTZU0SrGkoJgunQ-U"
}
