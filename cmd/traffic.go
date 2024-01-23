package main

import (
	"fmt"
	"github.com/packetriot/spokes-go"
)

var (
	token string = "TOKEN HERE"
)

func main() {
	client, err := spokes.NewClientWithURL("https://spokes.example.com")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		resp, err := client.Auth(token)
		if err != nil {
			fmt.Println(err.Error())
			return
		} else {
			fmt.Println(resp.Status)
		}
	}

	fmt.Println("Searching for active test tunnel")
	resp, err := client.SearchTunnels(&spokes.SearchTunRequest{Term: "API traffic test"})
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	} else if len(resp.Tunnels) == 0 {
		fmt.Println("Error: tunnel not found, cannot continue...")
		return
	}

	tun := resp.Tunnels[0]

	// Get the tunnel configuration and delete all of the rules.
	tunConfigResp, err := client.TunnelConfig(tun.ID)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	} else {
		fmt.Printf("Num HTTPS: %d\n", len(tunConfigResp.Config.Https))
		fmt.Printf("Num TCP Forwarding: %d\n", len(tunConfigResp.Config.Ports))
		fmt.Printf("Num Portmaps: %d\n", len(tunConfigResp.Config.PortMappings))
	}

	tunConfig := tunConfigResp.Config

	// Delete the existing rules for HTTP, TCP forwarding and Port Mapping
	for _, http := range tunConfig.Https {
		if _, err := client.DeleteHTTPSite(tun.ID, []string{http.Domain}); err != nil {
			fmt.Println("Error: " + err.Error())
		}
	}

	for _, port := range tunConfig.Ports {
		if _, err := client.ReleasePort(tun.ID, port.Port); err != nil {
			fmt.Println("Error: " + err.Error())
		}
	}

	for _, portmap := range tunConfig.PortMappings {
		if _, err := client.RemovePortMapping(tun.ID, []int{portmap.ListenPort}); err != nil {
			fmt.Println("Error: " + err.Error())
		}
	}

	// Create some new rules..
	client.AllocatePort(tun.ID)
	client.AllocatePort(tun.ID)
	client.AllocatePort(tun.ID)

	tunConfigResp, err = client.TunnelConfig(tun.ID)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	tunConfig = tunConfigResp.Config

	for _, port := range tunConfig.Ports {

		p := spokes.Port{
			Port:        port.Port,
			Destination: "127.0.0.1",
			DstPort:     2200,
			Firewall: spokes.FwRuleList{
				spokes.FwRule{
					Sequence: 1,
					Action:   spokes.FwActionAllow,
					Network:  "1.2.3.4/32",
				},
				spokes.FwRule{
					Sequence: 2,
					Action:   spokes.FwActionAllow,
					Network:  "2600:4041:4468:6f00::1e47/128",
				},
			},
		}

		if _, err := client.UpdatePortForwarding(tun.ID, []*spokes.Port{&p}); err != nil {
			fmt.Println("Error: " + err.Error())
		}
	}

	// Reset one of the ports
	if _, err := client.ResetPortForwarding(tun.ID, []int{tunConfig.Ports[1].Port}); err != nil {
		fmt.Println("Error: " + err.Error())
	}

	// Create some HTTP rules...
	basename := ".spokes.example.com"

	client.UpdateHTTPSite(tun.ID, []*spokes.Http{
		{
			Domain:      "api-test-0" + basename,
			Destination: "127.0.0.1",
			Port:        8080,
			WebRoot:     "/tmp",
			Redirect:    false,
			Firewall: spokes.FwRuleList{
				spokes.FwRule{
					Sequence: 1,
					Action:   spokes.FwActionAllow,
					Network:  "1.2.3.4/32",
				},
				spokes.FwRule{
					Sequence: 2,
					Action:   spokes.FwActionAllow,
					Network:  "2600:4041:4468:6f00::1e47/128",
				},
			},
		}})

	client.UpdateHTTPSite(tun.ID, []*spokes.Http{
		{
			Domain:      "api-test-1" + basename,
			UseLetsEnc:  true,
			Destination: "127.0.0.1",
			Port:        8080,
			WebRoot:     "/tmp",
			Redirect:    true,
			RewriteHost: "api-test-100.spokes.borak.co",
			Firewall: spokes.FwRuleList{
				spokes.FwRule{
					Sequence: 0,
					Action:   spokes.FwActionAllow,
					Network:  "71.172.242.211/32",
				},
				spokes.FwRule{
					Sequence: 0,
					Action:   spokes.FwActionAllow,
					Network:  "2600:4041:4468:6f00::1e47/128",
				},
			},
		}})

	_, err = client.UpdateHTTPSite(tun.ID, []*spokes.Http{
		{
			Domain:      "api-test-2" + basename,
			Destination: "127.0.0.1",
			Port:        8080,
			WebRoot:     "/tmp",
			Redirect:    true,
			CA:          certBase64,
			PrivateKey:  privateKeyBase64,
			Firewall: spokes.FwRuleList{spokes.FwRule{
				Sequence: 1,
				Action:   spokes.FwActionAllow,
				Network:  "1.2.3.4/32",
			}},
		}})

	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	}

	// Create some port mapping rules now..
	client.UpdatePortMapping(tun.ID, []*spokes.PortMap{
		{
			ListenPort:  15001,
			Destination: "",
			Transport:   "tcp",
			DstPort:     tunConfig.Ports[1].Port,
			HTTP:        false,
			Label:       "port-mapping-1",
		}})

	client.UpdatePortMapping(tun.ID, []*spokes.PortMap{
		{
			ListenPort:  15002,
			Destination: "",
			Transport:   "tcp",
			DstPort:     tunConfig.Ports[2].Port,
			HTTP:        false,
			Label:       "port-mapping-2",
		}})

	// Get the tunnel configuration and delete all of the rules.
	if resp, err := client.TunnelConfig(tun.ID); err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	} else if resp != nil {
		for _, http := range resp.Config.Https {
			if http.Domain != "api-test-1"+basename {
				continue
			}

			if len(http.Firewall) == 0 {
				fmt.Println("ERROR: expecting firewall rules but none found...")
			} else {
				fmt.Println("DEBUG: printing firewall rules")
				for _, rule := range http.Firewall {
					fmt.Printf("Seq: %d, Domain: %s, Action: %s, Network: %s\n",
						rule.Sequence, http.Domain, rule.Action, rule.Network)
				}
				fmt.Println("")
			}
		}
	}

}

var (
	certBase64       = `BASE64 encoded certificate (pem encoded)`
	privateKeyBase64 = `BASE64 encoded private key (pem encoded)`
)
