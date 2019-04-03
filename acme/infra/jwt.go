package infra

import (
	"crypto/rsa"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// JSONWebToken ...
type JSONWebToken struct {
	Protected string `json:"protected,omitempty"`
	Payload   string `json:"payload,omitempty"`
	Signature string `json:"signature,omitempty"`
}

// CreateJWT ...
func CreateJWT(header JWTHeader, payload JWTPayload, key *rsa.PrivateKey) (JSONWebToken, error) {
	err := header.Validate()
	if err != nil {
		return JSONWebToken{}, fmt.Errorf("Invalid JWT header: %s", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, payload.ToClaims())
	token.Header = header

	singedString, err := token.SignedString(key)
	if err != nil {
		return JSONWebToken{}, fmt.Errorf("Error signing JWT: %s", err)
	}

	splittetJwt := strings.Split(singedString, ".")

	if len(splittetJwt) != 3 {
		return JSONWebToken{}, fmt.Errorf("Generated JWT consists not of 3 parts")
	}

	jwtModel := JSONWebToken{
		Protected: splittetJwt[0],
		Payload:   splittetJwt[1],
		Signature: splittetJwt[2],
	}

	return jwtModel, nil
}
