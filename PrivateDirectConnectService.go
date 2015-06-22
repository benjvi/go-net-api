package netAPI

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type CreatePrivateDirectConnectParams struct {
	p map[string]interface{}
}

func (p *CreatePrivateDirectConnectParams) toURLValues() url.Values {
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
	if v, found := p.p["cidr"]; found {
		u.Set("cidr", v.(string))
	}
	if v, found := p.p["gateway"]; found {
		u.Set("gateway", v.(string))
	}
	if v, found := p.p["dcgname"]; found {
		u.Set("dcgname", v.(string))
	}
	if v, found := p.p["dcgid"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("dcgid", vv)
	}
	return u
}

func (p *CreatePrivateDirectConnectParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
	return
}

func (p *CreatePrivateDirectConnectParams) SetZonename(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zonename"] = v
	return
}

func (p *CreatePrivateDirectConnectParams) SetCidr(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidr"] = v
	return
}

func (p *CreatePrivateDirectConnectParams) SetGateway(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gateway"] = v
	return
}

func (p *CreatePrivateDirectConnectParams) SetDcgname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dcgname"] = v
	return
}

func (p *CreatePrivateDirectConnectParams) SetDcgid(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dcgid"] = v
	return
}

// You should always use this function to get a new params instance,
// as then you are sure you have configured all required params
func (s *PrivateDirectConnectService) NewCreatePrivateDirectConnectParams(displaytext, zonename, cidr, gateway string) *CreatePrivateDirectConnectParams {
	p := &CreatePrivateDirectConnectParams{}
	p.p = make(map[string]interface{})
	p.SetDisplaytext(displaytext)
	p.SetZonename(zonename)
	p.SetCidr(cidr)
	p.SetGateway(gateway)
	return p
}

// creates private direct connects
func (s *PrivateDirectConnectService) CreatePrivateDirectConnect(p *CreatePrivateDirectConnectParams) (*CreatePrivateDirectConnectResponse, error) {
	resp, err := s.cs.newRequest("createPrivateDirectConnect", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreatePrivateDirectConnectResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreatePrivateDirectConnectResponse struct {
	Count                     int                                `json:"count"`
	ListPrivateDirectConnects []*CreatePrivateDirectConnectEntry `json:"privatedirectconnect"`
}

type CreatePrivateDirectConnectEntry struct {
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
