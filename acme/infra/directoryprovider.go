package infra

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const directoryKey string = "/directory"

// AcmeDirectory ...
type AcmeDirectory struct {
	NewNonce   AcmeRessource
	NewAccount AcmeRessource
	NewOrder   AcmeRessource
	NewAuthz   AcmeRessource
	RevokeCert AcmeRessource
	KeyChange  AcmeRessource
}

// AcmeRessource ...
type AcmeRessource struct {
	URL       *url.URL
	Available bool
}

type directoryJSONModel struct {
	NewNonce   string `json:"newNonce"`
	NewAccount string `json:"newAccount"`
	NewOrder   string `json:"newOrder"`
	NewAuthz   string `json:"newAuthz"`
	RevokeCert string `json:"revokeCert"`
	KeyChange  string `json:"keyChange"`
}

// CreateDirectoryProvider ...
func CreateDirectoryProvider(authority *url.URL) (*AcmeDirectory, error) {
	requestURL, err := authority.Parse(directoryKey)
	if err != nil {
		return nil, fmt.Errorf("Error parsing directory URL: %s", err)
	}

	httpRequest, err := http.NewRequest(http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating directory request: %s", err)
	}

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("Error doing directory request: %s", err)
	}

	if httpResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error doing directory request: Invalid response code (%d)", httpResponse.StatusCode)
	}
	defer httpResponse.Body.Close()

	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading directory response content: %s", err)
	}

	var directoryJSONModel directoryJSONModel
	err = json.Unmarshal(body, &directoryJSONModel)
	if err != nil {
		return nil, fmt.Errorf("Error parsing directory response json: %s", err)
	}

	acmeDirectory := &AcmeDirectory{
		NewNonce:   AcmeRessource{Available: false},
		NewAccount: AcmeRessource{Available: false},
		NewOrder:   AcmeRessource{Available: false},
		NewAuthz:   AcmeRessource{Available: false},
		RevokeCert: AcmeRessource{Available: false},
		KeyChange:  AcmeRessource{Available: false},
	}

	if directoryJSONModel.NewNonce != "" {
		acmeDirectory.NewNonce.URL, err = url.Parse(directoryJSONModel.NewNonce)
		if err == nil {
			acmeDirectory.NewNonce.Available = true
		}
	}

	if directoryJSONModel.NewAccount != "" {
		acmeDirectory.NewAccount.URL, err = url.Parse(directoryJSONModel.NewAccount)
		if err == nil {
			acmeDirectory.NewAccount.Available = true
		}
	}

	if directoryJSONModel.NewOrder != "" {
		acmeDirectory.NewOrder.URL, err = url.Parse(directoryJSONModel.NewOrder)
		if err == nil {
			acmeDirectory.NewOrder.Available = true
		}
	}

	if directoryJSONModel.NewAuthz != "" {
		acmeDirectory.NewAuthz.URL, err = url.Parse(directoryJSONModel.NewAuthz)
		if err == nil {
			acmeDirectory.NewAuthz.Available = true
		}
	}

	if directoryJSONModel.RevokeCert != "" {
		acmeDirectory.RevokeCert.URL, err = url.Parse(directoryJSONModel.RevokeCert)
		if err == nil {
			acmeDirectory.RevokeCert.Available = true
		}
	}

	if directoryJSONModel.KeyChange != "" {
		acmeDirectory.KeyChange.URL, err = url.Parse(directoryJSONModel.KeyChange)
		if err == nil {
			acmeDirectory.KeyChange.Available = true
		}
	}

	return acmeDirectory, nil
}
