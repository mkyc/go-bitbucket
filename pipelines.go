package bitbucket

type Variable struct {
	Object
	Uuid    string `json:"uuid,omitempty"`
	Key     string `json:"key"`
	Value   string `json:"value"`
	Secured bool   `json:"secured"`
}

type VariablesPage struct {
	Page
	Values []Variable `json:"values"`
}

func (p *VariablesPage) GetValues() []Typer {
	var t []Typer
	for _, v := range p.Values {
		t = append(t, v)
	}
	return t
}

func (p *VariablesPage) GetSize() int {
	return p.Size
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
		HasNextPage: true,
		CurrentPage: 1,
	}

	var pages []VariablesPage
	for o.HasNextPage {
		var page VariablesPage
		err := p.c.execute(&o, &page)
		if err != nil {
			return nil, err
		}
		pages = append(pages, page)

		if o.CurrentPage*p.c.pagination.PageLength >= page.GetSize() {
			o.HasNextPage = false
		}
		o.CurrentPage++
	}

	var variables []Variable

	for _, page := range pages {
		for _, v := range page.Values {
			variables = append(variables, v)
		}
	}

	return variables, nil
}

func (p *PipelinesApiGroup) CreateVariableForEnvironment(workspace, repoSlug, environmentUuid string, variable Variable) (*Variable, error) {
	o := RequestOptions{
		Method: "POST",
		Path:   p.c.requestPath("/repositories/%s/%s/deployments_config/environments/%s/variables", workspace, repoSlug, environmentUuid),
		Data:   variable,
	}
	var createdVariable Variable

	err := p.c.execute(&o, &createdVariable)
	return &createdVariable, err
}

func (p *PipelinesApiGroup) UpdateVariableForEnvironment(workspace, repoSlug, environmentUuid, variableUuid string, variable Variable) (*Variable, error) {
	o := RequestOptions{
		Method: "PUT",
		Path:   p.c.requestPath("/repositories/%s/%s/deployments_config/environments/%s/variables/%s", workspace, repoSlug, environmentUuid, variableUuid),
		Data:   variable,
	}
	var result Variable

	err := p.c.execute(&o, &result)
	return &result, err
}
