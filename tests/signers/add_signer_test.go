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

	signerMock := models.Signer.CreateSigner(models.Signer{})

	responseRecorderStatusCode := httptest.NewRecorder().Code
	statusCode, _ := controllers.PostRequest(signerMock, getAPIToken, apiPath)
	assert.Equal(t, statusCode, responseRecorderStatusCode)
}
