package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/models"
	"github.com/ZapSign/SdkGo/utils"
	"github.com/stretchr/testify/assert"
)

//RUN FOLDER TESTS: go test -v ./tests/docs
//RUN UNIT TEST: go test -v -run TestShouldCreateDocument

func TestShouldCreateDocumentWithPdfFile(t *testing.T) {
	var apiPath = utils.GetApiRoute() + "docs/"

	signersMock := []models.Signer{
		{
			Name:                    "João Carlos",
			Email:                   "test@test.com",
			Auth_mode:               "assinaturaTela",
			Send_automatic_email:    false,
			Send_automatic_whatsapp: true,
		},
		{
			Name: "Siclano Almeida",
		},
	}

	docMock := models.Doc{
		Sandbox:             true,
		Name:                "My Contract Golang",
		Url_pdf:             "https://zapsign.s3.amazonaws.com/2022/1/pdf/63d19807-cbfa-4b51-8571-215ad0f4eb98/ca42e7be-c932-482c-b70b-92ad7aea04be.pdf",
		Brand_primary_color: "#000000",
		Lang:                "pt-br",
		Signers:             signersMock,
	}

	responseRecorderStatusCode := httptest.NewRecorder()
	statusCodeRequest, _ := controllers.PostRequest(docMock, utils.GetApiToken(), apiPath)
	assert.Equal(t, statusCodeRequest, responseRecorderStatusCode.Code)
}

func TestShouldCreateDocumentWithDocxFile(t *testing.T) {
	var apiPath = utils.GetApiRoute() + "docs/"

	signersMock := []models.Signer{
		{
			Name:                    "João Carlos",
			Email:                   "test@test.com",
			Auth_mode:               "assinaturaTela",
			Send_automatic_email:    false,
			Send_automatic_whatsapp: true,
		},
		{
			Name: "Siclano Almeida",
		},
	}

	docMock := models.Doc{
		Sandbox:             true,
		Name:                "My Contract Golang",
		Url_pdf:             "https://zapsign.s3.amazonaws.com/2022/1/docs/d7660fd2-fe74-4691-bec8-5c42c0ae2b3f/39a35070-8987-476d-86e3-75d91f588a5a.docx",
		Brand_primary_color: "#000000",
		Lang:                "pt-br",
		Signers:             signersMock,
	}

	responseRecorderStatusCode := httptest.NewRecorder()
	statusCodeRequest, _ := controllers.PostRequest(docMock, utils.GetApiToken(), apiPath)
	assert.Equal(t, statusCodeRequest, responseRecorderStatusCode.Code)
}

func TestShouldCreateDocFromTemplate(t *testing.T) {
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
		Template_id:         "e7dd40ea-5698-413d-91c9-6d1ee18b4c82",
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

func TestShould402TheSandboxFlag(t *testing.T) {
	var apiPath = utils.GetApiRoute() + "docs/"

	signersMock := []models.Signer{
		{
			Name:                    "My Second Signer",
			Email:                   "test@test.com",
			Lock_email:              true,
			Lock_phone:              true,
			Phone_country:           "YOUR_COUNTRY_PREFIX",
			Phone_number:            "YOUR_PHONE_NUMBER",
			Auth_mode:               "assinaturaTela",
			Send_automatic_email:    false,
			Send_automatic_whatsapp: true,
		},
	}

	docMock := models.Doc{
		Sandbox:             false,
		Name:                "My Contract Golang",
		Url_pdf:             "https://zapsign.s3.amazonaws.com/2022/1/pdf/63d19807-cbfa-4b51-8571-215ad0f4eb98/ca42e7be-c932-482c-b70b-92ad7aea04be.pdf",
		Brand_primary_color: "#000000",
		Lang:                "pt-br",
		Signers:             signersMock,
	}
	mockResponse := "É obrigatório contratar um Plano de API para utilizá-la em modo produção. Caso seja um teste, coloque a flag \"sandbox\" como true e use no modo desenvolvimento."
	statusCodeRequest, body := controllers.PostRequest(docMock, utils.GetApiToken(), apiPath)

	assert.Equal(t, statusCodeRequest, http.StatusPaymentRequired)
	assert.Equal(t, mockResponse, string(body))
}
