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
	o := RequestOptions{Method: "GET", Path: "/workspaces/" + name}
	req, err := w.c.prepareRequest(o)
	if err != nil {
		return nil, err
	}
	bytes, err := w.c.do(req)
	if err != nil {
		return nil, err
	}
	var workspace Workspace
	err = json.Unmarshal(bytes, &workspace)
	return &workspace, err
}
