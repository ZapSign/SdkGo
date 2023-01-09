package tests

import (
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/models"
	"github.com/ZapSign/SdkGo/utils"
)

//RUN FOLDER TESTS: go test -v ./tests/signers
//RUN UNIT TEST: go test -run TestShouldCreateDocFromUploadAsync ./tests/docs/ -v

func TestShouldDetailSigner(t *testing.T) {
	var apiPath = utils.GetApiRoute() + "signers/" + utils.GetSignerToken()
	controllers.GetRequest(utils.GetApiToken(), apiPath)
}

func TestShouldUpdateSigner(t *testing.T) {
	var apiPath = utils.GetApiRoute() + "signers/" + utils.GetSignerToken() + "/"

	signerMock := models.Signer {
		Name: "New Name",
		Email: "newEmail@test.com",
		Lock_email: true,
		Lock_phone: true,
		Phone_country: "55",
		Phone_number: "99999999999",
		Auth_mode: "assinaturaTela",
		Send_automatic_email: false,
		Send_automatic_whatsapp: false,
	}
	controllers.PostRequest(signerMock,utils.GetApiToken(), apiPath)
}

func TestShouldAddSigner(t *testing.T) {
	var apiPath = utils.GetApiRoute() + "docs/" + utils.GetDocToken() + "/add-signer/"

	signerMock := models.Signer {
		Name: "New Name",
		Email: "newEmail@test.com",
		Lock_email: true,
		Lock_phone: true,
		Phone_country: "55",
		Phone_number: "99999999999",
		Auth_mode: "assinaturaTela",
		Send_automatic_email: false,
		Send_automatic_whatsapp: false,
	}

	controllers.PostRequest(signerMock,utils.GetApiToken(), apiPath)
}

func TestShouldDeleteSigner(t *testing.T){
	var apiPath = utils.GetApiRoute() + "signer/" + utils.GetDocToken() + "/remove/"

	controllers.DeleteRequest(utils.GetApiToken(), apiPath)
}

func TestShouldSignInBatch(t *testing.T){
	const signBatchMock = `"{
		"user_token": "23d4d2d5-8998-4516-b2f8-6045c4eabc9d",
		"signer_tokens":[
		"3c8f16f8-949b-432a-adea-d1e544bb91be",
		"fed91ab3-bb69-46f9-8c2d-9594772b1186"
		]
	}"`
	var apiPath = utils.GetApiRoute() + "sign/"

	controllers.PostRequest(signBatchMock,utils.GetApiToken(), apiPath)
}
