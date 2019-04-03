package infra

import (
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
)

// JSONWebKey ...
type JSONWebKey struct {
	Kty string `json:"kty,omitempty"`
	N   string `json:"n,omitempty"`
	E   string `json:"e,omitempty"`
}

// GetJWK ...
func GetJWK(key *rsa.PrivateKey) JSONWebKey {
	const rsaKey string = "RSA"
	return JSONWebKey{
		Kty: rsaKey,
		N:   base64.RawURLEncoding.EncodeToString(key.N.Bytes()),
		E:   base64.RawURLEncoding.EncodeToString(getUInt64Bytes(uint64(key.E))),
	}
}

// GetThumbprint ...
func (jwk *JSONWebKey) GetThumbprint() ([]byte, error) {
	object := map[string]string{
		"e":   jwk.E,
		"kty": jwk.Kty,
		"n":   jwk.N,
	}

	jsonString, err := json.Marshal(object)
	if err != nil {
		return []byte{}, fmt.Errorf("Error marshalling JWK: %s", err)
	}

	hash := sha256.New()

	_, err = hash.Write(jsonString)
	if err != nil {
		return []byte{}, fmt.Errorf("Error hashing JWK: %s", err)
	}

	return hash.Sum(nil), nil
}

// getUInt64Bytes src: https://github.com/lestrrat-go/jwx/blob/master/internal/base64/base64.go
func getUInt64Bytes(v uint64) []byte {
	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, v)

	i := 0
	for ; i < len(data); i++ {
		if data[i] != 0x0 {
			break
		}
	}

	return data[i:]
}
