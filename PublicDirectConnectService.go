package netAPI

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type CreatePublicDirectConnectParams struct {
	p map[string]interface{}
}

func (p *CreatePublicDirectConnectParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["zonename"]; found {
		u.Set("zonename", v.(string))
	}
	if v, found := p.p["suffix"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("suffix", vv)
	}
	return u
}

func (p *CreatePublicDirectConnectParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
	return
}

func (p *CreatePublicDirectConnectParams) SetZonename(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zonename"] = v
	return
}

func (p *CreatePublicDirectConnectParams) SetSuffix(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["suffix"] = v
	return
}

// You should always use this function to get a new params instance,
// as then you are sure you have configured all required params
func (s *PublicDirectConnectService) NewCreatePublicDirectConnectParams(displaytext, zonename string, suffix int) *CreatePublicDirectConnectParams {
	p := &CreatePublicDirectConnectParams{}
	p.p = make(map[string]interface{})
	p.SetDisplaytext(displaytext)
	p.SetZonename(zonename)
	p.SetSuffix(suffix)
	return p
}

// lists direct connect groups
func (s *PublicDirectConnectService) CreatePublicDirectConnect(p *CreatePublicDirectConnectParams) (*CreatePublicDirectConnectResponse, error) {
	resp, err := s.cs.newRequest("createPublicDirectConnect", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreatePublicDirectConnectResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreatePublicDirectConnectResponse struct {
	Count                    int                               `json:"count"`
	ListPublicDirectConnects []*CreatePublicDirectConnectEntry `json:"publicdirectconnect"`
}

type CreatePublicDirectConnectEntry struct {
	Id              string `json:"id"`
	Gateway         string `json:"gateway"`
	Sid             string `json:"sid"`
	Zonename        string `json:"zonename"`
	Zoneid          string `json:"zoneid"`
	Domainid        string `json:"domainid"`
	Cidr            string `json:"cidr"`
	Vlan            string `json:"vlan"`
	Routerendpoint1 string `json:"routerendpoint1"`
	Routerendpoint2 string `json:"routerendpoint2"`
}
