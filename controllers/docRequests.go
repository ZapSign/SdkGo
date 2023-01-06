package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ZapSign/SdkGo/services"
)
func ConvertObjectToJSON (dockTemplate interface{}) string {
	jsonDoc, err := json.Marshal(dockTemplate)
	if err != nil {
        fmt.Printf("Error: %s", err)
    }
    return string(jsonDoc)
}

func PostRequest(docMock interface{}, apiToken string, apiRoute string)  {
  payload := ConvertObjectToJSON(docMock)

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




