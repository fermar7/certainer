package contract

import jwt "github.com/dgrijalva/jwt-go"

// AccountUpdatePayload represents the payload to update an acme account
type AccountUpdatePayload struct {
	Contact []string
}

// ToClaims converts the payload to jwt compatible claims
func (payload AccountUpdatePayload) ToClaims() jwt.Claims {
	return jwt.MapClaims{
		"contact": payload.Contact,
	}
}
