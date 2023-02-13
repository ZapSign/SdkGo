package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/ZapSign/SdkGo/controllers"
	"github.com/ZapSign/SdkGo/models"
	"github.com/ZapSign/SdkGo/utils"
	"github.com/stretchr/testify/assert"
)

func TestShouldStatus200WhenCreateExtraDoc(t *testing.T) {
	var apiPath = utils.GetDocsRoute() + utils.GetDocToken() + "/upload-extra-doc/"
	var getAPIToken = utils.GetApiToken()

	extraDoc := models.ExtraDoc.CreateExtraDoc(models.ExtraDoc{})

	statusCodeRequest, _ := controllers.PostRequest(extraDoc, getAPIToken, apiPath)
	responseRecorderStatusCode := httptest.NewRecorder().Code
	assert.Equal(t, statusCodeRequest, responseRecorderStatusCode)

}
