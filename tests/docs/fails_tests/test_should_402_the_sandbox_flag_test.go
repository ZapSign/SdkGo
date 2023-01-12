package tests

import (
	"net/http"
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/models"
	"github.com/ZapSign/SdkGo/utils"
	"github.com/stretchr/testify/assert"
)

const statusPaymentRequired = http.StatusPaymentRequired

func TestShould4Status02TheSandboxFlagIsDisabled(t *testing.T) {
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

	assert.Equal(t, statusCodeRequest, statusPaymentRequired)

	assert.Equal(t, mockResponse, string(body))
}
