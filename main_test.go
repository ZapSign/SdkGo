package main

import (
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/models"
)

// go test -v -run TestShouldCreateDocFromUploadAsync

const apiToken = "b1c75fcf-96d5-4951-b084-36870d2dde22d984198b-6186-4ac8-8135-bf3e586f0a34"
const apiRoute = "http://localhost:3001/api/v1/"

func TestShouldCreateADocFromAUploudPdf(t *testing.T) {
	const apiPath = apiRoute + "docs/"

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

	controllers.PostRequest(docMock,apiToken,apiPath)
}

func TestShouldCreateADocFromAUploudDocx(t *testing.T) {
	const apiPath = apiRoute + "docs/"
	
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

	controllers.PostRequest(docMock,apiToken,apiPath)
}

func TestShouldCreateDocFromUploadAsync(t *testing.T) {
	const apiPath = apiRoute + "docs/async/"

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

	controllers.PostRequest(docMock,apiToken,apiPath)
}

func TestShouldCreateDocFromTemplate(t *testing.T) {
	const apiPath = apiRoute + "models/create-doc/"

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

	controllers.PostRequest(docfromtemplateMock,apiToken,apiPath)
}

func TestShouldCreateDocFromTemplateAsync(t *testing.T) {
	const apiPath = apiRoute + "models/create-doc/async/"

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

	controllers.PostRequest(docfromtemplateMock,apiToken,apiPath)
}

func TestShouldCreateAExtraDoc(t *testing.T) {
	const doctoken = "2e3e805e-708e-433e-bf2e-2105d59d76f9"
	const apiPath = apiRoute + "docs/"+ doctoken + "/upload-extra-doc/"

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

	controllers.PostRequest(extradoc,apiToken,apiPath)
}

func TestShouldGetDetailsDocument(t *testing.T) {
	const doctoken = "2e3e805e-708e-433e-bf2e-2105d59d76f9"
	const apiPath = apiRoute + "docs/" + doctoken + "/"

	controllers.GetRequest(apiToken,apiPath)
}

func TestShouldGetAllDocuments(t *testing.T) {
	const apiPath = apiRoute + "docs/"

	controllers.GetRequest(apiToken,apiPath)
}

func TestShouldDeleteDocument(t *testing.T) {
	const doctoken = "dbbfae2b-617d-4443-b3cb-eff626bca499"
	const apiPath = apiRoute + "docs/" + doctoken + "/"
	
	controllers.DeleteRequest(apiToken,apiPath)
}

func TestShouldPlaceSignatures(t *testing.T) {
	const doctoken = "14f88fc1-8eb5-4bcd-83a6-340dc4a59a7f"
	const apiPath = apiRoute + "docs/" + doctoken + "/place-signatures/"
	
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

	controllers.PostRequest(rubricasMock, apiToken, apiPath)
}