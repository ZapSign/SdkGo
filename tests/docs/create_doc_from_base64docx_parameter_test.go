package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/models"
	"github.com/ZapSign/SdkGo/utils"
	"github.com/stretchr/testify/assert"
)

func TestShouldStatus200WhenCreateDocumentWithBase64DocxParameter(t *testing.T) {
	var docsRoutePath = utils.GetDocsRoute()
	var getAPIToken = utils.GetApiToken()

	signersMock := models.Signer.CreateSigners(models.Signer{})

	Base64DocxMock := models.Base64Docx.CreateDocWithBase64Docx(models.Base64Docx{}, signersMock)

	responseRecorderStatusCode := httptest.NewRecorder().Code
	statusCodeRequest, _ := controllers.PostRequest(Base64DocxMock, getAPIToken, docsRoutePath)

	assert.Equal(t, statusCodeRequest, responseRecorderStatusCode)
}
