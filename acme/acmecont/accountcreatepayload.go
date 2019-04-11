package acmecont

import (
	"github.com/dgrijalva/jwt-go"
)

// AccountCreatePayload represents the payload to request a new acme account
type AccountCreatePayload struct {
	TermsOfServiceAgreed bool
	Contact              []string
}

// ToClaims converts the payload to jwt compatible claims
func (payload AccountCreatePayload) ToClaims() jwt.Claims {
	return jwt.MapClaims{
		"termsOfServiceAgreed": payload.TermsOfServiceAgreed,
		"contact":              payload.Contact,
	}
}
