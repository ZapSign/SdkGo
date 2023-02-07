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

	signerTokens := [...]string{"b15edab2-b691-4ecb-833a-564b3490963e", "9934376a-58fd-443a-ab9d-f5210895e0e0"}

	signBatchMock := models.SignBatchAPI{
		User_token:    "1c166935-5b8a-4059-a205-4f87d7119f9b",
		Signer_tokens: signerTokens,
	}
	responseRecorderStatusCode := httptest.NewRecorder().Code
	statusCode, _ := controllers.PostRequest(signBatchMock, getAPIToken, apiSignBatchPath)

	assert.Equal(t, statusCode, responseRecorderStatusCode)

}
