package netAPI

/*
import (
	"fmt"
	"testing"
	"os"
)

 no automated creations because we cant delete :(
func TestCreatePublicDirectConnect(t *testing.T) {
	netAPI := NewClient("https://myservices.interoute.com/myservices/api/vdc", os.Getenv("NETAPI_PUBLIC_KEY"), os.Getenv("NETAPI_SECRET_KEY"), false)
	p := netAPI.PublicDirectConnect.NewCreatePublicDirectConnectParams("goipactest", "Slough", 29)
	r, err := netAPI.PublicDirectConnect.CreatePublicDirectConnect(p)
	if err != nil {
		fmt.Printf("Error creating : %s", err)
	}

	pdcList := r.ListPublicDirectConnects
	if pdcList == nil {
		t.Error("JSON unmarshalling failed")
	}
	for idx, pdc := range pdcList {
		fmt.Printf("Created PDC %d : %+v", idx+1, pdc)
		checkStringAttribute("gateway", pdc.Gateway, t)
		checkStringAttribute("sid", pdc.Sid, t)
		checkStringAttribute("zonename", pdc.Zonename, t)
		checkStringAttribute("zoneid", pdc.Zoneid, t)
		checkStringAttribute("domainid", pdc.Domainid, t)
		checkStringAttribute("cidr", pdc.Cidr, t)
		checkStringAttribute("id", pdc.Id, t)
		checkStringAttribute("vlan", pdc.Vlan, t)
		checkStringAttribute("routerendpoint1", pdc.Routerendpoint1, t)
		checkStringAttribute("routerendpoint2", pdc.Routerendpoint2, t)
	}
}
*/
