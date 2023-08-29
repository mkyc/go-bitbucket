package bitbucket

import (
	"encoding/json"
)

type WorkspaceLinks struct {
	Avatar Link `json:"avatar"`
	Html   Link `json:"html"`
	Self   Link `json:"self"`
}

type Workspace struct {
	Name      string         `json:"name"`
	Slug      string         `json:"slug"`
	Type      string         `json:"type"`
	Uuid      string         `json:"uuid"`
	IsPrivate bool           `json:"is_private"`
	CreatedOn string         `json:"created_on"`
	Links     WorkspaceLinks `json:"links"`
}

type WorkspacesApiGroup struct {
	c *Client
}

func (w *WorkspacesApiGroup) GetWorkspace(name string) (*Workspace, error) {
	o := RequestOptions{Method: "GET", Path: w.c.requestPath("/workspaces/%s", name)}
	req, err := w.c.newRequest(o)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := w.c.do(req)
	if err != nil {
		return nil, err
	}
	if w.c.Debug {
		w.c.logPrettyBody(bodyBytes)
	}
	var workspace Workspace
	err = json.Unmarshal(bodyBytes, &workspace)
	return &workspace, err
}
