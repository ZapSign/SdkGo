package services

import (
	"io/ioutil"
	"log"
	"net/http"
)

func ReadStatusCodeAndBodyRequest(response *http.Response) (int, string) {
	body, errorReadBody := ioutil.ReadAll(response.Body)
	if errorReadBody != nil {
		log.Fatalln(errorReadBody)
	}

	return response.StatusCode, string(body)
}
