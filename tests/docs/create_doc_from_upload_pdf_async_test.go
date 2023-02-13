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

func TestShouldStatus200WhenCreateDocumentWithPdfFileAsync(t *testing.T) {
	var docsRoutePath = fmt.Sprintf("%s%s", utils.GetDocsRoute(), "async/")
	var getAPIToken = utils.GetApiToken()

	fakeSignersMock := models.Signer.CreateSigners(models.Signer{})
	docWithUrlPdfMock := models.Doc.CreateDocWithUrlPdf(models.Doc{}, fakeSignersMock)

	responseRecorderStatusCode := httptest.NewRecorder().Code
	statusCodeRequest, _ := controllers.PostRequest(docWithUrlPdfMock, getAPIToken, docsRoutePath)

	assert.Equal(t, statusCodeRequest, responseRecorderStatusCode)
}
