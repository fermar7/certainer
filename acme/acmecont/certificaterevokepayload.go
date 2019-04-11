package acmecont

import jwt "github.com/dgrijalva/jwt-go"

// CertificateRevokePayload ...
type CertificateRevokePayload struct {
	Certificate string
	Reason      int16
}

// ToClaims ...
func (payload CertificateRevokePayload) ToClaims() jwt.Claims {
	return jwt.MapClaims{
		"certificate": payload.Certificate,
		"reason":      payload.Reason,
	}
}
