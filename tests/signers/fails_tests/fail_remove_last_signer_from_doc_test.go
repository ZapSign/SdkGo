package tests

import (
	"net/http"
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/utils"
	"github.com/stretchr/testify/assert"
)

const statusBadRequest = http.StatusBadRequest

func TestShouldStatus400WhenRemoveLastSignerFromDoc(t *testing.T) {
	var apiDeleteSignerPath = utils.GetSignerRoute() + utils.GetLastSignerTokenThatWillBeDeleted() + "/remove/"
	var getAPIToken = utils.GetApiToken()

	mockResponseErrorDeleteLastSigner := "Erro: Não é possível remover o único signatário do documento. Adicione outro signatário antes de remover este."

	statusCode, body := controllers.DeleteRequest(getAPIToken, apiDeleteSignerPath)

	assert.Equal(t, statusCode, statusBadRequest)
	assert.Equal(t, mockResponseErrorDeleteLastSigner, string(body))
}
