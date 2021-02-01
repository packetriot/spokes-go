# spokes-go

Go client library for spokes server.   Please check out our [docs](https://docs.packetriot.com/spokes_api/) for more details on endpoints and data descriptions.  This client library and the Admin API is being actively developed.  Please report any bugs or issues to [support](mailto:support@packetriot.com) or [DM](https://twitter.com/packetriot) us on Twitter.

# Requirements

You'll need the following to use this client library:

* Spokes setup on a server or virtual machine
* [License](https://packetriot.com/enterprise)

# Basic Usage

The Admin API for Spokes uses a key to authenticate a web session.  This is performed only once until the client is logged out or session expires.

Below is a following snippet to authenticate:
```

client, err := spokes.NewClientWithURL("https://spokes.example.com")
...
if _, err := client.Auth("INSERT KEY HERE"); err != nil {
	log.Fatal("authentication failed...")
}
```

Commonly, you may want to poll and check which tunnels are online:
```
tuns, err := client.OnlineTunnels()
if err != nil {
	log.Fatal("...")
}

for _, tun := range tuns {
	// Print out the host IP address for the tunnel
	fmt.Printf("Tun-Hostname: %s, hosted at %s\n", tun.Hostname, tun.Address)
}
```

You may want to shutdown a tunnel.  This will terminate the tunnel session and immediately drop all traffic between client and the tunnel.  

```

var tun *spokes.Tunnel
...

if resp, err := client.ShutdownTunnel(tun.ID); err != nil {
	log.Fataln(err.Error())
} else if resp.Status {
	fmt.Println("Tunnel was shut down")
}
```


