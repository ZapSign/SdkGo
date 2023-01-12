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
	var apiPath = utils.GetApiRoute() + "models/create-doc/"

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
		Template_id:         "778e4a2e-be65-4dc1-911d-57de6e3ffde0",
		Signer_name:         "João dos Santos",
		External_id:         "123",
		Brand_primary_color: "#000000",
		Lang:                "pt-br",
		Deparafromtemplate:  deParaTemplateMock,
	}

	statusCodeRequest, _ := controllers.PostRequest(docFromTemplateMock, utils.GetApiToken(), apiPath)
	responseRecorderStatusCode := httptest.NewRecorder().Code
	assert.Equal(t, statusCodeRequest, responseRecorderStatusCode)
}
