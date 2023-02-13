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
<<<<<<< HEAD
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
=======
	return "{{INSIRA_SEU_API_TOKEN}}"
}

func GetApiRoute() string {
	return "https://api.zapsign.com.br/api/v1/"
}

func GetSignerToken() string {
	return "{{INSIRA_TOKEN_DE_UM_SIGNATÁRIO_QUE_SERA_DELETADO}}"
}

func GetDocToken() string {
	return "{{INSIRA_TOKEN_DE_UM_DOCUMENTO}}"
}

// Use when you want to test a document token for deletion
func GetDocTokenThatWillBeDeleted() string {
	return "{{INSIRA_TOKEN_DE_UM_DOCUMENTO_QUE_SERA_DELETADO}}"
}

// Use when you want to test a signer token for deletion
func GetSignerTokenThatWillBeDeleted() string {
	return "{{INSIRA_TOKEN_DE_UM_SIGNATARIO_QUE_SERA_DELETADO}}"
}
>>>>>>> homolog
