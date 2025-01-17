package util

import (
	"bytes"
	"net/http"
	url_config "net/url"
)

func CreateHttpRequest(method string, url string, header map[string]string, body map[string]string) (*http.Response, error) {
	param := url_config.Values{}
	if body != nil {
		for k, v := range body { 
			param.Set(k, v)
		}	
	}	
	request_body := bytes.NewBuffer([]byte(param.Encode()))
	request, err := http.NewRequest(method, url, request_body)

	if err != nil { return nil, err }
	if header != nil { 
		for k, v := range header { 
			request.Header.Add(k, v) 
		} 
	}

	return http.DefaultClient.Do(request)
}
