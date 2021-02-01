package spokes

import (
	"time"
)

var (
	SystemID UID // all zeros
)

const (
	ScopeUser         uint16 = iota + 1 // used by pktriot for basic users
	ScopeAdmin                          // user by admins
	ScopeServer                         // used by servers
	ScopeRegistration                   // used by clients to get auth tokens

	defaultKeyLength int64 = 64
)

type APIKey struct {
	ID          UID       `json:"id"`
	Created     time.Time `json:"created"`
	Modified    time.Time `json:"modified"`
	Active      bool      `json:"active"`
	Description string    `json:"description"`
	Value       string    `json:"value"`
}

func (a *APIKey) Generate() {
	// Generate random value returned as hex string N-length hex characters.
	a.Value = randomString(defaultKeyLength)
}

func (a *APIKey) KeySnippet() string {
	if len(a.Value) == 0 {
		return "..."
	}
	return a.Value[:8]
}

func (a *APIKey) KeySnippetN(n int) string {
	if len(a.Value) == 0 {
		return "..."
	} else if n > len(a.Value) {
		n = len(a.Value)
	}
	return a.Value[:n]
}
