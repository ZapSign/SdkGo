package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ZapSign/SdkGo/services"
	"github.com/ZapSign/SdkGo/utils"
)

func PostRequest(docMock interface{}, apiToken string, apiRoute string) (int, string) {
	payload := utils.ConvertObjectToJSON(docMock)

	req, error := http.NewRequest("POST", apiRoute, strings.NewReader(payload))

	if error != nil {
		log.Fatal(error)
	}

	utils.AddHeadersFromRequest(req)
	utils.AddQueryParamsToRequest(req)

	fmt.Println(req.URL.String())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	statusCode, body := services.ReadStatusCodeAndBodyRequest(resp)
	return statusCode, body
}
