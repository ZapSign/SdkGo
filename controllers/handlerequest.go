package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ZapSign/SdkGo/services"
	"github.com/ZapSign/SdkGo/utils"
)

func PostRequest(docMock interface{}, apiToken string, apiRoute string)  {
  payload := utils.ConvertObjectToJSON(docMock)

  req, err := http.NewRequest("POST", apiRoute, strings.NewReader(payload))

  fmt.Println(apiRoute)
  fmt.Println(payload)

	if err != nil {
		fmt.Println(err)
		return
	}

  services.ProcessRequestToAPI(req,apiToken)
}

func GetRequest(apiToken string, apiRoute string)  {
  req, err := http.NewRequest("GET", apiRoute, nil)

  if err != nil {
		fmt.Println(err)
		return
	}

  services.ProcessRequestToAPI(req,apiToken)
}


func DeleteRequest(apiToken string, apiRoute string)  {
  req, err := http.NewRequest("DELETE", apiRoute, nil)

  if err != nil {
		fmt.Println(err)
		return
	}

  services.ProcessRequestToAPI(req,apiToken)
}




