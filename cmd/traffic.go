package main

import (
	"fmt"
	"github.com/packetriot/spokes-go"
)

var (
	token string = "TOKEN HERE"
)

func main() {
	client, err := spokes.NewClientWithURL("https://spokes.borak.co")
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

		p := spokes.Port{Port: port.Port, Destination: "127.0.0.1", DstPort: 22}

		if _, err := client.UpdatePortForwarding(tun.ID, []*spokes.Port{&p}); err != nil {
			fmt.Println("Error: " + err.Error())
		}
	}

	// Reset one of the ports
	if _, err := client.ResetPortForwarding(tun.ID, []int{tunConfig.Ports[1].Port}); err != nil {
		fmt.Println("Error: " + err.Error())
	}

	// Create some HTTP rules...
	basename := ".spokes.borak.co"

	client.UpdateHTTPSite(tun.ID, []*spokes.Http{
		{
			Domain:      "api-test-0" + basename,
			Destination: "127.0.0.1",
			Port:        8080,
			WebRoot:     "/tmp",
			Redirect:    false,
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
		}})

	client.UpdateHTTPSite(tun.ID, []*spokes.Http{
		{
			Domain:      "api-test-2" + basename,
			Destination: "127.0.0.1",
			Port:        8080,
			WebRoot:     "/tmp",
			Redirect:    true,
			CA:          certBase64,
			PrivateKey:  privateKeyBase64,
		}})

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
}

