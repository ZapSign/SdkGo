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

func TestShouldStatus200WhenCreateDocumentWithDocxFileAsync(t *testing.T) {
	var apiDocsRoutePath = utils.GetDocsRoute() + "async/"
	var getAPIToken = utils.GetApiToken()

	signersMock := models.Signer.CreateSigners(models.Signer{})

	docMock := models.Doc.CreateDocWithDocPdf(models.Doc{}, signersMock)

	responseRecorderStatusCode := httptest.NewRecorder()
	statusCodeRequest, body := controllers.PostRequest(docMock, getAPIToken, apiDocsRoutePath)
	fmt.Println(body)
	assert.Equal(t, statusCodeRequest, responseRecorderStatusCode.Code)
}
