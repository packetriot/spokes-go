package spokes

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"

	"golang.org/x/net/publicsuffix"

	"github.com/packetriot/spokes-go/json"
)

const (
	BasePath           = "/api/admin/"
	AuthPath           = "/api/admin/v1.0/auth"
	InfoPath           = "/api/admin/v1.0/info"
	PingPath           = "/api/admin/v1.0/ping"
	KeepAlivePath      = "/api/admin/v1.0/keepalive"
	LogoutPath         = "/api/admin/v1.0/logout"
	LicenseInfoPath    = "/api/admin/v1.0/license/info"
	ListTokensPath     = "/api/admin/v1.0/token/registration/list"
	CreateTokenPath    = "/api/admin/v1.0/token/registration/create"
	EditTokenPath      = "/api/admin/v1.0/token/registration/edit"
	DeleteTokenPath    = "/api/admin/v1.0/token/registration/delete"
	ListTunsPath       = "/api/admin/v1.0/tunnel/list"
	ListActiveTunsPath = "/api/admin/v1.0/tunnel/list/active"
	ListOnlineTunsPath = "/api/admin/v1.0/tunnel/list/online"
	SearchTunsPath     = "/api/admin/v1.0/tunnel/search"
	TunPagePath        = "/api/admin/v1.0/tunnel/page/"
	GetTunInfoPath     = "/api/admin/v1.0/tunnel/info/"
	GetTunAuthPath     = "/api/admin/v1.0/tunnel/auth/"
	GetTunConfigPath   = "/api/admin/v1.0/tunnel/config/"
	ShutdownTunPath    = "/api/admin/v1.0/tunnel/shutdown/"
	GetPortsPath       = "/api/admin/v1.0/port/list"
	GetPortRangePath   = "/api/admin/v1.0/port/range"
)

var (
	Debug = false
)

type Client struct {
	client   *http.Client
	address  string
	insecure bool // default is false
}

func NewClientWithURL(address string) (*Client, error) {
	if len(address) == 0 {
		return nil, fmt.Errorf("address is invalid")
	}

	c := &Client{
		client:  &http.Client{},
		address: address,
	}

	if _, err := url.Parse(address); err != nil {
		return nil, err
	}

	// Setup the cookie jar for clients so authenticated clients don't need
	// to manage cookies and sessions themselves.
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, err
	}
	c.client.Jar = jar

	return c, nil
}

// Passing a 'true' value will put this client in insecure-mode and accept
// any TLS certificate it receives from the server.  Passing in false will
// perform proper verification of server and client-side certificates.
func (c *Client) SetInsecure(b bool) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: b},
	}
	c.client.Transport = tr

	// Set this field value so it can used in the PoolConnect() function later.
	c.insecure = b
	return nil
}

func (c *Client) Auth(key string) (*BasicResponse, error) {
	response, err := c.request("POST", AuthPath, &AuthRequest{Key: key})
	if err == nil {
		ar := &BasicResponse{}
		if err = json.Decode(response.Body, ar); err == nil {
			response.Body.Close()
			if ar.Status {
				// Debug
				dumpPrettyJson(ar)
				return ar, nil
			} else {
				return ar, fmt.Errorf(ar.Error)
			}
		}
	}

	return nil, err
}

func (c *Client) Info() (*ServerInfoResponse, error) {
	response, err := c.request("GET", InfoPath, nil)
	if err == nil {
		sir := &ServerInfoResponse{}
		if err = json.Decode(response.Body, sir); err == nil {
			response.Body.Close()
			if sir.Status {
				// Debug
				dumpPrettyJson(sir)
				return sir, nil
			} else {
				return nil, fmt.Errorf(sir.Error)
			}
		}
	}

	return nil, err
}

func (c *Client) Ping() (*BasicResponse, error) {
	response, err := c.request("GET", PingPath, nil)
	if err == nil {
		br := &BasicResponse{}
		if err = json.Decode(response.Body, br); err == nil {
			response.Body.Close()
			if br.Status {
				// Debug
				dumpPrettyJson(br)
				return br, nil
			} else {
				return nil, fmt.Errorf(br.Error)
			}
		}
	}

	return nil, err
}

func (c *Client) KeepAlive() (*BasicResponse, error) {
	response, err := c.request("GET", KeepAlivePath, nil)
	if err == nil {
		br := &BasicResponse{}
		if err = json.Decode(response.Body, br); err == nil {
			response.Body.Close()
			if br.Status {
				// Debug
				dumpPrettyJson(br)
				return br, nil
			} else {
				return nil, fmt.Errorf(br.Error)
			}
		}
	}

	return nil, err
}

func (c *Client) Logout() (*BasicResponse, error) {
	response, err := c.request("GET", LogoutPath, nil)
	if err == nil {
		br := &BasicResponse{}
		if err = json.Decode(response.Body, br); err == nil {
			response.Body.Close()
			if br.Status {
				// Debug
				dumpPrettyJson(br)
				return br, nil
			} else {
				return nil, fmt.Errorf(br.Error)
			}
		}
	}

	return nil, err
}

