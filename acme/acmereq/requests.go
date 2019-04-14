package acmereq

import (
	"crypto/rsa"
	"fmt"
	"net/http"

	"github.com/fermar7/certainer/acme/acmecont"
	"github.com/fermar7/certainer/acme/infra"
)

// CreateNewNonceRequest ...
func CreateNewNonceRequest(directory *infra.AcmeDirectory) (*AcmeRequest, error) {

	if !directory.NewNonce.Available {
		return nil, fmt.Errorf("Ressource 'new-nonce' not available")
	}

	request := &AcmeRequest{
		URL:                  directory.NewNonce.URL,
		Method:               http.MethodHead,
		ExpectedResponseCode: http.StatusOK,
	}

	return request, nil
}

// CreateAccountCreationRequest ...
func CreateAccountCreationRequest(directory *infra.AcmeDirectory, accountKey *rsa.PrivateKey, contact []string) (*AcmeRequest, error) {

	if !directory.NewAccount.Available {
		return nil, fmt.Errorf("Ressource 'new-acct' not available")
	}

	nonce, err := getNonce(directory)
	if err != nil {
		return nil, err
	}

	payload := acmecont.AccountCreatePayload{
		Contact:              contact,
		TermsOfServiceAgreed: true,
	}

	jwtHeader := infra.NewHeader(nonce, directory.NewAccount.URL.String(), infra.WithJWK(infra.GetJWK(accountKey)))

	jwt, err := infra.CreateJWT(jwtHeader, payload, accountKey)
	if err != nil {
		return nil, err
	}

	request := &AcmeRequest{
		URL:                  directory.NewAccount.URL,
		Method:               http.MethodPost,
		ExpectedResponseCode: http.StatusCreated,
		Body:                 &jwt,
	}

	return request, nil
}

// CreateNewOrderRequest ...
func CreateNewOrderRequest(directory *infra.AcmeDirectory, accountKey *rsa.PrivateKey, kid string, domains []string) (*AcmeRequest, error) {

	if !directory.NewOrder.Available {
		return nil, fmt.Errorf("Ressource 'new-order' not available")
	}

	nonce, err := getNonce(directory)
	if err != nil {
		return nil, err
	}

	identifiers := make([]acmecont.Identifier, len(domains))

	for i, domain := range domains {
		identifiers[i] = acmecont.Identifier{Type: "dns", Value: domain}
	}

	payload := acmecont.OrderCreatePayload{
		Identifiers: identifiers,
		NotAfter:    nil,
		NotBefore:   nil,
	}

	jwtHeader := infra.NewHeader(nonce, directory.NewOrder.URL.String(), infra.WithKID(kid))

	jwt, err := infra.CreateJWT(jwtHeader, payload, accountKey)
	if err != nil {
		return nil, err
	}

	request := &AcmeRequest{
		URL:                  directory.NewOrder.URL,
		Method:               http.MethodPost,
		ExpectedResponseCode: http.StatusCreated,
		Body:                 &jwt,
	}

	return request, nil
}

func getNonce(directory *infra.AcmeDirectory) (string, error) {
	newNonceReq, err := CreateNewNonceRequest(directory)
	if err != nil {
		return "", err
	}

	newNonceRes, err := Execute(newNonceReq)
	if err != nil {
		return "", err
	}

	return newNonceRes.Header.Get("Replay-Nonce"), nil
}
