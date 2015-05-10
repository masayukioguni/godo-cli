package command

import (
	"fmt"
	"github.com/digitalocean/godo"
)

func GetNetworksV4IPAddress(networks *godo.Networks) string {
	var ips []string
	for _, v4Network := range networks.V4 {
		ips = append(ips, v4Network.IPAddress)
	}
	return fmt.Sprintf("%v", ips)
}
