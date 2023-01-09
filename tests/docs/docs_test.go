package tests

import (
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/models"
	"github.com/ZapSign/SdkGo/utils"
)

//RUN FOLDER TESTS: go test -v ./tests/docs
//RUN UNIT TEST: go test -run TestShouldCreateDocFromUploadAsync ./tests/docs/ -v

func TestShouldCreateADocFromAUploudPdf(t *testing.T) {
	var apiPath = utils.GetApiRoute() + "docs/"

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

	controllers.PostRequest(docMock,utils.GetApiToken(),apiPath)
}
func TestShouldCreateADocFromAUploudDocx(t *testing.T) {
	var apiPath = utils.GetApiRoute() + "docs/"

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
		Url_pdf: "https://zapsign.s3.amazonaws.com/2022/1/docs/d7660fd2-fe74-4691-bec8-5c42c0ae2b3f/39a35070-8987-476d-86e3-75d91f588a5a.docx",
		Brand_primary_color: "#000000",
		Lang: "pt-br",
		Signers: signersMock,
	}

	controllers.PostRequest(docMock,utils.GetApiToken(),apiPath)
}

func TestShouldCreateDocFromUploadAsync(t *testing.T) {
	var apiPath = utils.GetApiRoute() + "docs/async/"

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

	controllers.PostRequest(docMock,utils.GetApiToken(),apiPath)
}

func TestShouldCreateDocFromTemplate(t *testing.T) {
	var apiPath = utils.GetApiRoute() + "models/create-doc/"

	deparatemplateMock := []models.DeParaTemplate {
		{
			De: "{{EMAIL CLIENTE}}",
			Para: "cliente@gmail.com",
		},
		{
			De: "{{NOME COMPLETO}}",
			Para: "João dos Santos",
		},
		{
			De: "{{NÚMERO DO CPF}}",
			Para: "123.456.789-10",
		},
	}
	docfromtemplateMock := models.DocFromTemplate{
			Sandbox: false,
			Template_id:"21620517-abaa-4cb9-8f19-d7f136d61810",
			Signer_name:"João dos Santos",
			External_id: "123",
			Brand_primary_color: "#000000",
			Lang: "pt-br",
			Deparafromtemplate: deparatemplateMock,
	}

	controllers.PostRequest(docfromtemplateMock,utils.GetApiToken(),apiPath)
}

func TestShouldCreateDocFromTemplateAsync(t *testing.T) {
	var apiPath = utils.GetApiRoute() + "models/create-doc/async/"

	deparatemplateMock := []models.DeParaTemplate {
		{
			De: "{{EMAIL CLIENTE}}",
			Para: "cliente@gmail.com",
		},
		{
			De: "{{NOME COMPLETO}}",
			Para: "João dos Santos",
		},
		{
			De: "{{NÚMERO DO CPF}}",
			Para: "123.456.789-10",
		},
	}
	docfromtemplateMock := models.DocFromTemplate{
			Sandbox: false,
			Template_id:"21620517-abaa-4cb9-8f19-d7f136d61810",
			Signer_name:"João dos Santos",
			External_id: "123",
			Brand_primary_color: "#000000",
			Lang: "pt-br",
			Deparafromtemplate: deparatemplateMock,
	}

	controllers.PostRequest(docfromtemplateMock,utils.GetApiToken(),apiPath)
}

func TestShouldCreateAExtraDoc(t *testing.T) {
	var apiPath = utils.GetApiRoute() + "docs/"+ utils.GetDocToken() + "/upload-extra-doc/"

	extradoc := models.ExtraDoc {
		Name: "Anexo ao Contrato de Admissão João",
		Base64_pdf: `"JVBERi0xLjcKCjEgMCBvYmogICUgZW50cnkgcG9pbnQKPDwKICAvVHlwZSAvQ2F0YWxvZwog
		IC9QYWdlcyAyIDAgUgo+PgplbmRvYmoKCjIgMCBvYmoKPDwKICAvVHlwZSAvUGFnZXMKICAvTWVkaWFCb3ggWyAw
		IDAgMjAwIDIwMCBdCiAgL0NvdW50IDEKICAvS2lkcyBbIDMgMCBSIF0KPj4KZW5kb2JqCgozIDAgb2JqCjw8CiAgL1
		R5cGUgL1BhZ2UKICAvUGFyZW50IDIgMCBSCiAgL1Jlc291cmNlcyA8PAogICAgL0ZvbnQgPDwKICAgICAgL0YxIDQg
		MCBSIAogICAgPj4KICA+PgogIC9Db250ZW50cyA1IDAgUgo+PgplbmRvYmoKCjQgMCBvYmoKPDwKICAvVHlwZSAvRm9ud
		AogIC9TdWJ0eXBlIC9UeXBlMQogIC9CYXNlRm9udCAvVGltZXMtUm9tYW4KPj4KZW5kb2JqCgo1IDAgb2JqICAlIHBhZ2U
		gY29udGVudAo8PAogIC9MZW5ndGggNDQKPj4Kc3RyZWFtCkJUCjcwIDUwIFRECi9GMSAxMiBUZgooSGVsbG8sIHdvcmxkI
		SkgVGoKRVQKZW5kc3RyZWFtCmVuZG9iagoKeHJlZgowIDYKMDAwMDAwMDAwMCA2NTUzNSBmIAowMDAwMDAwMDEwIDAwMDAwI
		G4gCjAwMDAwMDAwNzkgMDAwMDAgbiAKMDAwMDAwMDE3MyAwMDAwMCBuIAowMDAwMDAwMzAxIDAwMDAwIG4gCjAwMDAwMDAzO
		DAgMDAwMDAgbiAKdHJhaWxlcgo8PAogIC9TaXplIDYKICAvUm9vdCAxIDAgUgo+PgpzdGFydHhyZWYKNDkyCiUlRU9G",`,
	}

	controllers.PostRequest(extradoc,utils.GetApiToken(),apiPath)
}

func TestShouldGetDetailsDocument(t *testing.T) {
	var apiPath = utils.GetApiRoute() + "docs/" + utils.GetDocToken() + "/"
	controllers.GetRequest(utils.GetApiToken(),apiPath)
}

func TestShouldGetAllDocuments(t *testing.T) {
	var apiPath = utils.GetApiRoute() + "docs/"

	controllers.GetRequest(utils.GetApiToken(),apiPath)
}

func TestShouldDeleteDocument(t *testing.T) {
	var apiPath = utils.GetApiRoute() + "docs/" + utils.GetDocToken() + "/"

	controllers.DeleteRequest(utils.GetApiToken(),apiPath)
}

func TestShouldPlaceSignatures(t *testing.T) {
	var apiPath = utils.GetApiRoute() + "docs/" + utils.GetDocToken() + "/place-signatures/"

	rubricasMock := []models.Rubricas{
        {
            Page: 0,
            Relative_position_bottom: 52.50,
            Relative_position_left: 75.71,
            Relative_size_x: 19.55,
            Relative_size_y: 9.42,
            Type: "signature",
            Signer_token: "afd082e6-bfeb-4d21-8658-d845349b9dff",
        },
        {
            Page: 0,
            Relative_position_bottom: 13.50,
            Relative_position_left: 20.71,
            Relative_size_x: 19.55,
            Relative_size_y: 9.42,
            Type: "visto",
            Signer_token: "afd082e6-bfeb-4d21-8658-d845349b9dff",
        },
	}

	controllers.PostRequest(rubricasMock, utils.GetApiToken(), apiPath)
}