func (c *Client) License() (*LicenseInfoResponse, error) {
	response, err := c.request("GET", LicenseInfoPath, nil)
	if err == nil {
		lir := &LicenseInfoResponse{}
		if err = json.Decode(response.Body, lir); err == nil {
			response.Body.Close()
			if lir.Status {
				// Debug
				dumpPrettyJson(lir)
				return lir, nil
			} else {
				return nil, fmt.Errorf(lir.Error)
			}
		}
	}

	return nil, err
}

func (c *Client) Tokens() (*TokenResponse, error) {
	response, err := c.request("GET", ListTokensPath, nil)
	if err == nil {
		tr := &TokenResponse{}
		if err = json.Decode(response.Body, tr); err == nil {
			response.Body.Close()
			if tr.Status {
				// Debug
				dumpPrettyJson(tr)
				return tr, nil
			} else {
				return nil, fmt.Errorf(tr.Error)
			}
		}
	}

	return nil, err
}

func (c *Client) CreateToken(description string, active bool) (*TokenResponse, error) {
	ctr := CreateTokenRequest{
		Description: description,
		Active:      active,
	}

	response, err := c.request("POST", CreateTokenPath, &ctr)
	if err == nil {
		tr := &TokenResponse{}
		if err = json.Decode(response.Body, tr); err == nil {
			response.Body.Close()
			if tr.Status {
				// Debug
				dumpPrettyJson(tr)
				return tr, nil
			} else {
				dumpPrettyJson(tr)
				return nil, fmt.Errorf(tr.Error)
			}
		}
	}

	return nil, err
}

func (c *Client) EditToken(tokenID UID, description string, active bool) (*TokenResponse, error) {
	etr := EditTokenRequest{
		TokenID:     tokenID,
		Description: description,
		Active:      active,
	}

	response, err := c.request("POST", EditTokenPath, &etr)
	if err == nil {
		tr := &TokenResponse{}
		if err = json.Decode(response.Body, tr); err == nil {
			response.Body.Close()
			if tr.Status {
				// Debug
				dumpPrettyJson(tr)
				return tr, nil
			} else {
				dumpPrettyJson(tr)
				return nil, fmt.Errorf(tr.Error)
			}
		}
	}

	return nil, err
}

func (c *Client) DeleteToken(tokenID UID) (*BasicResponse, error) {
	tr := TokenRequest{TokenID: tokenID}
	response, err := c.request("POST", DeleteTokenPath, &tr)
	if err == nil {
		br := &BasicResponse{}
		if err = json.Decode(response.Body, br); err == nil {
			response.Body.Close()
			if br.Status {
				// Debug
				dumpPrettyJson(br)
				return br, nil
			} else {
				return nil, fmt.Errorf(br.Error)
			}
		}
	}

	return nil, err
}

func (c *Client) Tunnels() (*TunResponse, error) {
	response, err := c.request("GET", ListTunsPath, nil)
	if err == nil {
		tr := &TunResponse{}
		if err = json.Decode(response.Body, tr); err == nil {
			response.Body.Close()
			if tr.Status {
				// Debug
				dumpPrettyJson(tr)
				return tr, nil
			} else {
				return nil, fmt.Errorf(tr.Error)
			}
		}
	}

	return nil, err
}

func (c *Client) ActiveTunnels(details bool) (*TunResponse, error) {
	path := ListActiveTunsPath
	if details {
		path = path + "?details=true"
	}
	response, err := c.request("GET", path, nil)
	if err == nil {
		tr := &TunResponse{}
		if err = json.Decode(response.Body, tr); err == nil {
			response.Body.Close()
			if tr.Status {
				// Debug
				dumpPrettyJson(tr)
				return tr, nil
			} else {
				return nil, fmt.Errorf(tr.Error)
			}
		}
	}

	return nil, err
}

func (c *Client) OnlineTunnels(details bool) (*TunResponse, error) {
	path := ListOnlineTunsPath
	if details {
		path = path + "?details=true"
	}
	response, err := c.request("GET", path, nil)
	if err == nil {
		tr := &TunResponse{}
		if err = json.Decode(response.Body, tr); err == nil {
			response.Body.Close()
			if tr.Status {
				// Debug
				dumpPrettyJson(tr)
				return tr, nil
			} else {
				return nil, fmt.Errorf(tr.Error)
			}
		}
	}

	return nil, err
}

func (c *Client) SearchTunnels(input *SearchTunRequest) (*TunResponse, error) {
	response, err := c.request("POST", SearchTunsPath, input)
	if err == nil {
		tr := &TunResponse{}
		if err = json.Decode(response.Body, tr); err == nil {
			response.Body.Close()
			if tr.Status {
				// Debug
				dumpPrettyJson(tr)
				return tr, nil
			} else {
				return nil, fmt.Errorf(tr.Error)
			}
		}
	}

	return nil, err
}

