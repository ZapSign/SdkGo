package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/models"
	"github.com/ZapSign/SdkGo/utils"
	"github.com/stretchr/testify/assert"
)

func TestShouldStatus200WhenCreateDocumentWithBase64PDFParameter(t *testing.T) {
	var docsRoutePath = utils.GetDocsRoute()
	var getAPIToken = utils.GetApiToken()

	signersMock := models.Signer.CreateSigners(models.Signer{})

	Base64PdfMock := models.Base64Pdf.CreateDocWithBase64Pdf(models.Base64Pdf{}, signersMock)

	responseRecorderStatusCode := httptest.NewRecorder().Code
	statusCodeRequest, _ := controllers.PostRequest(Base64PdfMock, getAPIToken, docsRoutePath)

	assert.Equal(t, statusCodeRequest, responseRecorderStatusCode)
}
