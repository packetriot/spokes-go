package spokes

import (
	"fmt"

	"github.com/packetriot/spokes-go/json"
)

const (
	UpdateHTTPSitePath       = "/api/admin/v1.0/tunnel/traffic/http/update/"
	DeleteHTTPSitePath       = "/api/admin/v1.0/tunnel/traffic/http/remove/"
	AllocatePortPath         = "/api/admin/v1.0/tunnel/traffic/tcp/allocate/"
	ReleasePortPath          = "/api/admin/v1.0/tunnel/traffic/tcp/release/"
	UpdatePortForwardingPath = "/api/admin/v1.0/tunnel/traffic/tcp/update/"
	ResetPortForwardingPath  = "/api/admin/v1.0/tunnel/traffic/tcp/reset/"
	UpdatePortMappingPath    = "/api/admin/v1.0/tunnel/traffic/portmap/update/"
	RemovePortMappingPath    = "/api/admin/v1.0/tunnel/traffic/portmap/remove/"
)

func (c *Client) UpdateHTTPSite(tunID UID, sites []*Http) (*BasicResponse, error) {
	if sites == nil {
		return nil, fmt.Errorf("invalid (nil) http argument")
	}

	path := UpdateHTTPSitePath + tunID.String()
	response, err := c.request("POST", path, &UpdateHTTPRequest{Sites: sites})
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

func (c *Client) DeleteHTTPSite(tunID UID, domains []string) (*BasicResponse, error) {
	path := DeleteHTTPSitePath + tunID.String()
	response, err := c.request("POST", path, &RemoveHTTPRequest{Domains: domains})
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

func (c *Client) AllocatePort(tunID UID) (*PortResponse, error) {
	path := AllocatePortPath + tunID.String()
	response, err := c.request("GET", path, nil)
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

func (c *Client) ReleasePort(tunID UID, portNo int) (*BasicResponse, error) {
	path := ReleasePortPath + tunID.String()
	response, err := c.request("POST", path, &Port{Port: portNo})
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

func (c *Client) UpdatePortForwarding(tunID UID, ports []*Port) (*BasicResponse, error) {
	if ports == nil {
		return nil, fmt.Errorf("invalid (nil) user argument")
	}

	path := UpdatePortForwardingPath + tunID.String()
	response, err := c.request("POST", path, &UpdatePortRequest{Ports: ports})
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

func (c *Client) ResetPortForwarding(tunID UID, ports []int) (*BasicResponse, error) {
	path := ResetPortForwardingPath + tunID.String()
	response, err := c.request("POST", path, &ResetPortRequest{Ports: ports})
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

func (c *Client) UpdatePortMapping(tunID UID, portmaps []*PortMap) (*BasicResponse, error) {
	if portmaps == nil {
		return nil, fmt.Errorf("invalid (nil) portmap argument")
	}

	path := UpdatePortMappingPath + tunID.String()
	response, err := c.request("POST", path, &UpdatePortMapRequest{PortMappings: portmaps})
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

func (c *Client) RemovePortMapping(tunID UID, listenPorts []int) (*BasicResponse, error) {
	path := RemovePortMappingPath + tunID.String()
	response, err := c.request("POST", path, &RemovePortMapRequest{ListenPorts: listenPorts})
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
