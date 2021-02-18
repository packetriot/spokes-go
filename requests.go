package spokes

type AuthRequest struct {
	Key     string `json:"key"`
	Version string `json:"version,omitempty"`
	OS      string `json:"os,omitempty"`
	Arch    string `json:"arch,omitempty"`
}

type TunRequest struct {
	TunID UID `json:"tunnelID"`
}

type CreateTunRequest struct {
	UserID   UID    `json:"userID"`
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
}

type CreateTokenRequest struct {
	Description string `json:"description"`
	Active      bool   `json:"active"`
}

type EditTokenRequest struct {
	TokenID     UID    `json:"tokenID"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}

type TokenRequest struct {
	TokenID UID `json:"tokenID"`
}

type UserRequest struct {
	UserID       UID    `json:"userID"`
	Email        string `json:"email,omitempty"`
	Fullname     string `json:"fullname,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Active       bool   `json:"active,omitempty`
	MaxBandwidth int    `json:"maxBandwidth"`
}

type BandwidthRequest struct {
	UserID UID `json:"userID"`
	Max    int `json:"max"` // MB, e.g. 1000 == 1 Gigabyte
}

type DomainRequest struct {
	UserID  UID      `json:"userID"`
	Domains []string `json:"domains"`
}
