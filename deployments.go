package bitbucket

import (
	"encoding/json"
)

type Environment struct {
	Object
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}

type DeploymentsApiGroup struct {
	c *Client
}

func (d *DeploymentsApiGroup) GetEnvironment(workspace, repoSlug, environmentUuid string) (*Environment, error) {
	o := RequestOptions{Method: "GET", Path: d.c.requestPath("/repositories/%s/%s/environments/%s", workspace, repoSlug, environmentUuid)}
	req, err := d.c.newRequest(o)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := d.c.do(req)
	if err != nil {
		return nil, err
	}
	if d.c.Debug {
		d.c.logPrettyBody(bodyBytes)
	}
	var environment Environment
	err = json.Unmarshal(bodyBytes, &environment)
	return &environment, err
}
