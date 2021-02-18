package spokes

import (
	"time"
)

type BasicResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

type TunResponse struct {
	Status bool   `json:"status"`
	Error  string `json:"error"`

	// List of tunnels (all, active, online)
	Tunnels []*Tunnel `json:"tunnels,omitempty"`

	// Detailed single tunnel response
	Tunnel *Tunnel `json:"tunnel,omitempty"`

	// Authentication Token (used during creation
	Token *APIKey `json:"token,omitempty"`
}

type PortResponse struct {
	Status bool       `json:"status"`
	Error  string     `json:"error"`
	Ports  []*PortTun `json:"ports"`
}

type PortTun struct {
	Port  int `json:"port"`
	TunID UID `json:"tunID"`
}

type PortRangeResponse struct {
	Status bool      `json:"status"`
	Error  string    `json:"error"`
	Range  PortRange `json:"range"`
}

type PortRange struct {
	Begin int `json:"begin"`
	End   int `json:"end"`
}

type TokenResponse struct {
	Status bool   `json:"status"`
	Error  string `json:"error"`

	// List token operations (registration, auth)
	Tokens []*APIKey `json:"tokens,omitempty"`

	// Create (registration, authtoken token response
	Token *APIKey `json:"token,omitempty"`
}

type TunConfigResponse struct {
	Status bool       `json:"status"`
	Error  string     `json:"error"`
	Config *TunConfig `json:"config"`
}

// Supports pktriot client HTTP API.  Can be used as a request or a response.
type TunConfig struct {
	Version string  `json:"version,omitempty"`
	OS      string  `json:"os,omitempty"`
	Arch    string  `json:"arch,omitempty"`
	Https   []*Http `json:"https,omitempty"`
	Ports   []*Port `json:"ports,omitempty"`
}

type Http struct {
	Domain        string `json:"domain,omitempty"`        // domain request, e.g. example.com
	Secure        bool   `json:"secure,omitempty"`        // indicates http (80) and/or https (443)
	Destination   string `json:"destination,omitempty"`   // forward to this host/address
	Port          int    `json:"port,omitempty"`          // port to forward on
	TLS           int    `json:"tls,omitempty"`           // port to forward to for TLS
	WebRoot       string `json:"webRoot,omitempty"`       // static document root to serve content
	UseLetsEnc    bool   `json:"useLetsEnc,omitempty"`    // use lets-encrypt for TLS certificates
	CA            string `json:"ca,omitempty"`            // path to custom certificate authority
	PrivateKey    string `json:"privateKey,omitempty"`    // path to custom privacy key
	Redirect      bool   `json:"redirect,omitempty"`      // redirect to https
	Password      string `json:"password,omitempty"`      // salted-hash of password to permit traffic
	BasicHTTPAuth string `json:"basicHttpAuth,omitempty"` // salted-hash of user:password (HTTP basic auth) to permit traffic
	RedirectURL   string `json:"redirectURL,omitempty"`   // redirect all requests to URL
	RewriteHost   string `json:"rewriteHost,omitempty"`   // modify host header with this value
}

type Port struct {
	Port        int    `json:"port,omitempty"`        // port used by servers, e.g. 22001 (for ssh)
	Destination string `json:"destination,omitempty"` // forward to this host/address
	DstPort     int    `json:"dstPort,omitempty"`
}

type DomainResponse struct {
	Status  bool     `json:"status"`
	Error   string   `json:"error"`
	Domains []string `json:"domains"`
}

type UserResponse struct {
	Status bool   `json:"status"`
	Error  string `json:"error"`

	// List of all (basic) users on the system
	Users []*User `json:"users,omitempty"`

	// User object for creation, import
	User *User `json:"user,omitempty"`
}

type LicenseInfoResponse struct {
	Status     bool      `json:"status"`
	Error      string    `json:"error"`
	MaxTunnels int       `json:"maxTunnels"`
	Expires    time.Time `json:"expires"`
}
