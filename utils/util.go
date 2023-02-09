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
	return "f452fca2-7c4c-47bd-99e4-ab4dcf28bb3590b0cf7b-47c7-416d-b545-aa7ce47ade44"
}

func GetApiRoute() string {
	return "https://api.zapsign.com.br/api/v1/"
}

func GetSignerToken() string {
	return "1b29bf51-9451-45be-98ac-8ff4474da75c"
}

func GetDocToken() string {
	return "8222b8d4-ed76-4af0-a8ae-0f056decbee9"
}

// Use when you want to test a document token for deletion
func GetDocTokenThatWillBeDeleted() string {
	return "65e8f427-a6e2-4901-848a-29b12deb3e10/"
}

// Use when you want to test a signer token for deletion
func GetSignerTokenThatWillBeDeleted() string {
	return "495c7250-609f-4a6c-b7c2-0cdd0685866f"
}

// Use when you have only one signer on the document and try to delete
func GetLastSignerTokenThatWillBeDeleted() string {
	return "91916c66-b85c-4a90-a862-6a8b2ccadf33"
}
