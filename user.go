package spokes

import (
	"time"
)

type User struct {
	ID       UID       `json:"id"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	Active   bool      `json:"active"`
	Fullname string    `json:"fullname"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`

	// Runtime
	Limits UserLimits `json:"-"`
}

type UserLimits struct {
	Bandwidth    int64        `json:"bandwidth"`    // max bandwidth (MB) for user
	LimitMessage string       `json:"limitMessage"` // any limit-related messages sent to the user
	Domains      DomainLimits `json:"domains"`
}

type DomainLimits struct {
	Allowed []string `json:"allowed"`
}
