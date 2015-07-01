package netAPI

import (
	"fmt"
	"strconv"
	"testing"
)

func TestListDirectConnectGroups(t *testing.T) {
	httpServer, netAPI := NewMockServerAndClient(200, "{\"response\": {\"directconnectgroups\": [{\"id\": \"1\", \"name\": \"dcg1\", \"sids\":[\"network1-sid\",\"network2-sid\"], \"networks\": [\"net1\",\"net2\"]}], \"count\": 1}}")
	defer httpServer.Close()

	p := netAPI.DirectConnectGroup.NewListDirectConnectGroupsParams()
	r, err := netAPI.DirectConnectGroup.ListDirectConnectGroups(p)
	if err != nil {
		fmt.Printf("Error getting list DCGs: %s", err)
	}

	dcgList := r.ListDirectConnectGroups
	if dcgList == nil {
		t.Error("JSON unmarshalling failed")
	}
	checkStringAttributeValue("id", dcgList[0].Id, "1", t)
	checkStringAttributeValue("name", dcgList[0].Name, "dcg1", t)
	checkStringAttributeValue("sid[0]", dcgList[0].Sids[0], "network1-sid", t)
	checkStringAttributeValue("sid[1]", dcgList[0].Sids[1], "network2-sid", t)
	checkStringAttributeValue("networks[0]", dcgList[0].Networks[0], "net1", t)
	checkStringAttributeValue("networks[1]", dcgList[0].Networks[1], "net2", t)
}

func TestCreateDirectConnectGroup(t *testing.T) {
	httpServer, netAPI := NewMockServerAndClient(200, "{\"response\": {\"directconnectgroup\": [{\"id\": 1, \"name\": \"dcg1\"}], \"count\": 1}}")
	defer httpServer.Close()

	p := netAPI.DirectConnectGroup.NewCreateDirectConnectGroupParams("dcg1")
	r, err := netAPI.DirectConnectGroup.CreateDirectConnectGroup(p)
	if err != nil {
		fmt.Printf("Error creating DCG: %s", err)
	}

	dcgList := r.ListDirectConnectGroups
	if dcgList == nil {
		t.Error("JSON unmarshalling failed")
	}
	checkStringAttributeValue("id", strconv.FormatInt(dcgList[0].Id, 10), "1", t)
	checkStringAttributeValue("name", dcgList[0].Name, "dcg1", t)
}

func TestRenameDirectConnectGroup(t *testing.T) {
	httpServer, netAPI := NewMockServerAndClient(200, "{\"response\": {\"directconnectgroup\": [{\"id\": \"1\", \"name\": \"dcg-renamed\"}], \"count\": 1}}")
	defer httpServer.Close()

	p := netAPI.DirectConnectGroup.NewUpdateDirectConnectGroupParams("dcg-orig", "dcg-renamed")
	r, err := netAPI.DirectConnectGroup.UpdateDirectConnectGroup(p)
	if err != nil {
		fmt.Printf("Error updating DCG: %s", err)
	}

	dcgList := r.ListDirectConnectGroups
	if dcgList == nil {
		t.Error("JSON unmarshalling failed")
	}
	checkStringAttributeValue("id", dcgList[0].Id, "1", t)
	checkStringAttributeValue("name", dcgList[0].Name, "dcg-renamed", t)

}
