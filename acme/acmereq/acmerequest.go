package acmereq

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

// Execute ...
func Execute(acmeRequest *AcmeRequest, responseHandler func(httpResponse *http.Response) interface{}) (interface{}, error) {
	if acmeRequest.URL == nil {
		return nil, errors.New("Missing value for: URL")
	}

	if acmeRequest.Method == "" {
		return nil, errors.New("Missing value for: Method")
	}

	httpRequest, err := http.NewRequest(acmeRequest.Method, acmeRequest.URL.String(), nil)
	if err != nil {
		return nil, errors.New("Error creating HTTP request")
	}

	if acmeRequest.HasBody() {
		jsonBytes, err := json.Marshal(acmeRequest.Body)
		if err != nil {
			return nil, errors.New("Error marshalling request body")
		}

		httpRequest.Body = ioutil.NopCloser(bytes.NewReader(jsonBytes))
		httpRequest.Header.Add("Content-Type", "application/jose+json")
	}

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("Error executing HTTP request: %s", err)
	}

	if httpResponse.StatusCode != acmeRequest.ExpectedResponseCode {
		return nil, fmt.Errorf("actual status code (%d) does not match expected status code (%d)", httpResponse.StatusCode, acmeRequest.ExpectedResponseCode)
	}

	return responseHandler(httpResponse), nil
}
