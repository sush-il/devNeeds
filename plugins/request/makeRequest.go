package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func MakeRequest(url string, requestMethod string, headers string, body string) (*http.Response, error){

	request, err := http.NewRequest(requestMethod, url, bytes.NewBufferString(body))

	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	
	var headerObject map[string]string
	if headers != "" {
		if err := json.Unmarshal([]byte(headers), &headerObject); err != nil {
			return nil, fmt.Errorf("Headers could not be parsed correctly.")
		}

		for key, value := range headerObject {
			request.Header.Set(key, value)
		}
	}

	response, err := http.DefaultClient.Do(request)
	
	if err != nil {
		return nil, fmt.Errorf("Failed to get a response, %w", err)
	}

	return response, nil
}
