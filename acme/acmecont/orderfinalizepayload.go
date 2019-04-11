package acmecont

import jwt "github.com/dgrijalva/jwt-go"

// OrderFinalizePayload ...
type OrderFinalizePayload struct {
	CSR string
}

// ToClaims ...
func (payload OrderFinalizePayload) ToClaims() jwt.Claims {
	return jwt.MapClaims{
		"csr": payload.CSR,
	}
}
