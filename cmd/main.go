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

	fmt.Println("Create Token")
	tokenResp, _ := client.CreateToken("this tunnel is a test", false)

	fmt.Println("Edit Token")
	client.EditToken(tokenResp.Token.ID, tokenResp.Token.Description, true)

	fmt.Println("List Token")
	client.Tokens()

	fmt.Println("Delete Token")
	client.DeleteToken(tokenResp.Token.ID)

	fmt.Println("Active Tunnels")
	client.ActiveTunnels(false)

	fmt.Println("Online Tunnels")
	tunResp, _ := client.OnlineTunnels(false)
	if len(tunResp.Tunnels) > 0 {
		fmt.Println("Tunnel Detail")
		client.Tunnel(tunResp.Tunnels[0].ID)

		fmt.Println("Tunnel Configuration")
		client.TunnelConfig(tunResp.Tunnels[0].ID)
	}

	fmt.Println("All Ports")
	client.Ports()

	fmt.Println("Port Range")
	client.PortRange()

	fmt.Println("Getting *ALL* Tunnels")
	if resp, _ := client.Tunnels(); resp != nil {
		fmt.Printf("Number of pages: %d\n", len(resp.Links))
		for i := 0; i < len(resp.Links); i++ {

			pageResp, err := client.GetTunPageResult(resp.Links[i].UID)
			if err != nil {
				fmt.Println("Error: " + err.Error())
			} else {
				fmt.Printf("\tNum tunnels in page %d - %d\n", resp.Links[i].Order, len(pageResp.Tunnels))
			}
		}
	}

	fmt.Println("Searching Tunnels")
	if resp, _ := client.SearchTunnels(&spokes.SearchTunRequest{OS: "windows"}); resp != nil {
		fmt.Printf("Number of pages: %d\n", len(resp.Links))
		for i := 0; i < len(resp.Links); i++ {

			pageResp, err := client.GetTunPageResult(resp.Links[i].UID)
			if err != nil {
				fmt.Println("Error: " + err.Error())
			} else {
				fmt.Printf("\tNum tunnels in page %d - %d\n", resp.Links[i].Order, len(pageResp.Tunnels))
			}
		}
	}
}
