package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ZapSign/SdkGo/services"
	"github.com/ZapSign/SdkGo/utils"
)

func GetRequest(apiRoute string) (int, string) {
	request, errorRequest := http.NewRequest(http.MethodGet, apiRoute, nil)
	if errorRequest != nil {
		log.Fatalln(errorRequest)
	}

	utils.AddQueryParamsToRequest(request)
	fmt.Println(request.URL.String())

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(errorRequest)
	}

	statusCode, body := services.ReadStatusCodeAndBodyRequest(resp)

	return statusCode, body
}
func PostRequest(docMock interface{}, apiToken string, apiRoute string) (int, string) {
	payload := utils.ConvertObjectToJSON(docMock)

	request, error := http.NewRequest(http.MethodPost, apiRoute, strings.NewReader(payload))

	if error != nil {
		log.Fatalln(error)
	}

	utils.AddHeadersFromRequest(request)
	utils.AddQueryParamsToRequest(request)
	fmt.Println(request.URL.String())

	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	statusCode, body := services.ReadStatusCodeAndBodyRequest(resp)
	return statusCode, body
}

func DeleteRequest(apiRoute string) (int, string) {
	request, errorRequest := http.NewRequest(http.MethodDelete, apiRoute, nil)
	utils.AddQueryParamsToRequest(request)

	if errorRequest != nil {
		log.Fatalln(errorRequest)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	statusCode, body := services.ReadStatusCodeAndBodyRequest(resp)
	return statusCode, body
}
