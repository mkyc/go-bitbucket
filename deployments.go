package bitbucket

type EnvironmentTypeName string

const (
	EnvironmentTypeProduction EnvironmentTypeName = "Production"
	EnvironmentTypeStaging    EnvironmentTypeName = "Staging"
	EnvironmentTypeTest       EnvironmentTypeName = "Test"
)

type EnvironmentTypeRank int

const (
	EnvironmentTypeRankProduction EnvironmentTypeRank = 2
	EnvironmentTypeRankStaging    EnvironmentTypeRank = 1
	EnvironmentTypeRankTest       EnvironmentTypeRank = 0
)

type EnvironmentType struct {
	Object
	Name EnvironmentTypeName `json:"name"`
	Rank EnvironmentTypeRank `json:"rank"`
}

type Environment struct {
	Object
	Uuid            string          `json:"uuid,omitempty"`
	Name            string          `json:"name"`
	Slug            string          `json:"slug,omitempty"`
	EnvironmentType EnvironmentType `json:"environment_type"`
}

type EnvironmentsPage struct {
	Page
	Values []Environment `json:"values"`
}

func (p *EnvironmentsPage) GetValues() []Typer {
	var t []Typer
	for _, v := range p.Values {
		t = append(t, v)
	}
	return t
}

func (p *EnvironmentsPage) GetSize() int {
	return p.Size
}

type DeploymentsApiGroup struct {
	c *Client
}

func (d *DeploymentsApiGroup) ListEnvironments(workspace, repoSlug string) ([]Environment, error) {
	o := RequestOptions{
		Method:      "GET",
		Path:        d.c.requestPath("/repositories/%s/%s/environments", workspace, repoSlug),
		IsPageable:  true,
		HasNextPage: true,
		CurrentPage: 1,
	}
	var pages []EnvironmentsPage
	for o.HasNextPage {
		var page EnvironmentsPage
		err := d.c.execute(&o, &page)
		if err != nil {
			return nil, err
		}
		pages = append(pages, page)

		if o.CurrentPage*d.c.pagination.PageLength >= page.GetSize() {
			o.HasNextPage = false
		}
		o.CurrentPage++
	}

	var environments []Environment
	for _, page := range pages {
		for _, v := range page.Values {
			environments = append(environments, v)
		}
	}
	return environments, nil
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

func (d *DeploymentsApiGroup) CreateEnvironment(workspace, repoSlug string, environment Environment) (*Environment, error) {
	environment.EnvironmentType.Type = "deployment_environment_type"
	o := RequestOptions{
		Method: "POST",
		Path:   d.c.requestPath("/repositories/%s/%s/environments", workspace, repoSlug),
		Data:   environment,
	}
	var createdEnvironment Environment

	err := d.c.execute(&o, &createdEnvironment)
	return &createdEnvironment, err
}

func (d *DeploymentsApiGroup) DeleteEnvironment(workspace, repoSlug, environmentUuid string) error {
	o := RequestOptions{
		Method: "DELETE",
		Path:   d.c.requestPath("/repositories/%s/%s/environments/%s", workspace, repoSlug, environmentUuid),
	}
	return d.c.execute(&o, nil)
}
