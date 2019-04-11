package acmecont

import "time"

// Challenge ...
type Challenge struct {
	URL       string    `json:"url"`
	Type      string    `json:"type"`
	Status    string    `json:"status"`
	Token     string    `json:"token"`
	Validated time.Time `json:"validated"`
}
