package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/utils"
	"github.com/stretchr/testify/assert"
)

func TestShouldReceiveStatus200WhenWantAllDocuments(t *testing.T) {
	var apiRoute = utils.GetApiRoute() + "docs/"

	responseRecorderStatusCode := httptest.NewRecorder().Code
	statusCodeRequest, _ := controllers.GetRequest(apiRoute)
	assert.Equal(t, statusCodeRequest, responseRecorderStatusCode)
}
