package netAPI

import (
	"fmt"
	"testing"
)

// Checking all the network attributes we care about are unmarshalled correctly
func TestListNetworks(t *testing.T) {
	httpServer, netAPI := NewMockServerAndClient(200, "{\"response\": {\"network\": [{\"id\": \"1\", \"name\": \"network1\", \"displaytext\": \"network1-text\", \"cidr\": \"192.168.34.0/24\", \"dcgfriendlyname\":\"mydcg\", \"gateway\": \"192.168.34.1\"}], \"count\": 1}}")
        defer httpServer.Close()

	p := netAPI.Network.NewListNetworksParams()
	r, err := netAPI.Network.ListNetworks(p)
	if err != nil {
		t.Error(fmt.Sprintf("Error getting list networks: %s\n", err))
	}

	networks := r.Networks
	if networks == nil {
		t.Error("Error unmarshalling JSON into list\n")
	}
	// Only testing the fields we use with NetAPI resources
	checkStringAttributeValue("name", networks[0].Name, "network1", t)
	checkStringAttributeValue("id", networks[0].Id, "1", t)
	checkStringAttributeValue("displaytext", networks[0].Displaytext, "network1-text", t)
	checkStringAttributeValue("cidr", networks[0].Cidr, "192.168.34.0/26", t)
	checkStringAttributeValue("dcgfriendlyname", networks[0].Dcgfriendlyname, "mydcg", t)
	checkStringAttributeValue("gateway", networks[0].Gateway, "192.168.34.1", t)
}

// testing query for IPACs - duplicated unmarshalling test
func TestListIPACs(t *testing.T) {
	httpServer, netAPI := NewMockServerAndClient(200, "{\"response\": {\"network\": [{\"id\": \"1\", \"name\": \"ipac1\"}], \"count\": 1}}")
        defer httpServer.Close()

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
	checkStringAttributeValue("name", networks[0].Name, "ipac1", t)
	checkStringAttributeValue("id", networks[0].Id, "1", t)
}

// testing query for VPNs - duplicated unmarsjhalling test
func TestListVPNs(t *testing.T) {
	httpServer, netAPI := NewMockServerAndClient(200, "{\"response\":{\"network\": [{\"id\": \"1\", \"name\": \"vpn1\"}], \"count\": 1}}") 
        defer httpServer.Close()

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
	checkStringAttributeValue("name", networks[0].Name, "vpn1", t)
	checkStringAttributeValue("id", networks[0].Id, "1", t)
}

func TestRenameVPNs(t *testing.T) {
	httpServer, netAPI := NewMockServerAndClient(200, "{\"response\": {\"network\": [{\"id\": \"1\", \"displaytext\": \"renamed-vpn\"}], \"count\": 1}}")
        defer httpServer.Close()

	p := netAPI.Network.NewModifyNetworkParams("1")
	p.SetDisplaytext("renamed-vpn")

	r, err := netAPI.Network.ModifyNetwork(p)
	if err != nil {
		t.Error(fmt.Sprintf("Error modifying networks: %s\n", err))
	}

	networks := r.ListNetworks
	if networks == nil {
		t.Error("Error unmarshalling JSON into list\n")
	}
	for _, netw := range networks {
		checkStringAttributeValue("id", netw.Id, "1", t)
		checkStringAttributeValue("displaytext", netw.Displaytext, "renamed-vpn", t)
	}
}

func TestChangeVPNGatewayAndCidr(t *testing.T) {
	httpServer, netAPI := NewMockServerAndClient(200, "{\"response\": {\"network\": [{\"id\": \"2\", \"originalid\": \"1\", \"cidr\": \"10.0.8.0/24\", \"gateway\": \"10.0.28.1\"}], \"count\": 1}}")
        defer httpServer.Close()

	p := netAPI.Network.NewModifyNetworkParams("1")
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
		checkStringAttributeValue("originalid", netw.Originalid, "1", t)
		checkStringAttributeValue("id", netw.Id, "2", t)
		checkStringAttributeValue("gateway", netw.Gateway, "10.0.28.1", t)
		checkStringAttributeValue("cidr", netw.Cidr, "10.0.28.0/24", t)
	}
}
