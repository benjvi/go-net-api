package netAPI

import (
	"fmt"
	"os"
	"testing"
)

func TestListDirectConnectGroups(t *testing.T) {
	// TODO: read in auth info from environment variables
	netAPI := NewClient("https://myservices.interoute.com/myservices/api/vdc", os.Getenv("NETAPI_PUBLIC_KEY"), os.Getenv("NETAPI_SECRET_KEY"), false)
	p := netAPI.DirectConnectGroup.NewListDirectConnectGroupsParams()
	r, err := netAPI.DirectConnectGroup.ListDirectConnectGroups(p)
	if err != nil {
		fmt.Printf("Error getting list DCGs: %s", err)
	}

	dcgList := r.ListDirectConnectGroups
	if dcgList == nil {
		t.Error("JSON unmarshalling failed")
	}
	for idx, dcg := range dcgList {
		fmt.Printf("DCG list result %d : %+v", idx+1, dcg)
		if dcg.Name == "" {
			t.Error("Empty name found on DCG. Check raw response and unmarshalling.")
		}
		if dcg.Id == "" {
			t.Error("Empty ID found on DCG. Check raw response and unmarshalling.")
		}
	}
}

/* Cannot delete DCGs so don't create them automatically :(
func TestCreateDirectConnectGroup(t *testing.T) {
	netAPI := NewClient("https://myservices.interoute.com/myservices/api/vdc", os.Getenv("NETAPI_PUBLIC_KEY"), os.Getenv("NETAPI_SECRET_KEY"), false)
	p := netAPI.DirectConnectGroup.NewCreateDirectConnectGroupParams("gotest-4")
	r, err := netAPI.DirectConnectGroup.CreateDirectConnectGroup(p)
	if err != nil {
		fmt.Printf("Error creating DCG: %s", err)
	}

	dcgList := r.ListDirectConnectGroups
        if dcgList == nil {
                t.Error("JSON unmarshalling failed")
        }
        for idx, dcg := range dcgList {
                fmt.Printf("Created DCG %d : %+v", idx+1, dcg)
                if dcg.Name == "" {
                        t.Error("Empty name found on DCG. Check raw response and unmarshalling.")
                }
                if dcg.Id == 0 {
                        t.Error("Empty ID found on DCG. Check raw response and unmarshalling.")
                }
        }
}
*/

func TestUpdateDirectConnectGroup(t *testing.T) {
	netAPI := NewClient("https://myservices.interoute.com/myservices/api/vdc", os.Getenv("NETAPI_PUBLIC_KEY"), os.Getenv("NETAPI_SECRET_KEY"), false)
	p := netAPI.DirectConnectGroup.NewUpdateDirectConnectGroupParams("goupdate-test2", "goupdate-test")
	r, err := netAPI.DirectConnectGroup.UpdateDirectConnectGroup(p)
	if err != nil {
		fmt.Printf("Error updating DCG: %s", err)
	}

	dcgList := r.ListDirectConnectGroups
	if dcgList == nil {
		t.Error("JSON unmarshalling failed")
	}
	for idx, dcg := range dcgList {
		fmt.Printf("Updated DCG %d : %+v", idx+1, dcg)
		if dcg.Name == "" {
			t.Error("Empty name found on DCG. Check raw response and unmarshalling.")
		}
		if dcg.Id == "" {
			t.Error("Empty ID found on DCG. Check raw response and unmarshalling.")
		}
	}

	// Succeeded so we need to put back the original values for the next run
	p = netAPI.DirectConnectGroup.NewUpdateDirectConnectGroupParams("goupdate-test", "goupdate-test2")
	netAPI.DirectConnectGroup.UpdateDirectConnectGroup(p)
}
