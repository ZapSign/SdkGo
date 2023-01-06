package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func ProcessRequestToAPI(req *http.Request, apiToken string) {
	client := &http.Client{}

	req.Header.Add("Authorization", "Bearer "+apiToken)
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