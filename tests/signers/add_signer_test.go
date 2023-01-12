package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/models"
	"github.com/ZapSign/SdkGo/utils"
	"github.com/stretchr/testify/assert"
)

func TestShouldStatus200AddSignerFromDoc(t *testing.T) {
	var apiPath = utils.GetDocsRoute() + utils.GetDocToken() + "/add-signer/"
	var getAPIToken = utils.GetApiToken()

	signerMock := models.Signer{
		Name:                    "Novinho",
		Email:                   "email@test.com",
		Lock_email:              true,
		Lock_phone:              true,
		Phone_country:           "55",
		Phone_number:            "99999999999",
		Auth_mode:               "assinaturaTela",
		Send_automatic_email:    false,
		Send_automatic_whatsapp: false,
	}
	responseRecorderStatusCode := httptest.NewRecorder().Code
	statusCode, _ := controllers.PostRequest(signerMock, getAPIToken, apiPath)
	assert.Equal(t, statusCode, responseRecorderStatusCode)
}
