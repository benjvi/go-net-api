package netAPI

import (
	"encoding/json"
	"net/url"
)

type CreateDirectConnectGroupParams struct {
	p map[string]interface{}
}

func (p *CreateDirectConnectGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}

	return u
}

func (p *CreateDirectConnectGroupParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

// You should always use this function to get a new params instance,
// as then you are sure you have configured all required params
func (s *DirectConnectGroupService) NewCreateDirectConnectGroupParams(name string) *CreateDirectConnectGroupParams {
	p := &CreateDirectConnectGroupParams{}
	p.p = make(map[string]interface{})
	p.SetName(name)
	return p
}

// lists direct connect groups
func (s *DirectConnectGroupService) CreateDirectConnectGroup(p *CreateDirectConnectGroupParams) (*CreateDirectConnectGroupResponse, error) {
	resp, err := s.cs.newRequest("createDirectConnectGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateDirectConnectGroupResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateDirectConnectGroupResponse struct {
	Count                   int                              `json:"count"`
	ListDirectConnectGroups []*CreateDirectConnectGroupEntry `json:"directconnectgroup"`
}

type CreateDirectConnectGroupEntry struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type ListDirectConnectGroupsParams struct {
	p map[string]interface{}
}

func (p *ListDirectConnectGroupsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *ListDirectConnectGroupsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListDirectConnectGroupsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new params instance,
// as then you are sure you have configured all required params
func (s *DirectConnectGroupService) NewListDirectConnectGroupsParams() *ListDirectConnectGroupsParams {
	p := &ListDirectConnectGroupsParams{}
	p.p = make(map[string]interface{})
	return p
}

// lists direct connect groups
func (s *DirectConnectGroupService) ListDirectConnectGroups(p *ListDirectConnectGroupsParams) (*ListDirectConnectGroupsResponse, error) {
	resp, err := s.cs.newRequest("listDirectConnectGroups", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDirectConnectGroupsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListDirectConnectGroupsResponse struct {
	Count                   int                             `json:"count"`
	ListDirectConnectGroups []*ListDirectConnectGroupsEntry `json:"directconnectgroups"`
}

type ListDirectConnectGroupsEntry struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Sids     []string `json:"sids"`
	Networks []string `json:"networks"`
}

type UpdateDirectConnectGroupParams struct {
	p map[string]interface{}
}

func (p *UpdateDirectConnectGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["newname"]; found {
		u.Set("newname", v.(string))
	}
	return u
}

func (p *UpdateDirectConnectGroupParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *UpdateDirectConnectGroupParams) SetNewname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["newname"] = v
	return
}

// You should always use this function to get a new params instance,
// as then you are sure you have configured all required params
func (s *DirectConnectGroupService) NewUpdateDirectConnectGroupParams(name, newName string) *UpdateDirectConnectGroupParams {
	p := &UpdateDirectConnectGroupParams{}
	p.p = make(map[string]interface{})
	p.SetName(name)
	p.SetNewname(newName)
	return p
}

// lists direct connect groups
func (s *DirectConnectGroupService) UpdateDirectConnectGroup(p *UpdateDirectConnectGroupParams) (*UpdateDirectConnectGroupResponse, error) {
	resp, err := s.cs.newRequest("updateDirectConnectGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateDirectConnectGroupResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateDirectConnectGroupResponse struct {
	Count                   int                              `json:"count"`
	ListDirectConnectGroups []*UpdateDirectConnectGroupEntry `json:"directconnectgroup"`
}

type UpdateDirectConnectGroupEntry struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
