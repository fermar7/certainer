package infra

import (
	"github.com/dgrijalva/jwt-go"
)

// JWTPayload acts as the base interface for all jwt payload contracts
type JWTPayload interface {
	ToClaims() jwt.Claims
}
