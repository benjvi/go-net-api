package netAPI

import (
	"fmt"
	"os"
	"testing"
)

func TestListNetworks(t *testing.T) {
	netAPI := NewClient("https://myservices.interoute.com/myservices/api/vdc", os.Getenv("NETAPI_PUBLIC_KEY"), os.Getenv("NETAPI_SECRET_KEY"), false)
	p := netAPI.Network.NewListNetworksParams()
	r, err := netAPI.Network.ListNetworks(p)
	if err != nil {
		t.Error(fmt.Sprintf("Error getting list networks: %s\n", err))
	}

	networks := r.Networks
	if networks == nil {
		t.Error("Error unmarshalling JSON into list\n")
	}
	for _, netw := range networks {
		// Only testing the fields we're likely to use most
		checkStringAttribute("name", netw.Name, t)
		checkStringAttribute("id", netw.Id, t)
	}
}

func TestListIPACs(t *testing.T) {
	netAPI := NewClient("https://myservices.interoute.com/myservices/api/vdc", os.Getenv("NETAPI_PUBLIC_KEY"), os.Getenv("NETAPI_SECRET_KEY"), false)
	p := netAPI.Network.NewListNetworksParams()
	p.SetSubtype("publicdirectconnect")

	r, err := netAPI.Network.ListNetworks(p)
	if err != nil {
		t.Error(fmt.Sprintf("Error getting list networks: %s\n", err))
	}

	networks := r.Networks
	if networks == nil {
		t.Error("Error unmarshalling JSON into list\n")
	}
	for _, netw := range networks {
		// Only testing the fields we're likely to use most
		checkStringAttribute("name", netw.Name, t)
		checkStringAttribute("id", netw.Id, t)
	}
}

func TestListVPNs(t *testing.T) {
	netAPI := NewClient("https://myservices.interoute.com/myservices/api/vdc", os.Getenv("NETAPI_PUBLIC_KEY"), os.Getenv("NETAPI_SECRET_KEY"), false)
	p := netAPI.Network.NewListNetworksParams()
	p.SetSubtype("privatedirectconnect")

	r, err := netAPI.Network.ListNetworks(p)
	if err != nil {
		t.Error(fmt.Sprintf("Error getting list networks: %s\n", err))
	}

	networks := r.Networks
	if networks == nil {
		t.Error("Error unmarshalling JSON into list\n")
	}
	for _, netw := range networks {
		// Only testing the fields we're likely to use most
		checkStringAttribute("name", netw.Name, t)
		checkStringAttribute("id", netw.Id, t)
	}
}

func TestRenameVPNs(t *testing.T) {
	netAPI := NewClient("https://myservices.interoute.com/myservices/api/vdc", os.Getenv("NETAPI_PUBLIC_KEY"), os.Getenv("NETAPI_SECRET_KEY"), false)
	p := netAPI.Network.NewModifyNetworkParams("53b874b8-88a5-494b-b4b0-00bc1219e970")
	p.SetDisplaytext("terraform-test-vpn-renamed")

	r, err := netAPI.Network.ModifyNetwork(p)
	if err != nil {
		t.Error(fmt.Sprintf("Error modifying networks: %s\n", err))
	}

	networks := r.ListNetworks
	if networks == nil {
		t.Error("Error unmarshalling JSON into list\n")
	}
	for _, netw := range networks {
		checkStringAttribute("id", netw.Id, t)
		checkStringAttribute("displaytext", netw.Displaytext, t)
	}

	// succeeded - so put the original name back
	p = netAPI.Network.NewModifyNetworkParams("53b874b8-88a5-494b-b4b0-00bc1219e970")
        p.SetDisplaytext("terraform-test-vpn")

        r, err = netAPI.Network.ModifyNetwork(p)
        if err != nil {
                t.Error(fmt.Sprintf("Error reverting network to original state: %s\n", err))
        }

}

/* unsafe method - id changes every time we run this
func TestChangeVPNGatewayAndCidr(t *testing.T) {
	netAPI := NewClient("https://myservices.interoute.com/myservices/api/vdc", os.Getenv("NETAPI_PUBLIC_KEY"), os.Getenv("NETAPI_SECRET_KEY"), false)
	p := netAPI.Network.NewModifyNetworkParams("42c1f230-e478-4a89-898a-c48de38d00bc")
	p.SetCidr("10.0.28.0/24")
	p.SetGateway("10.0.28.1")

	r, err := netAPI.Network.ModifyNetwork(p)
	if err != nil {
		t.Error(fmt.Sprintf("Error modifying networks: %s\n", err))
	}

	networks := r.ListNetworks
	if networks == nil {
		t.Error("Error unmarshalling JSON into list\n")
	}
	for _, netw := range networks {
		checkStringAttribute("originalid", netw.Originalid, t)
		checkStringAttribute("id", netw.Id, t)
		checkStringAttribute("gateway", netw.Gateway, t)
		checkStringAttribute("cidr", netw.Cidr, t)
	}
}*/
