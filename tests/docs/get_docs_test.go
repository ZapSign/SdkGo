package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/utils"
	"github.com/stretchr/testify/assert"
)

func TestShouldReceiveStatus200WhenWantAllDocuments(t *testing.T) {
	var docsRoutePath = utils.GetDocsRoute()

	responseRecorderStatusCode := httptest.NewRecorder().Code
	statusCodeRequest, _ := controllers.GetRequest(docsRoutePath)
	assert.Equal(t, statusCodeRequest, responseRecorderStatusCode)
}
