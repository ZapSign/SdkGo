package tests

import (
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/utils"
	"github.com/stretchr/testify/assert"
)

func TestShouldStatus500WhenNullBodyInUpdateSignerFromDoc(t *testing.T) {
	var apiUpdateSignerPath = utils.GetSignersRoute() + utils.GetSignerToken() + "/"
	var getAPIToken = utils.GetApiToken()

	statusCode, _ := controllers.PostRequest(nil, getAPIToken, apiUpdateSignerPath)

	assert.Equal(t, statusCode, 500)

}
