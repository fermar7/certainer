package acmereq

import (
	"fmt"
	"net/http"

	"github.com/fermar7/certainer/acme/infra"
)

// CreateNewNonceRequest ...
func CreateNewNonceRequest(directory *infra.AcmeDirectory) (*AcmeRequest, error) {

	if !directory.NewNonce.Available {
		return nil, fmt.Errorf("Ressource 'new-nonce' not available")
	}

	request := &AcmeRequest{
		URL:    directory.NewNonce.URL,
		Method: http.MethodHead,
	}

	return request, nil
}
