package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/models"
	"github.com/ZapSign/SdkGo/utils"
	"github.com/stretchr/testify/assert"
)

func TestShouldStatus200WhenCreateDocumentFromTemplateAsync(t *testing.T) {
	var apiCreateTemplatePath = utils.GetApiRoute() + "models/create-doc/async/"
	var getAPIToken = utils.GetApiToken()

	deParaTemplateMock := []models.DeParaTemplate{
		{
			De:   "{{EMAIL CLIENTE}}",
			Para: "cliente@gmail.com",
		},
		{
			De:   "{{NOME COMPLETO}}",
			Para: "João dos Santos",
		},
		{
			De:   "{{NÚMERO DO CPF}}",
			Para: "123.456.789-10",
		},
	}
	docFromTemplateMock := models.DocFromTemplate{
		Sandbox:             true,
		Template_id:         "aea8a7f9-5591-4a3e-b243-38926c2b2e00",
		Signer_name:         "João dos Santos",
		External_id:         "123",
		Brand_primary_color: "#000000",
		Lang:                "pt-br",
		Deparafromtemplate:  deParaTemplateMock,
	}

	statusCodeRequest, _ := controllers.PostRequest(docFromTemplateMock, getAPIToken, apiCreateTemplatePath)
	responseRecorderStatusCode := httptest.NewRecorder().Code
	assert.Equal(t, statusCodeRequest, responseRecorderStatusCode)
}
