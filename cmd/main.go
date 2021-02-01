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
	client.ActiveTunnels()

	fmt.Println("Online Tunnels")
	tunResp, _ := client.OnlineTunnels()
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
}