var (
	certBase64       = `LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURIekNDQWdlZ0F3SUJBZ0lSQUlIeEI3YWhlUGhwQnlCaFpjRXNCL1l3RFFZSktvWklodmNOQVFFTEJRQXcKQURBZUZ3MHlNakF4TURjeU1UTTRORGRhRncweU16QXhNRGN5TVRNNE5EZGFNQUF3Z2dFaU1BMEdDU3FHU0liMwpEUUVCQVFVQUE0SUJEd0F3Z2dFS0FvSUJBUUN3bkZvSU1xdnVJQ0Nwek01NDNrMTFCQjRuSGdqUDZvSGYrcC9NCm9abW5RUUVBSjBHbStJNEdPaEtMWWFDS3ZNWlUra2duVkNyME9zNFRtVUd1UU1ENmdSY3hFT1NhUUVBY29DQ0kKK1NVME1meG9zT0ZlRHRhVDZ6c0RvNEpZZ2xxQS8rUW9rQS8yTXUxTk0zUFF5VFNrNXZrZnV6MWo1bTNPa0FCcApYOHBUVFVxMVVTeVEwYjRNMndCSmtTNEEwNWlaeGl5NGd0SUE1V3Z5ekVrZGxQcmVOblRkUEovV09TenNzTnYxClNWSmozT2IwZXQrbTRVbTFCbkVQWEJCUFNtdWNpRy9wbittSDloWlowVkxlRWJzcmdFYU96Ymg0Ly9sUjhsbE0KQXJyb2JCMmw4TVpwNlJrQmFkUHRXV3JreTUydXlQcTJQYzFROXpvclpkWXNLQTR2QWdNQkFBR2pnWk13Z1pBdwpEZ1lEVlIwUEFRSC9CQVFEQWdXZ01CTUdBMVVkSlFRTU1Bb0dDQ3NHQVFVRkJ3TUJNQXdHQTFVZEV3RUIvd1FDCk1BQXdIUVlEVlIwT0JCWUVGQzEwVXZwcjJVRnZ5bTFqcS84dW54NHhSSDVuTUI4R0ExVWRJd1FZTUJhQUZGQ2MKanEvY21uZU1CMWp3OWVXcVcrTG5uV09YTUJzR0ExVWRFUUVCL3dRUk1BK0NEWFJsYzNScGJtY3ViRzlqWVd3dwpEUVlKS29aSWh2Y05BUUVMQlFBRGdnRUJBSFhidjAyRkRQQ2RSVEtoWis0MEU0ZU02SDBGMHRUb0o1dStZQ3BnCkUrdmtFaXA2ZERZdUZTMUQ5RVlmeWdkWjJuc0J3SGRPYW1GblBDUEVsT1RzK0V0N3lQZ0FRaE5OU1g0eWovU3MKb1hoV1pTR3ZiZloySXZaenBOd0lJUXgxdzZkRDNaOVd0dUN4WmRISUNKT2VLVjQrWGR5SG9SajFualVNbnp3LwpLVkNhbmpKV1JQWDhibUlzaG9nVHVQNGxFMUdxNGJhanB3bjVkVEpqanJ3UlNsK2llMXhiRGNlb3VtWHBpMWd3CjNlREVCbHY1bTVwR21FeEVRSEcraXBxUGVFZ3Avbno3dTZKeUNJcTFvaEtRdjBPRFdWRFVndXJEaTFyRWJLQzAKcEN4SlQzS0hBMHJPNDArZlV6K0VnaXl3NklReEFsTUQxdXQvYlFGMDljY1dleGc9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K`
	privateKeyBase64 = `LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBc0p4YUNES3I3aUFncWN6T2VONU5kUVFlSng0SXorcUIzL3FmektHWnAwRUJBQ2RCCnB2aU9Cam9TaTJHZ2lyekdWUHBJSjFRcTlEck9FNWxCcmtEQStvRVhNUkRrbWtCQUhLQWdpUGtsTkRIOGFMRGgKWGc3V2srczdBNk9DV0lKYWdQL2tLSkFQOWpMdFRUTnowTWswcE9iNUg3czlZK1p0enBBQWFWL0tVMDFLdFZFcwprTkcrRE5zQVNaRXVBTk9ZbWNZc3VJTFNBT1ZyOHN4SkhaVDYzalowM1R5ZjFqa3M3TERiOVVsU1k5em05SHJmCnB1Rkp0UVp4RDF3UVQwcHJuSWh2NlovcGgvWVdXZEZTM2hHN0s0QkdqczI0ZVAvNVVmSlpUQUs2Nkd3ZHBmREcKYWVrWkFXblQ3VmxxNU11ZHJzajZ0ajNOVVBjNksyWFdMQ2dPTHdJREFRQUJBb0lCQVFDSHNmbnk1b3YzVURRSQpndFg1UkVTYkxlakZBV3lmNDN5YVRRUk93N005TU5MRi9XT1NHTk4yc1ZQaVI4YUVFbnNJZTZ6SW13RE4yZ2pRClFpYVc3aVhYMHl1d2hWdy9zRElTVEczekVBcW55ZWczdi8vSXR2bmplUTlFd25LNThwMzdFNkdJRVBoU245cDQKUUpveHh3WnQxUkdBY2VxL1FTTUpYTmNVSDRkaTdHN1o5bERyZ2VmOURxMmt6UUpUaUdrMUJCUk9xcm03TXBhVQpUR1I4bVlIL0NOeW5lOHorVnl6U1oxZ21sV3d6ZmhyN0FtM1BCcWlUMFhEUnpnM3FVQXlVaXJzbVFSVi93RFppCmNuSVFrSjNUUUNHcGhKNzJYOEpvOEYxUFM1cXhDTnVhV2RvQlBjOGZzQjVjNUlKQVQ4SmVMU2htd1FmRkpzK2EKdlFXdVFucUJBb0dCQU1vbDNiK3dtN01Kd0NqZ202M1lmL3NwUEYyeWZHZUM5N09pOWZ1YlZER3dVb2FzVjB4bQo2K1BMN3RUZWtuNTM0Zy9iOVlWRU8remlUaXFudVFYVHBQZHFHclJpcVh1Q25WMlBEeWJUdlE3Wk1xL3F6RThMCnFJckF0VEJ5bTcwOFVUZENsWkIvTUZ2YjNIV2l0YVUwOGFRYnp0SGVhM1kyTEpSQW5nbTdoc3VQQW9HQkFOK28KNVQzVVQzSDlKSnY5cEhubXhRdXA2NGJXUisvR1N1NW1zelYvUU9lanplMnVTNm9keUxyRWpiWGdqNytId0pDZwpOTEk5WGg0TEhJeDNYTUxXV3ZWZzZQMmYya3ZQdTJ3S1lDRWxNUERWaitqa053NUpRd0t6bDdtS3IrVkdwcVZXClV0eGh5anh6NW8wN2JJaWs2UjhzL3BSbWxWbVFlNzJSK0VtbkNjTmhBb0dBYmVjK2M0eWlhaW0vOXJsL0NucHQKd25DOTZEYzFHZFlEcy8vQ2V3UU5OMktreTZHQmFRRi8rSDZVbjlsT2prSEJmMXZZVlpjdWVYRGtqSjVab3NoWgpwVVpqdEhUN3JqSHFPc2FmdHoxaFNXUmZBWFBIbHFaQkFRY2F1M2RrSXZOYk9JOFQzOWEyeFFwNUJ0L3FvQ0p6CnlnUndZbnZwc3dCckprTW5hU1V5ZVJNQ2dZQURPTHVVbUdwTWlnanM5a3BZTnlxL2NFTWtQaEhyTWtBQ0R4aWwKdkorQ3RxbFFzeUlENFVueTVzSGp6TWhGU2Y5TUZnS3NUcFg4ZU15QWVYZXNsc25ZbnZ5OEtvRktka09NMnRsbgpvTkdEcG8vY0d1MXdGejRQMitaamxjdlMwYUcwMy9seGt6Y0doK1RhUS9EY1J3UFVueUZMb2U2a3k1LzhxdzJZCmdmOXlBUUtCZ1FDUFpkSGdBbm5rSnA1dlhycGc4SitRdkQ3UmNhV3VQKytHNTV4MGhrRHdVUEhob1B5RTZHS3UKd3VrNmUzZjhSaFJWRTQ4YlRSTFpTOEZoNFo2TkQzMS90b0JLSjhkQlhiWHN2SVBnYXM0TDg0VENZSFJzU0xZOApvdDduMmtKOUFUS0N0T3V5V3d5bE5vQzlkeVNydmd0bDVucEZvRUZiQXJmWTk5QUlBMWE2Y2c9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=`
)
