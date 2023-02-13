package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/models"
	"github.com/ZapSign/SdkGo/utils"
	"github.com/stretchr/testify/assert"
)

func TestShouldSignInBatch(t *testing.T) {
	var apiSignBatchPath = utils.GetApiRoute() + "sign/"
	var getAPIToken = utils.GetApiToken()

	signBatchMock := models.SignBatchAPI.CreateSignBatch(models.SignBatchAPI{})

	responseRecorderStatusCode := httptest.NewRecorder().Code
	statusCode, _ := controllers.PostRequest(signBatchMock, getAPIToken, apiSignBatchPath)

	assert.Equal(t, statusCode, responseRecorderStatusCode)

}
