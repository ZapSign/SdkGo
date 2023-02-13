package tests

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/utils"
	"github.com/stretchr/testify/assert"
)

func TestShouldStatus200WhenDetailSignerFromDoc(t *testing.T) {
	var apiDetailSignerPath = utils.GetSignersRoute() + utils.GetSignerToken()

	responseRecorderStatusCode := httptest.NewRecorder().Code
	statusCode, body := controllers.GetRequest(apiDetailSignerPath)
	fmt.Println(body)
	assert.Equal(t, statusCode, responseRecorderStatusCode)
}
