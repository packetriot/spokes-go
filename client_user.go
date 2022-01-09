package spokes

import (
	"fmt"

	"github.com/packetriot/spokes-go/json"
)

const (
	ListUsersPath    = "/api/admin/v1.0/user/list"
	CreateUserPath   = "/api/admin/v1.0/user/create"
	EditUserPath     = "/api/admin/v1.0/user/edit"
	ImportUserPath   = "/api/admin/v1.0/user/import"
	DeleteUserPath   = "/api/admin/v1.0/user/delete/"
	ActivateUserPath = "/api/admin/v1.0/user/activate"
	CreateTunPath    = "/api/admin/v1.0/user/tunnel/create"
	StopTunPath      = "/api/admin/v1.0/user/tunnel/stop/"
	ListUserTunsPath = "/api/admin/v1.0/user/tunnel/list/"
	BandwidthCapPath = "/api/admin/v1.0/user/bandwidth/cap"
	AllowDomainPath  = "/api/admin/v1.0/user/domain/allow"
	RemoveDomainPath = "/api/admin/v1.0/user/domain/remove"
	ListDomainPath   = "/api/admin/v1.0/user/domain/list/"
	ResetDomainPath  = "/api/admin/v1.0/user/domain/reset/"
)

func (c *Client) ListUsers() (*UserResponse, error) {
	response, err := c.request("GET", ListUsersPath, nil)
	if err == nil {
		ur := &UserResponse{}
		if err = json.Decode(response.Body, ur); err == nil {
			response.Body.Close()
			if ur.Status {
				// Debug
				dumpPrettyJson(ur)
				return ur, nil
			} else {
				return nil, fmt.Errorf(ur.Error)
			}
		}
	}

	return nil, err
}

func (c *Client) CreateUser(user *UserRequest) (*UserResponse, error) {
	if user == nil {
		return nil, fmt.Errorf("invalid (nil) user argument")
	}

	response, err := c.request("POST", CreateUserPath, user)
	if err == nil {
		ur := &UserResponse{}
		if err = json.Decode(response.Body, ur); err == nil {
			response.Body.Close()
			if ur.Status {
				// Debug
				dumpPrettyJson(ur)
				return ur, nil
			} else {
				return nil, fmt.Errorf(ur.Error)
			}
		}
	}

	return nil, err
}

func (c *Client) EditUser(user *UserRequest) (*BasicResponse, error) {
	if user == nil {
		return nil, fmt.Errorf("invalid (nil) user argument")
	} else if user.UserID.IsZero() {
		return nil, fmt.Errorf("invalid ID for user, it's a zero value")
	}

	response, err := c.request("POST", EditUserPath, user)
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

func (c *Client) DeleteUser(userID UID) (*BasicResponse, error) {
	path := DeleteUserPath + userID.String()
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

func (c *Client) ImportUser(user *UserRequest) (*UserResponse, error) {
	if user == nil {
		return nil, fmt.Errorf("invalid (nil) user argument")
	} else if user.UserID.IsZero() {
		return nil, fmt.Errorf("invalid ID for user, it's a zero value")
	}

	response, err := c.request("POST", ImportUserPath, user)
	if err == nil {
		ur := &UserResponse{}
		if err = json.Decode(response.Body, ur); err == nil {
			response.Body.Close()
			if ur.Status {
				// Debug
				dumpPrettyJson(ur)
				return ur, nil
			} else {
				return nil, fmt.Errorf(ur.Error)
			}
		}
	}

	return nil, err
}

func (c *Client) ActivateUser(userID UID, active bool) (*TunResponse, error) {
	eur := &UserRequest{UserID: userID, Active: active}
	response, err := c.request("POST", ActivateUserPath, &eur)
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

func (c *Client) CreateTunnel(userID UID, name, hostname string) (*TunResponse, error) {
	tr := CreateTunRequest{UserID: userID, Name: name, Hostname: hostname}
	response, err := c.request("POST", CreateTunPath, &tr)
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

func (c *Client) StopTunnel(tunID UID) (*BasicResponse, error) {
	path := StopTunPath + tunID.String()
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

func (c *Client) ListUserTunnels(userID UID) (*TunResponse, error) {
	path := ListUserTunsPath + userID.String()
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

func (c *Client) SetBandwidthCap(userID UID, max int) (*BasicResponse, error) {
	response, err := c.request("POST", BandwidthCapPath, &BandwidthRequest{UserID: userID, Max: max})
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

func (c *Client) AllowDomains(userID UID, domains []string) (*BasicResponse, error) {
	rd := &DomainRequest{UserID: userID, Domains: domains}
	response, err := c.request("POST", AllowDomainPath, &rd)
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

func (c *Client) RemoveDomains(userID UID, domains []string) (*BasicResponse, error) {
	rd := &DomainRequest{UserID: userID, Domains: domains}
	response, err := c.request("POST", RemoveDomainPath, &rd)
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

func (c *Client) ListDomains(userID UID) (*DomainResponse, error) {
	path := ListDomainPath + userID.String()
	response, err := c.request("GET", path, nil)
	if err == nil {
		dr := &DomainResponse{}
		if err = json.Decode(response.Body, dr); err == nil {
			response.Body.Close()
			if dr.Status {
				// Debug
				dumpPrettyJson(dr)
				return dr, nil
			} else {
				return nil, fmt.Errorf(dr.Error)
			}
		}
	}

	return nil, err
}

func (c *Client) ResetDomains(userID UID) (*BasicResponse, error) {
	path := ResetDomainPath + userID.String()
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
