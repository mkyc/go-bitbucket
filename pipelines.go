package bitbucket

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Variable struct {
	Object
	Uuid    string `json:"uuid"`
	Key     string `json:"key"`
	Value   string `json:"value"`
	Secured bool   `json:"secured"`
}

type VariablesPage struct {
	Page
	Values []Variable `json:"values"`
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
	err := p.c.execute(o, &variable)
	return &variable, err
}

func (p *PipelinesApiGroup) ListVariablesForEnvironment(workspace, repoSlug, environmentUuid string) ([]Variable, error) {
	o := RequestOptions{
		Method: "GET",
		Path:   p.c.requestPath("/repositories/%s/%s/deployments_config/environments/%s/variables", workspace, repoSlug, environmentUuid),
	}
	req, err := p.c.newRequest(o)
	if err != nil {
		return nil, err
	}

	currentPage := 1
	hasNextPage := true
	var variables []Variable
	for hasNextPage {
		variablesPage, size, err := p.listVariablesForEnvironmentPage(req, currentPage)
		if err != nil {
			return nil, err
		}
		variables = append(variables, variablesPage...)
		if currentPage*p.c.pagination.PageLength >= size {
			hasNextPage = false
		}
		currentPage++
	}
	return variables, nil
}

func (p *PipelinesApiGroup) listVariablesForEnvironmentPage(req *http.Request, page int) ([]Variable, int, error) {
	q := req.URL.Query()
	q.Set("page", strconv.Itoa(page))
	req.URL.RawQuery = q.Encode()

	bodyBytes, err := p.c.do(req)
	if err != nil {
		return nil, -1, err
	}
	if p.c.Debug {
		p.c.logPrettyBody(bodyBytes)
	}
	var variablesPage VariablesPage
	err = json.Unmarshal(bodyBytes, &variablesPage)
	return variablesPage.Values, variablesPage.Size, err
}
