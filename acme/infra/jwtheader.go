package infra

import (
	"fmt"
	"net/url"
)

const (
	algRS256 = "RS256"
	keyAlg   = "alg"
	keyNonce = "nonce"
	keyURL   = "url"
	keyJWK   = "jwk"
	keyKID   = "kid"
)

// JWTHeader ...
type JWTHeader map[string]interface{}

// NewHeader ...
func NewHeader(nonce, URL string, keyFunc func(header JWTHeader) JWTHeader) JWTHeader {
	header := JWTHeader{
		keyAlg:   algRS256,
		keyNonce: nonce,
		keyURL:   URL,
	}

	return keyFunc(header)
}

// Validate ...
func (header JWTHeader) Validate() error {
	const err string = "Required field '%s' not conainted"

	alg, algExists := header[keyAlg]
	if !algExists {
		return fmt.Errorf(err, keyAlg)
	}

	if alg != algRS256 {
		return fmt.Errorf("Only alg of type '%s' supported", algRS256)
	}

	_, nonceExists := header[keyNonce]
	if !nonceExists {
		return fmt.Errorf(err, keyNonce)
	}

	URL, urlExists := header[keyURL]
	if !urlExists {
		return fmt.Errorf(err, keyURL)
	}

	if _, err := url.ParseRequestURI(URL.(string)); err != nil {
		return fmt.Errorf("Malformed url")
	}

	_, kidExists := header[keyKID]
	_, jwkExists := header[keyJWK]

	switch {
	case kidExists && jwkExists:
		return fmt.Errorf("'%s' and '%s' are mutually exclusive", keyJWK, keyKID)
	case !kidExists && !jwkExists:
		return fmt.Errorf("Either '%s' or '%s' required", keyJWK, keyKID)
	}

	if !kidExists {
		return fmt.Errorf(err, keyKID)
	}

	if !jwkExists {
		return fmt.Errorf(err, keyJWK)
	}

	return nil
}

// WithKID ...
func WithKID(kid string) func(header JWTHeader) JWTHeader {
	return func(header JWTHeader) JWTHeader {
		header[keyKID] = kid
		return header
	}
}

// WithJWK ...
func WithJWK(jwk JSONWebKey) func(header JWTHeader) JWTHeader {
	return func(header JWTHeader) JWTHeader {
		header[keyJWK] = jwk
		return header
	}
}
