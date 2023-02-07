package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/models"
	"github.com/ZapSign/SdkGo/utils"
	"github.com/stretchr/testify/assert"
)

func TestShouldStatus200WhenCreateDocumentFromTemplate(t *testing.T) {
	var apiCreateTemplatePath = utils.GetApiRoute() + "models/create-doc/"
	var getAPIToken = utils.GetApiToken()

	deParaTemplateMock := models.DeParaTemplate.CreateDeParaTemplate(models.DeParaTemplate{})
	docFromTemplateMock := models.DocFromTemplate.CreateDocFromTemplate(models.DocFromTemplate{}, deParaTemplateMock)

	statusCodeRequest, _ := controllers.PostRequest(docFromTemplateMock, getAPIToken, apiCreateTemplatePath)
	responseRecorderStatusCode := httptest.NewRecorder().Code
	assert.Equal(t, statusCodeRequest, responseRecorderStatusCode)
}