func (c *Client) GetTunPageResult(uid string) (*TunResponse, error) {
	path := TunPagePath + uid
	response, err := c.request("GET", path, nil)
	if err == nil {
		tr := &TunResponse{}
		if err = json.Decode(response.Body, tr); err == nil {
			response.Body.Close()
			if tr.Status {
				// Debug
				dumpPrettyJson(tr)
				return tr, nil
			} else {
				return nil, fmt.Errorf(tr.Error)
			}
		}
	}

	return nil, err
}

func (c *Client) Tunnel(id UID) (*TunResponse, error) {
	path := GetTunInfoPath + id.String()
	response, err := c.request("GET", path, &TunRequest{TunID: id})
	if err == nil {
		tr := &TunResponse{}
		if err = json.Decode(response.Body, tr); err == nil {
			response.Body.Close()
			if tr.Status {
				// Debug
				dumpPrettyJson(tr)
				return tr, nil
			} else {
				return nil, fmt.Errorf(tr.Error)
			}
		}
	}

	return nil, err
}

func (c *Client) TunnelAuth(id UID) (*TokenResponse, error) {
	path := GetTunAuthPath + id.String()
	response, err := c.request("GET", path, nil)
	if err == nil {
		tr := &TokenResponse{}
		if err = json.Decode(response.Body, tr); err == nil {
			response.Body.Close()
			if tr.Status {
				// Debug
				dumpPrettyJson(tr)
				return tr, nil
			} else {
				return nil, fmt.Errorf(tr.Error)
			}
		}
	}

	return nil, err
}

func (c *Client) TunnelConfig(id UID) (*TunConfigResponse, error) {
	path := GetTunConfigPath + id.String()
	response, err := c.request("GET", path, nil)
	if err == nil {
		tcr := &TunConfigResponse{}
		if err = json.Decode(response.Body, tcr); err == nil {
			response.Body.Close()
			if tcr.Status {
				// Debug
				dumpPrettyJson(tcr)
				return tcr, nil
			} else {
				return nil, fmt.Errorf(tcr.Error)
			}
		}
	}

	return nil, err
}

func (c *Client) ShutdownTunnel(id UID) (*BasicResponse, error) {
	path := ShutdownTunPath + id.String()
	response, err := c.request("GET", path, nil)
	if err == nil {
		br := &BasicResponse{}
		if err = json.Decode(response.Body, br); err == nil {
			response.Body.Close()
			if br.Status {
				// Debug
				dumpPrettyJson(br)
				return br, nil
			} else {
				return nil, fmt.Errorf(br.Error)
			}
		}
	}

	return nil, err
}

func (c *Client) Ports() (*PortResponse, error) {
	response, err := c.request("GET", GetPortsPath, nil)
	if err == nil {
		pr := &PortResponse{}
		if err = json.Decode(response.Body, pr); err == nil {
			response.Body.Close()
			if pr.Status {
				// Debug
				dumpPrettyJson(pr)
				return pr, nil
			} else {
				return nil, fmt.Errorf(pr.Error)
			}
		}
	}

	return nil, err
}

func (c *Client) PortRange() (*PortRangeResponse, error) {
	response, err := c.request("GET", GetPortRangePath, nil)
	if err == nil {
		prr := &PortRangeResponse{}
		if err = json.Decode(response.Body, prr); err == nil {
			response.Body.Close()
			if prr.Status {
				// Debug
				dumpPrettyJson(prr)
				return prr, nil
			} else {
				return nil, fmt.Errorf(prr.Error)
			}
		}
	}

	return nil, err
}

// HTTP client request utility.  Return response and error message.
func (c *Client) request(method, path string, message interface{}) (response *http.Response, err error) {
	body := &bytes.Buffer{}
	if err = json.EncodePretty(body, message); err == nil {
		var request *http.Request

		if Debug {
			fmt.Printf("URL: %s\n\nMessage:\n%s\n", path, body.String())
		}

		request, err = http.NewRequest(method, c.address+path, body)

		if err == nil {
			request.Header.Set("Accept", "application/json")
			request.Header.Set("Content-Type", "application/json")
			response, err = c.client.Do(request)
		} else {
			return nil, err
		}
	}

	// JB: maybe found an odd corner case when server requested is down
	// but an error is not returned and a nil response is returned.  Not
	// sure but this will provide logic to make sure an error is returned
	// when response is nil.
	if response == nil && err == nil {
		err = fmt.Errorf("request to '%s' return nil response", c.address+path)
	}

	return response, err
}

func dumpPrettyJson(v interface{}) {
	if Debug {
		fmt.Println("Response:")
		json.EncodePretty(os.Stdout, v)
		fmt.Println("")
	}
}
