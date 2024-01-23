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
	TunID    UID    `json:"tunID"`
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
}

type SearchTunRequest struct {
	Term    string `json:"term"`
	OS      string `json:"os"`
	Arch    string `json:"arch"`
	Version string `json:"version"`
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
	Bandwidth    int    `json:"bandwidth,omitempty"`
	MaxBandwidth bool   `json:"maxBandwidth,omitempty"`
	CheckDomains bool   `json:"checkDomains,omitempty"`
	MaxPorts     int    `json:"maxPorts,omitempty"`
	CheckPorts   bool   `json:"checkPorts,omitempty"`
}

type BandwidthRequest struct {
	UserID UID `json:"userID"`
	Max    int `json:"max"` // MB, e.g. 1000 == 1 Gigabyte
}

type DomainRequest struct {
	UserID  UID      `json:"userID"`
	Domains []string `json:"domains"`
}

type UpdateHTTPRequest struct {
	Sites []*Http `json:"sites"`
}

type RemoveHTTPRequest struct {
	Domains []string `json:"domains"`
}

type UpdatePortRequest struct {
	Ports []*Port `json:"ports"`
}

type ResetPortRequest struct {
	Ports []int `json:"ports"`
}

type UpdatePortMapRequest struct {
	PortMappings []*PortMap `json:"portmaps"`
}

type RemovePortMapRequest struct {
	ListenPorts []int `json:"listenPorts"`
}
