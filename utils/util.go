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
	return "15247ee5-9080-4c4a-81e1-6ebd0643fe54fa5156ab-7584-4df3-9b99-41c8e0ab1e28"
}

func GetApiRoute() string {
	return "http://localhost:3001/api/v1/"
}

func GetSignerToken() string {
	return "57cc1429-2064-4a6b-ac47-4cf95372c56a"
}

func GetDocToken() string {
	return "2d9263e1-abd2-4228-9e12-01677925610b"
}

// Use when you want to test a document token for deletion
func GetDocTokenThatWillBeDeleted() string {
	return "4e5130c7-f3d4-44cd-97fe-c4ae8480e156"
}

// Use when you want to test a signer token for deletion
func GetSignerTokenThatWillBeDeleted() string {
	return "5fd1e947-ff99-4e3c-928e-1a84dd4f301f"
}

func getBearerToken() string {
	return "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiZXhwIjoxNjczNjI1NTAzLCJqdGkiOiI2OGQwODQwODkyNmE0YjVlYTRjMjBjOWM5ZGE4ZGJiOSIsInVzZXJfaWQiOjF9.TriGJGxWC2YWOA2mT_MYkQ2Inn0iDYQdLVKhPlDf7g4"
}
