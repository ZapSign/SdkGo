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
	var apiUpdateSignerPath = utils.GetApiRoute() + "signers/" + utils.GetSignerToken() + "/"
	var getAPIToken = utils.GetApiToken()

	signerMock := models.Signer{
		Lock_email: false,
		Lock_phone: false,
		Name:       "Jhon Doe",
	}
	responseRecorderStatusCode := httptest.NewRecorder().Code
	statusCode, _ := controllers.PostRequest(signerMock, getAPIToken, apiUpdateSignerPath)

	assert.Equal(t, statusCode, responseRecorderStatusCode)
}
