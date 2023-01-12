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
	return "8c9e3294-8bde-4bdb-bea2-41da2fac0a2c"
}

// Use when you want to test a signer token for deletion
func GetSignerTokenThatWillBeDeleted() string {
	return "c5d60a4e-4ee6-4ba6-b934-7e78ce78d778"
}

func getBearerToken() string {
	return "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiZXhwIjoxNjczNjI4NDAzLCJqdGkiOiJiYzdhZDc1ZmJiNDM0YzAxOGFlNjA0ZmE4MGFiZjliOCIsInVzZXJfaWQiOjF9.XEBHyceqBTNgCC-kWiE4zG5ZNvsUq4RSdRpHkfOabr0"
}
