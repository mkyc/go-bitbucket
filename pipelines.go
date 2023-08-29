package bitbucket

import "encoding/json"

type Variable struct {
	Type    string `json:"type"`
	Uuid    string `json:"uuid"`
	Key     string `json:"key"`
	Value   string `json:"value"`
	Secured bool   `json:"secured"`
}
type PipelinesApiGroup struct {
	c *Client
}

func (p *PipelinesApiGroup) GetVariableForWorkspace(workspace, uuid string) (*Variable, error) {
	o := RequestOptions{Method: "GET", Path: p.c.requestPath("/workspaces/%s/pipelines-config/variables/%s", workspace, uuid)}
	req, err := p.c.newRequest(o)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := p.c.do(req)
	if err != nil {
		return nil, err
	}
	if p.c.Debug {
		p.c.logPrettyBody(bodyBytes)
	}
	var variable Variable
	err = json.Unmarshal(bodyBytes, &variable)
	return &variable, err
}
