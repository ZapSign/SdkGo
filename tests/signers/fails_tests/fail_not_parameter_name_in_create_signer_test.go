package tests

import (
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
	signerMockNotParameterName := models.Signer{
		Email: "test@test.com",
	}

	statusCode, body := controllers.PostRequest(signerMockNotParameterName, getAPIToken, apiAddSignerPath)

	assert.Equal(t, statusCode, statusBadRequest)
	assert.Equal(t, mockResponseErrorNotParameterName, string(body))

}
