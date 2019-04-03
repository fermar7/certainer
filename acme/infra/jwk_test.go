package infra

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"strings"
	"testing"
)

func TestGetJWK(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 2048)

	jwk := GetJWK(key)

	if jwk.Kty != "RSA" {
		t.Errorf("KTY = %s, want 'RSA'", jwk.Kty)
	}
}

func TestGetThumbprint(t *testing.T) {
	const expected string = "NzbLsXh8uDCcd-6MNwXF4W_7noWXFZAfHkxZsRGC9Xs"

	jwk := JSONWebKey{
		E:   "AQAB",
		Kty: "RSA",
		N:   "0vx7agoebGcQSuuPiLJXZptN9nndrQmbXEps2aiAFbWhM78LhWx4cbbfAAtVT86zwu1RK7aPFFxuhDR1L6tSoc_BJECPebWKRXjBZCiFV4n3oknjhMstn64tZ_2W-5JsGY4Hc5n9yBXArwl93lqt7_RN5w6Cf0h4QyQ5v-65YGjQR0_FDW2QvzqY368QQMicAtaSqzs8KJZgnYb9c7d0zgdAZHzu6qMQvRL5hajrn1n91CbOpbISD08qNLyrdkt-bFTWhAI4vMQFh6WeZu0fM4lFd2NcRwr3XPksINHaQ-G_xBniIqbw0Ls1jF44-csFCur-kEgU8awapJzKnqDKgw",
	}

	generated, err := jwk.GetThumbprint()

	if err != nil {
		t.Errorf("err should be nil, got: %s", err)
	}

	base64string := base64.URLEncoding.EncodeToString(generated)

	base64string = strings.TrimSuffix(base64string, "=")

	if base64string != expected {
		t.Errorf("Generated string '%s' not matching expected '%s'", base64string, expected)
	}
}
