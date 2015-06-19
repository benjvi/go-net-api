package netAPI

import (
	"fmt"
	"testing"
	"os"
)

func TestListNetworks(t *testing.T) {
	// TODO: read in auth info from environment variables
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
	// TODO: read in auth info from environment variables
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
	// TODO: read in auth info from environment variables
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
