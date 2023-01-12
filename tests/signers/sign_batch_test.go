package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/utils"
	"github.com/stretchr/testify/assert"
)

func TestShouldSignInBatch(t *testing.T) {
	var apiPath = utils.GetApiRoute() + "sign/"

	const signBatchMock = `"{

		"user_token": "bec3dd89-a2d8-4d64-ad5e-86b4b98b8557",
		"signer_tokens":[
			"57cc1429-2064-4a6b-ac47-4cf95372c56a",
			"56097301-86c1-42a8-b48f-bc59d798467c"
		]
	}"`

	mockResponse := `{"message":"Documento(s) assinado(s) com sucesso. Lembrete: este endpoint é assíncrono, então aguarde os PDF finais ficarem prontos via webhooks ou confira-os daqui alguns minutos."}`
	responseRecorderStatusCode := httptest.NewRecorder().Code
	statusCode, body := controllers.PostRequest(signBatchMock, utils.GetApiToken(), apiPath)

	assert.Equal(t, statusCode, responseRecorderStatusCode)
	assert.Equal(t, mockResponse, string(body))

}
