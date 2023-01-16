package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/utils"
	"github.com/stretchr/testify/assert"
)

func TestShouldHaveStatus200WhenWantDetailsOfDocument(t *testing.T) {
	var apiDetailDocumentPath = utils.GetDocsRoute() + utils.GetDocToken()

	responseRecorderStatusCode := httptest.NewRecorder().Code
	statusCodeRequest, _ := controllers.GetRequest(apiDetailDocumentPath)
	assert.Equal(t, statusCodeRequest, responseRecorderStatusCode)
}
