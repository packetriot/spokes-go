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
	Bandwidth    int       `json:"bandwidth"`    // monthly limit (MB)
	MaxBandwidth bool      `json:"maxBandwidth"` // flag indicates max bandwidth limit hit for user
	CheckDomains bool      `json:"checkDomains"` // flag indicates a validation check should be performed on HTTP rules
	MaxPorts     int       `json:"maxPorts"`     // max number of port allocations/rules for user
	CheckPorts   bool      `json:"checkPorts"`   // flag indicates max ports limit should be enforced

	// Runtime field...
	Domains *DomainLimits `json:"-"`
}

type DomainLimits struct {
	Allowed []string `json:"allowed"`
}
