package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/utils"
	"github.com/stretchr/testify/assert"
)

func TestShouldStatus200WhenDetailSignerFromDoc(t *testing.T) {
	var apiDetailSignerPath = utils.GetApiRoute() + "signers/" + utils.GetSignerToken()

	responseRecorderStatusCode := httptest.NewRecorder().Code
	statusCode, _ := controllers.GetRequest(apiDetailSignerPath)
	assert.Equal(t, statusCode, responseRecorderStatusCode)
}
