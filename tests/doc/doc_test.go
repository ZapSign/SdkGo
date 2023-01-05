package tests

import (
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/models"
)

var apiToken = "b1c75fcf-96d5-4951-b084-36870d2dde22d984198b-6186-4ac8-8135-bf3e586f0a34"

func CriaDocMock() models.Doc {
	signersMock := []models.Signer{
        {
            Name: "My First Signer",
        },
		{
			Name:"My Second Signer",
			Email:"test@test.com",
			Lock_email:true,
			Lock_phone:true,
			Phone_country:"YOUR_COUNTRY_PREFIX",
			Phone_number:"YOUR_PHONE_NUMBER",
			Auth_mode:"assinaturaTela",
			Send_automatic_email: false,
			Send_automatic_whatsapp: true,
		},	
	}

	docMock := models.Doc{
		Sandbox: false,
		Name: "My Contract Golang",
		Url_pdf: "https://zapsign.s3.amazonaws.com/2022/1/pdf/63d19807-cbfa-4b51-8571-215ad0f4eb98/ca42e7be-c932-482c-b70b-92ad7aea04be.pdf",
		Brand_primary_color: "#000000",
		Lang: "pt-br",
		Signers: signersMock,
	}

	return docMock

}

func TestShouldCreateADocFromAUploudPdf(t *testing.T) {
	const apiRoute = "http://localhost:3001/api/v1/docs/"

	docMock := CriaDocMock()
	controllers.CreateDocFromUploadPdf(docMock,apiToken,apiRoute)

}

func TestShouldFailedCreateADocFromAUploudPdf(t *testing.T) {
	const apiRoute = "http://localhost:3001/api/v1/docs/"

	docMock := CriaDocMock()
	controllers.CreateDocFromUploadPdf(docMock,apiToken,apiRoute)

}
