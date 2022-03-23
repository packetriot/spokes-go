package spokes

import (
	"time"
)

type BasicResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

type TunResponse struct {
	Status bool   `json:"status"`
	Error  string `json:"error"`

	// List of tunnels (all, active, online)
	Total   int            `json:"total,omitempty"`
	Tunnels []*Tunnel      `json:"tunnels,omitempty"`
	Links   []*TunRespLink `json:"links,omitempty"`

	// Detailed single tunnel response
	Tunnel *Tunnel `json:"tunnel,omitempty"`

	// Authentication Token (used during creation
	Token *APIKey `json:"token,omitempty"`
}

type TunRespLink struct {
	UID   string `json:"uid"`
	Order int    `json:"order"`
	Count int    `json:"count"`
	URL   string `json:"url"`
}

type PortResponse struct {
	Status bool       `json:"status"`
	Error  string     `json:"error,omitempty"`
	Ports  []*PortTun `json:"ports"`
}

type PortTun struct {
	Port  int `json:"port"`
	TunID UID `json:"tunID"`
}

type PortRangeResponse struct {
	Status bool      `json:"status"`
	Error  string    `json:"error,omitempty"`
	Range  PortRange `json:"range"`
}

type PortRange struct {
	Begin int `json:"begin"`
	End   int `json:"end"`
}

type TokenResponse struct {
	Status bool   `json:"status"`
	Error  string `json:"error,omitempty"`

	// List token operations (registration, auth)
	Tokens []*APIKey `json:"tokens,omitempty"`

	// Create (registration, authtoken token response
	Token *APIKey `json:"token,omitempty"`
}

type TunConfigResponse struct {
	Status bool       `json:"status"`
	Error  string     `json:"error,omitempty"`
	Config *TunConfig `json:"config"`
}

// Supports pktriot client HTTP API.  Can be used as a request or a response.
type TunConfig struct {
	Version      string     `json:"version,omitempty"`
	OS           string     `json:"os,omitempty"`
	Arch         string     `json:"arch,omitempty"`
	Https        []*Http    `json:"https,omitempty"`
	Ports        []*Port    `json:"ports,omitempty"`
	PortMappings []*PortMap `json:"portmaps,omitempty"`
}

type Http struct {
	Domain           string `json:"domain,omitempty"`           // domain request, e.g. example.com
	Secure           bool   `json:"secure,omitempty"`           // indicates http (80) and/or https (443)
	Destination      string `json:"destination,omitempty"`      // forward to this host/address
	Port             int    `json:"port,omitempty"`             // port to forward on
	TLS              int    `json:"tls,omitempty"`              // port to forward to for TLS
	UpstreamURL      string `json:"upstreamURL,omitempty"`      // upstream service addressed w/URL, e.g. http://127.0.0.1:8080
	WebRoot          string `json:"webRoot,omitempty"`          // static document root to serve content
	UseLetsEnc       bool   `json:"useLetsEnc,omitempty"`       // use lets-encrypt for TLS certificates
	CA               string `json:"ca,omitempty"`               // path to custom certificate authority
	PrivateKey       string `json:"privateKey,omitempty"`       // path to custom privacy key
	Redirect         bool   `json:"redirect,omitempty"`         // redirect to https
	Password         string `json:"password,omitempty"`         // salted-hash of password to permit traffic
	Requires2FA      bool   `json:"requires2FA,omitempty"`      // indicates 2FA is used for authentication
	BasicHTTPAuth    string `json:"basicHttpAuth,omitempty"`    // salted-hash of user:password (HTTP basic auth) to permit traffic
	RedirectURL      string `json:"redirectURL,omitempty"`      // redirect all requests to URL
	RewriteHost      string `json:"rewriteHost,omitempty"`      // modify host header with this value
	InsecureUpstream bool   `json:"insecureUpstream,omitempty"` // Accept insecure TLS upstream servers
}

type Port struct {
	Port        int    `json:"port,omitempty"`        // port used by servers, e.g. 22001 (for ssh)
	Destination string `json:"destination,omitempty"` // forward to this host/address
	DstPort     int    `json:"dstPort,omitempty"`
}

type PortMap struct {
	ListenPort  int    `json:"listenPort"`
	DstPort     int    `json:"dstPort"`
	Destination string `json:"destination"`     // hostname, IP address
	Transport   string `json:"transport"`       // tcp, udp
	Label       string `json:"label,omitempty"` // e.g. ssh, smtp, docker
	HTTP        bool   `json:"http,omitempty"`  // flag indicates http traffic
}

type DomainResponse struct {
	Status  bool     `json:"status"`
	Error   string   `json:"error,omitempty"`
	Domains []string `json:"domains"`
}

type UserResponse struct {
	Status bool   `json:"status"`
	Error  string `json:"error,omitempty"`

	// List of all (basic) users on the system
	Users []*User `json:"users,omitempty"`

	// User object for creation, import
	User *User `json:"user,omitempty"`
}

type LicenseInfoResponse struct {
	Status     bool      `json:"status"`
	Error      string    `json:"error,omitempty"`
	MaxTunnels int       `json:"maxTunnels"`
	Expires    time.Time `json:"expires"`
}
