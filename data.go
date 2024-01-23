package spokes

import (
	"fmt"
	"math"
	"time"
)

const (
	TunInit uint16 = iota + 1
	TunOnline
	TunOffline
	TunShutdown
	TunDeleted
)

const (
	TCP uint16 = iota + 1
	UDP
)

const (
	// Values for ConnMetric.Service
	ServiceHttp uint16 = iota + 1
	ServicePort
)

var (
	minute = (60.0 * 60.0)
	day    = (60.0 * 60.0 * 24.0)
)

type Tunnel struct {
	ID         UID           `json:"id"`
	UserID     UID           `json:"userID"`
	Created    time.Time     `json:"created"`
	LastActive time.Time     `json:"lastActive"`
	State      uint16        `json:"state"`
	Uptime     time.Duration `json:"uptime"` // value is saved in seconds
	Bandwidth  DataUsage     `json:"bandwidth"`
	Name       string        `json:"name"`
	Hostname   string        `json:"hostname"` // domain assigned to tunnel
	Address    string        `json:"address"`  // ip address of client
	Version    string        `json:"version"`  // client version (most recent session)
	OS         string        `json:"os"`       // operating system
	Arch       string        `json:"arch"`     // architecture

	// Used during runtime
	Https []*HttpService `json:"https,omitempty"`
	Ports []*PortService `json:"ports,omitempty"`
}

func (t *Tunnel) IsVisible() bool {
	return (t.State >= TunInit && t.State <= TunShutdown)
}

func (t *Tunnel) IsActive() bool {
	return (t.State >= TunInit && t.State < TunShutdown)
}

func (t *Tunnel) IsOnline() bool {
	return (t.State == TunOnline)
}

func (t *Tunnel) NumServices() int {
	numActive := 0
	for _, hs := range t.Https {
		if hs.Active {
			numActive++
		}
	}
	for _, ps := range t.Ports {
		if ps.Active {
			numActive++
		}
	}

	return numActive
}

func (t *Tunnel) StateString() string {
	switch t.State {
	case TunInit:
		return "Init"
	case TunOnline:
		return "Running"
	case TunOffline:
		return "Stopped"
	case TunShutdown:
		return "Shutdown"
	case TunDeleted:
		return "Deleted"
	}

	return "Unknown"
}

func (t *Tunnel) UptimeString() string {
	// Need to return this value back nanoseconds
	duration := (t.Uptime * time.Second)
	uptime := float64(t.Uptime)

	if uptime < 60.0 {
		return fmt.Sprintf("%d secs", t.Uptime)
	} else if uptime < minute {
		// Minutes
		return fmt.Sprintf("%.1f min", duration.Minutes())
	} else if uptime < day {
		// Hours
		return fmt.Sprintf("%.1f hrs", duration.Hours())
	}

	// Days
	return fmt.Sprintf("%.1f days", (duration.Hours() / 24.0))
}

func (t *Tunnel) DomainNames() (domains []string) {
	domainMap := make(map[string]bool)
	for _, http := range t.Https {
		domainMap[http.DomainName] = true
	}

	for domainName := range domainMap {
		domains = append(domains, domainName)
	}

	return domains
}

func (t *Tunnel) DataUsage() int64 {
	var usage int64
	for _, hs := range t.Https {
		usage += int64(hs.Bandwidth.Monthly)
	}
	for _, ps := range t.Ports {
		usage += int64(ps.Bandwidth.Monthly)
	}

	return usage
}

type HttpService struct {
	ID         UID       `json:"id"`
	UserID     UID       `json:"userID"`
	TunID      UID       `json:"tunID"`
	Active     bool      `json:"active"`
	Available  bool      `json:"available"`
	Secure     bool      `json:"secure"` // true == https, false == http
	DomainName string    `json:"domainName"`
	Bandwidth  DataUsage `json:"bandwidth"`
}

type PortService struct {
	ID        UID       `json:"id"`
	UserID    UID       `json:"userID"`
	TunID     UID       `json:"tunID"`
	Active    bool      `json:"active"`
	Available bool      `json:"available"`
	Protocol  uint16    `json:"protocol"`  // tcp/udp
	Port      int       `json:"port"`      // port used by servers, e.g. 22001 (for ssh)
	Bandwidth DataUsage `json:"bandwidth"` // bandwidth stats
	Label     string    `json:"label"`     // e.g. ssh, smtp, docker
}

type DataUsage struct {
	Daily   ByteSize `json:"daily"`   // bytes
	Monthly ByteSize `json:"monthly"` // bytes
}

type ByteSize int64

func (bs ByteSize) String() string {
	size, unit := bs.Units()
	return fmt.Sprintf("%.02f %s", size, unit)
}

func (bs ByteSize) Units() (float64, string) {

	value := float64(bs)

	kb := float64(math.Pow10(3))
	mb := float64(math.Pow10(6))
	gb := float64(math.Pow10(9))
	tb := float64(math.Pow10(12))

	if value < kb {
		return value, "B"
	} else if value < mb {
		return (value / kb), "KB"
	} else if value < gb {
		return (value / mb), "MB"
	} else if value < tb {
		return (value / gb), "GB"
	}

	return (value / tb), "TB"
}

type ConnMetric struct {
	ID          UID       `json:"id"`
	TunnelID    UID       `json:"tunID"`
	ServiceID   UID       `json:"serviceID"` // unique ID for http/port service
	Service     uint16    `json:"service"`   // indicates the service type
	Address     string    `json:"address"`   // ip address of client
	Bandwidth   int64     `json:"bandwidth"` // read + write
	Established time.Time `json:"established"`
	Closed      time.Time `json:"closed"`

	// Runtime field
	ServerName string `json:"-"`
	Port       uint16 `json:"-"`
}

func (cm *ConnMetric) DataUsage() string {
	bs := ByteSize(cm.Bandwidth)
	return bs.String()
}

// TunService is used to associate a tunnel with an HTTP *or* a PortService
// and an availability state for the service. This struct is used to communicate
// when a service goes up/down.
type TunService struct {
	Tun       *Tunnel      `json:"tunnel"`
	HTTP      *HttpService `json:"http,omitempty"`
	Port      *PortService `json:"port,omitempty"`
	Available bool         `json:"available"`
	Timestamp time.Time    `json:"timestamp"`
}
