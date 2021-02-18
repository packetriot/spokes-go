package spokes

import (
	"time"
)

type User struct {
	ID           UID       `json:"id"`
	Created      time.Time `json:"created"`
	Modified     time.Time `json:"modified"`
	Active       bool      `json:"active"`
	Fullname     string    `json:"fullname"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	MaxBandwidth int       `json:"maxBandwidth"` // monthly data cap (MB)

	// Runtime field...
	Domains *DomainLimits `json:"-"`
}

type DomainLimits struct {
	Allowed []string `json:"allowed"`
}
