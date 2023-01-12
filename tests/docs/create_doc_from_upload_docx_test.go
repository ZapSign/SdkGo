package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/models"
	"github.com/ZapSign/SdkGo/utils"
	"github.com/stretchr/testify/assert"
)

func TestShouldStatus200WhenCreateDocumentWithDocxFile(t *testing.T) {
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
