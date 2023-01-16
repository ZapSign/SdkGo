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
	return "ac3bfbdf-8588-480f-b7a9-c2f7e0f010f7dbd4edff-f7b5-4f42-b923-4daccc775a59"
}

func GetApiRoute() string {
	return "https://sandbox.api.zapsign.com.br/api/v1/"
}

func GetSignerToken() string {
	return "420222c4-143e-4f45-8e1f-8c99a1140c35"
}

func GetDocToken() string {
	return "token2"
}

// Use when you want to test a document token for deletion
func GetDocTokenThatWillBeDeleted() string {
	return "token976"
}

// Use when you want to test a signer token for deletion
func GetSignerTokenThatWillBeDeleted() string {
	return "464f0544-279a-48d4-96d5-e78a1394f571"
}

// Use when you have only one signer on the document and try to delete
func GetLastSignerTokenThatWillBeDeleted() string {
	return "91916c66-b85c-4a90-a862-6a8b2ccadf33"
}
