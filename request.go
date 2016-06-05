package pagarme

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

const (
	apiURL = "api.pagar.me/1"
)

//Do execute a new request to the API
func Do(method, path string, body []byte) (Response, error) {
	var response Response

	reqURL := fmt.Sprintf("http://%s/%s"), apiURL, path)
	req, err := http.NewRequest(method, reqURL, bytes.NewBuffer(body))

	if err != nil {
		return response, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return response, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)

	return response, err
}

