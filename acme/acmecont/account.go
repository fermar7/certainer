package acmecont

import (
	"time"

	"github.com/fermar7/certainer/acme/infra"
)

// Account represents an ACME account
type Account struct {
	ID        int              `json:"id"`
	Status    string           `json:"status"`
	Contact   []string         `json:"contact"`
	Orders    string           `json:"orders"`
	Key       infra.JSONWebKey `json:"key"`
	InitialIP string           `json:"initialIp"`
	CreatedAt time.Time        `json:"createdAt"`
}
