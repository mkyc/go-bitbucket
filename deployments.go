package bitbucket

type Environment struct {
	Object
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}

type DeploymentsApiGroup struct {
	c *Client
}

func (d *DeploymentsApiGroup) GetEnvironment(workspace, repoSlug, environmentUuid string) (*Environment, error) {
	o := RequestOptions{
		Method: "GET",
		Path:   d.c.requestPath("/repositories/%s/%s/environments/%s", workspace, repoSlug, environmentUuid),
	}
	var environment Environment
	err := d.c.execute(&o, &environment)
	return &environment, err
}
