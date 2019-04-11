package acmecont

import "time"

// Authorization ...
type Authorization struct {
	Status     string      `json:"status"`
	Expires    time.Time   `json:"expires"`
	Identifier Identifier  `json:"identifier"`
	Challenges []Challenge `json:"challenges"`
	Wildcard   bool        `json:"wildcard"`
}
