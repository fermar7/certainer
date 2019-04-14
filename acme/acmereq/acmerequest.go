package acmereq

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/fermar7/certainer/acme/infra"
)

// AcmeRequest ...
type AcmeRequest struct {
	Body                 *infra.JSONWebToken
	URL                  *url.URL
	Method               string
	ExpectedResponseCode int
}

// HasBody ...
func (req *AcmeRequest) HasBody() bool {
	return req.Body != nil
}

// AcmeResponse ...
type AcmeResponse struct {
	Content    []byte
	Header     http.Header
	StatusCode int
}

// HasBody ...
func (req *AcmeResponse) HasBody() bool {
	return len(req.Content) > 0
}

// Execute ...
func Execute(acmeRequest *AcmeRequest) (*AcmeResponse, error) {
	if acmeRequest.URL == nil {
		return nil, errors.New("Missing value for: URL")
	}

	if acmeRequest.Method == "" {
		return nil, errors.New("Missing value for: Method")
	}

	var err error
	var httpRequest *http.Request

	if acmeRequest.HasBody() {
		requestBodyBytes, err := json.Marshal(acmeRequest.Body)
		if err != nil {
			return nil, errors.New("Error marshalling request body")
		}

		httpRequest, err = http.NewRequest(acmeRequest.Method, acmeRequest.URL.String(), bytes.NewReader(requestBodyBytes))
		if err != nil {
			return nil, errors.New("Error creating HTTP request with body")
		}

		httpRequest.Header.Add("Content-Type", "application/jose+json")

	} else {
		httpRequest, err = http.NewRequest(acmeRequest.Method, acmeRequest.URL.String(), nil)
		if err != nil {
			return nil, errors.New("Error creating HTTP request without body")
		}
	}

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("Error executing HTTP request: %s", err)
	}

	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != acmeRequest.ExpectedResponseCode {
		content, err := ioutil.ReadAll(httpResponse.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading http response body")
		}

		log.Fatalf(string(content))

		return nil, fmt.Errorf("actual status code (%d) does not match expected status code (%d)", httpResponse.StatusCode, acmeRequest.ExpectedResponseCode)
	}

	content, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading http response body")
	}

	return &AcmeResponse{
		Content:    content,
		Header:     httpResponse.Header,
		StatusCode: httpResponse.StatusCode,
	}, nil
}
