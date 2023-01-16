package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/models"
	"github.com/ZapSign/SdkGo/utils"
	"github.com/stretchr/testify/assert"
)

func TestShouldStatus200WhenUpdateSignerFromDoc(t *testing.T) {
	var apiUpdateSignerPath = utils.GetSignersRoute() + utils.GetSignerToken() + "/"
	var getAPIToken = utils.GetApiToken()

	signerMock := models.Signer{
		Lock_email:    true,
		Lock_phone:    true,
		Phone_country: "55",
		Name:          "Felipe S",
		Email:         "test@test.com",
	}
	responseRecorderStatusCode := httptest.NewRecorder().Code
	statusCode, _ := controllers.PostRequest(signerMock, getAPIToken, apiUpdateSignerPath)

	assert.Equal(t, statusCode, responseRecorderStatusCode)
}
