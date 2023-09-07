package bitbucket

type Variable struct {
	Object
	Uuid    string `json:"uuid,omitempty"`
	Key     string `json:"key"`
	Value   string `json:"value"`
	Secured bool   `json:"secured"`
}

type PipelinesApiGroup struct {
	c *Client
}

func (p *PipelinesApiGroup) GetVariableForWorkspace(workspace, variableUuid string) (*Variable, error) {
	o := RequestOptions{
		Method: "GET",
		Path:   p.c.requestPath("/workspaces/%s/pipelines-config/variables/%s", workspace, variableUuid),
	}
	var variable Variable
	err := p.c.execute(&o, &variable)
	return &variable, err
}

func (p *PipelinesApiGroup) ListVariablesForEnvironment(workspace, repoSlug, environmentUuid string) ([]Variable, error) {
	o := RequestOptions{
		Method:      "GET",
		Path:        p.c.requestPath("/repositories/%s/%s/deployments_config/environments/%s/variables", workspace, repoSlug, environmentUuid),
		IsPageable:  true,
		CurrentPage: 1,
	}
	var typers []Typer
	err := p.c.executePageable(&o, &typers)
	if err != nil {
		return nil, err
	}
	var variables []Variable
	for _, typer := range typers {
		variables = append(variables, typer.(Variable))
	}

	return variables, err
}

func (p *PipelinesApiGroup) CreateVariableForEnvironment(workspace, repoSlug, environmentUuid string, variable Variable) (*Variable, error) {
	o := RequestOptions{
		Method: "POST",
		Path:   p.c.requestPath("/repositories/%s/%s/deployments_config/environments/%s/variables", workspace, repoSlug, environmentUuid),
		Data:   variable,
	}
	var result Variable

	err := p.c.execute(&o, &result)
	return &result, err
}
