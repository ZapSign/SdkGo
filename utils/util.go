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
	const apiToken = "b1c75fcf-96d5-4951-b084-36870d2dde22d984198b-6186-4ac8-8135-bf3e586f0a34"
	return apiToken
}

func GetApiRoute() string {
	const apiRoute = "http://localhost:3001/api/v1/"
	return apiRoute
}

func GetSignerToken() string {
	const signerToken = "8ce2bcf4-731d-41e5-9c4c-7fe0bdba8a5b"
	return signerToken
}

func GetDocToken() string {
	const docToken = "dbbfae2b-617d-4443-b3cb-eff626bca499"
	return docToken
}