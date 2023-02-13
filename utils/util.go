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
