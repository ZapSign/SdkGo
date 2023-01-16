package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/utils"
	"github.com/stretchr/testify/assert"
)

func TestShouldSignInBatch(t *testing.T) {
	var apiSignBatchPath = utils.GetApiRoute() + "sign/"
	var getAPIToken = utils.GetApiToken()

	const signBatchMock = `"{

		"user_token": "98dad7ef-0271-4ed0-a0fe-605e2bd55eb3",
		"signer_tokens":[
			"7a319e9b-9579-47b0-aa67-4234706a07a3",
			"0a034962-80cd-4c4e-b27f-2d7e4e76ae5d"
		]
	}"`

	mockResponse := `{"message":"Documento(s) assinado(s) com sucesso. Lembrete: este endpoint é assíncrono, então aguarde os PDF finais ficarem prontos via webhooks ou confira-os daqui alguns minutos."}`
	responseRecorderStatusCode := httptest.NewRecorder().Code
	statusCode, body := controllers.PostRequest(signBatchMock, getAPIToken, apiSignBatchPath)

	assert.Equal(t, statusCode, responseRecorderStatusCode)
	assert.Equal(t, mockResponse, string(body))

}
