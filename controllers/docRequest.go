package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ZapSign/SdkGo/models"
)

func ConvertObjectToJSON (docMock models.Doc) string {
	jsonDoc, err := json.Marshal(docMock)
	if err != nil {
        fmt.Printf("Error: %s", err)
    }
    return string(jsonDoc)
}

func CreateDocFromUploadPdf(docMock models.Doc, apiToken string, apiRoute string)  {
  payload := ConvertObjectToJSON(docMock)

  client := &http.Client {}
  req, err := http.NewRequest("POST", apiRoute, strings.NewReader(payload))

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Authorization", "Bearer " + apiToken)
  req.Header.Add("Content-Type", "application/json")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}