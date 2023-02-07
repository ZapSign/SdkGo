package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/models"
	"github.com/ZapSign/SdkGo/utils"
	"github.com/stretchr/testify/assert"
)

func TestShouldStatus200WhenCreateDocumentWithPdfFileAsync(t *testing.T) {
	var docsRoutePath = utils.GetDocsRoute() + "async/"
	var getAPIToken = utils.GetApiToken()

	signersMock := []models.Signer{
		{
			Name:                    "João Mateus",
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

	responseRecorderStatusCode := httptest.NewRecorder().Code
	statusCodeRequest, _ := controllers.PostRequest(docMock, getAPIToken, docsRoutePath)

	assert.Equal(t, statusCodeRequest, responseRecorderStatusCode)
}
