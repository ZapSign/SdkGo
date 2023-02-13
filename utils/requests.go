package utils

import "net/http"

func AddHeadersFromRequest(req *http.Request) {
	req.Header.Add("Content-Type", "application/json")
}

func AddQueryParamsToRequest(req *http.Request) {
	getAPIToken := GetApiToken()

	q := req.URL.Query()
	q.Add("api_token", getAPIToken)
	req.URL.RawQuery = q.Encode()
}
