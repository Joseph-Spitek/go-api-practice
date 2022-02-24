package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func PerformGetRequest(url string, response interface{}) error {
	log.Printf("Fetching data from url: %s", url)

	httpResp, httpErr := http.Get(url)

	if httpErr != nil {
		return httpErr
	}
	defer httpResp.Body.Close()

	log.Printf("Response returned with status: %s and status code: %d", httpResp.Status, httpResp.StatusCode)

	respBody, readErr := ioutil.ReadAll(httpResp.Body)

	if readErr != nil {
		return readErr
	}

	jsonErr := json.Unmarshal(respBody, &response)
	if jsonErr != nil {
		return jsonErr
	}

	return nil
}
