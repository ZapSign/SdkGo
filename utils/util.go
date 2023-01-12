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
	return "862f2b2c-84d3-4162-a2ab-2970e0e68668e81b45c4-8205-4ad1-9fbe-cd23227ce460"
}

func GetApiRoute() string {
	return "http://localhost:3001/api/v1/"
}

func GetSignerToken() string {
	return "8ce2bcf4-731d-41e5-9c4c-7fe0bdba8a5b"
}

func GetDocToken() string {
	return "7ccc795e-3f1b-4d78-a93c-2bd1d97814fd"
}

// Use when you want to test a document token for deletion
func GetDocTokenThatWillBeDeleted() string {
	return "4e5130c7-f3d4-44cd-97fe-c4ae8480e156"
}

func getBearerToken() string {
	return "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiZXhwIjoxNjczNTc4MDEwLCJqdGkiOiJhNmFiNzU4OGMyYzc0NmFlYTIwMzhiOGRhZDBkYzI5OCIsInVzZXJfaWQiOjF9.umiPo2uvaYmh58lDeMd2iVQwLa-qrYr63MhYDEOSlD0"
}
