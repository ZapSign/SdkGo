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
	return "c5d61c24-18b5-4665-82e3-1683d18ef9b5b5ae6c12-9e45-4e97-ab5a-1c7cfd71c7d5"
}

func GetApiRoute() string {
	return "http://localhost:3001/api/v1/"
}

func GetSignerToken() string {
	return "8ce2bcf4-731d-41e5-9c4c-7fe0bdba8a5b"
}

func GetDocToken() string {
	return "5b5a3045-9f48-4042-a995-ecdcfb3326ab"
}

func getBearerToken() string {
	return "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiZXhwIjoxNjczNTQ0OTc4LCJqdGkiOiI0YjA3MGY0MmVjNTk0Mjg3YjQ5NDExOTRjODU4ODZkNiIsInVzZXJfaWQiOjF9.RavdinhpjSSSwnjuGHGP26cqNK33OuhYhzff893EZQ4"
}
