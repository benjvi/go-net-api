package netAPI

import (
	"fmt"
	"testing"
)

func TestCreatePublicDirectConnect(t *testing.T) {
	httpServer, netAPI := NewMockServerAndClient(200, "{\"response\": {\"publicdirectconnect\": [{\"id\": \"4d22713a-0709-49e1-aae9-bc965471e442\", \"routerendpoint2\": \"10.0.25.3\", \"zoneid\": \"5343ddc2-919f-4d1b-a8e6-59f91d901f8e\", \"gateway\": \"10.0.25.1\", \"domainid\": \"38175F38-0755-4C3F-9E17-D83B8B42D7CE\", \"vlan\": \"61\", \"routerendpoint1\": \"10.0.25.2\", \"cidr\": \"10.0.25.0/24\", \"name\": \"Network Private Direct Connect Interoute cstore test 1\", \"sid\": \"INT38/IPVPN/pLC0dsDg\", \"zonename\": \"Slough\"}], \"count\": 1}}")
        defer httpServer.Close()

	p := netAPI.PublicDirectConnect.NewCreatePublicDirectConnectParams("goipactest", "Slough", 29)
	r, err := netAPI.PublicDirectConnect.CreatePublicDirectConnect(p)
	if err != nil {
		fmt.Printf("Error creating : %s", err)
	}

	pdcList := r.ListPublicDirectConnects
	if pdcList == nil {
		t.Error("JSON unmarshalling failed")
	}
        checkStringAttributeValue("gateway", pdcList[0].Gateway, "10.0.25.1", t)
        checkStringAttributeValue("sid", pdcList[0].Sid, "INT38/IPVPN/pLC0dsDg", t)
        checkStringAttributeValue("zonename", pdcList[0].Zonename, "Slough", t)
        checkStringAttributeValue("zoneid", pdcList[0].Zoneid, "5343ddc2-919f-4d1b-a8e6-59f91d901f8e", t)
        checkStringAttributeValue("domainid", pdcList[0].Domainid, "38175F38-0755-4C3F-9E17-D83B8B42D7CE", t)
        checkStringAttributeValue("cidr", pdcList[0].Cidr, "10.0.25.0/24", t)
        checkStringAttributeValue("id", pdcList[0].Id, "4d22713a-0709-49e1-aae9-bc965471e442", t)
        checkStringAttributeValue("vlan", pdcList[0].Vlan, "61", t)
        checkStringAttributeValue("routerendpoint1", pdcList[0].Routerendpoint1, "10.0.25.2", t)
        checkStringAttributeValue("routerendpoint2", pdcList[0].Routerendpoint2, "10.0.25.3", t)
}

