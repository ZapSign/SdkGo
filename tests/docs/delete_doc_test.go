package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/utils"
	"github.com/stretchr/testify/assert"
)

func TestShouldHaveStatus200WhenWantDeleteDoc(t *testing.T) {
	var apiDeleteDocumentPath = utils.GetDocsRoute() + utils.GetDocTokenThatWillBeDeleted()

	responseRecorderStatusCode := httptest.NewRecorder().Code
	statusCodeRequest, _ := controllers.DeleteRequest(apiDeleteDocumentPath)

	assert.Equal(t, statusCodeRequest, responseRecorderStatusCode)
}
