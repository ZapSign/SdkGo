package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/models"
	"github.com/ZapSign/SdkGo/utils"
	"github.com/stretchr/testify/assert"
)

func TestShouldStatus200WhenCreateDocumentWithPdfFile(t *testing.T) {
	var docsRoutePath = utils.GetDocsRoute()
	var getAPIToken = utils.GetApiToken()

	signersMock := models.Signer.CreateSigners(models.Signer{})

	docMock := models.Doc.CreateDocWithUrlPdf(models.Doc{}, signersMock)

	responseRecorderStatusCode := httptest.NewRecorder().Code
	statusCodeRequest, _ := controllers.PostRequest(docMock, getAPIToken, docsRoutePath)

	assert.Equal(t, statusCodeRequest, responseRecorderStatusCode)
}
