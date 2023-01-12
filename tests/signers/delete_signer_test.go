package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/utils"
	"github.com/stretchr/testify/assert"
)

func TestShouldStatus200WhenRemoveSignerFromDoc(t *testing.T) {
	var apiPath = utils.GetApiRoute() + "signer/" + utils.GetSignerTokenThatWillBeDeleted() + "/remove/"

	mockResponse := "Signatário removido com sucesso."

	responseRecorderStatusCode := httptest.NewRecorder().Code
	statusCode, body := controllers.DeleteRequest(utils.GetApiToken(), apiPath)

	assert.Equal(t, statusCode, responseRecorderStatusCode)
	assert.Equal(t, mockResponse, string(body))
}
