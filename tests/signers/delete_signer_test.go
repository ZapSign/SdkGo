package tests

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/utils"
	"github.com/stretchr/testify/assert"
)

func TestShouldStatus200WhenRemoveSignerFromDoc(t *testing.T) {
	var apiDeleteSignerPath = utils.GetSignerRoute() + utils.GetSignerTokenThatWillBeDeleted() + "/remove/"
	var getAPIToken = utils.GetApiToken()

	mockResponse := "Signatário removido com sucesso."

	responseRecorderStatusCode := httptest.NewRecorder().Code
	statusCode, body := controllers.DeleteRequest(getAPIToken, apiDeleteSignerPath)

	assert.Equal(t, statusCode, responseRecorderStatusCode)
	fmt.Println(body)
	assert.Equal(t, mockResponse, string(body))
}
