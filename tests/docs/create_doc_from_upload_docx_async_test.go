package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/models"
	"github.com/ZapSign/SdkGo/utils"
	"github.com/stretchr/testify/assert"
)

func TestShouldStatus200WhenCreateDocumentWithDocxFileAsync(t *testing.T) {
	var apiDocsRoutePath = utils.GetDocsRoute() + "async/"
	var getAPIToken = utils.GetApiToken()

	signersMock := models.Signer.CreateSigner(models.Signer{})
	docMock := models.Doc.CreateDoc(models.Doc{}, signersMock)

	responseRecorderStatusCode := httptest.NewRecorder()
	statusCodeRequest, _ := controllers.PostRequest(docMock, getAPIToken, apiDocsRoutePath)

	assert.Equal(t, statusCodeRequest, responseRecorderStatusCode.Code)
}
