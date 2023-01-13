package tests

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/models"
	"github.com/ZapSign/SdkGo/utils"
	"github.com/stretchr/testify/assert"
)

func TestShouldStatus400WhenParameterNameNotExistsInCreateSigner(t *testing.T) {
	var apiAddSignerPath = utils.GetDocsRoute() + utils.GetDocToken() + "/add-signer/"
	var getAPIToken = utils.GetApiToken()

	mockResponseErrorNotParameterName := "Missing required parameter 'name'."
	signerMock := models.Signer{
		Email: "email@test.com",
	}

	responseRecorderStatusCode := httptest.NewRecorder().Code
	statusCode, body := controllers.PostRequest(signerMock, getAPIToken, apiAddSignerPath)

	fmt.Println(string(body))
	assert.Equal(t, statusCode, responseRecorderStatusCode)
	assert.Equal(t, mockResponseErrorNotParameterName, string(body))

}
