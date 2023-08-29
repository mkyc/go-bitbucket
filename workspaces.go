package bitbucket

type WorkspaceLinks struct {
	Avatar Link `json:"avatar"`
	Html   Link `json:"html"`
	Self   Link `json:"self"`
}

type Workspace struct {
	Object
	Name      string         `json:"name"`
	Slug      string         `json:"slug"`
	Uuid      string         `json:"uuid"`
	IsPrivate bool           `json:"is_private"`
	CreatedOn string         `json:"created_on"`
	Links     WorkspaceLinks `json:"links"`
}

type WorkspacesApiGroup struct {
	c *Client
}

func (w *WorkspacesApiGroup) GetWorkspace(name string) (*Workspace, error) {
	o := RequestOptions{
		Method: "GET",
		Path:   w.c.requestPath("/workspaces/%s", name),
	}
	var workspace Workspace
	err := w.c.execute(&o, &workspace)
	return &workspace, err
}
