package acmecont

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// OrderCreatePayload ...
type OrderCreatePayload struct {
	Identifiers []Identifier
	NotBefore   *time.Time
	NotAfter    *time.Time
}

// ToClaims ...
func (payload OrderCreatePayload) ToClaims() jwt.Claims {
	claims := jwt.MapClaims{
		"identifiers": payload.Identifiers,
	}

	if payload.NotBefore != nil {
		claims["notBefore"] = payload.NotBefore
	}

	if payload.NotAfter != nil {
		claims["notAfter"] = payload.NotAfter
	}

	return claims
}
