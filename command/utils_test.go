package command

import (
	"github.com/digitalocean/godo"
	"reflect"
	"testing"
)

func TestCommond_GetNetworksV4IPAddress(t *testing.T) {
	networks := &godo.Networks{
		V4: []godo.NetworkV4{
			{
				IPAddress: "1.1.1.1",
			},
			{
				IPAddress: "2.2.2.2",
			},
		},
	}

	want := "[1.1.1.1 2.2.2.2]"
	res := GetNetworksV4IPAddress(networks)
	if !reflect.DeepEqual(res, want) {
		t.Errorf("GetNetworksV4IPAddress returned %+v, want %+v", res, want)
	}

}
