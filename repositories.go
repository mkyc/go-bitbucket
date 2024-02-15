package bitbucket

import "fmt"

type Project struct {
	Object
	Key  string `json:"key"`
	Uuid string `json:"uuid"`
}

type Repository struct {
	Object
	Uuid      string  `json:"uuid"`
	Slug      string  `json:"slug"`
	FullName  string  `json:"full_name"`
	Name      string  `json:"name"`
	IsPrivate bool    `json:"is_private"`
	Project   Project `json:"project"`
}

type RepositoriesPage struct {
	Page
	Values []Repository `json:"values"`
}

func (p *RepositoriesPage) GetValues() []Typer {
	var t []Typer
	for _, v := range p.Values {
		t = append(t, v)
	}
	return t
}

func (p *RepositoriesPage) GetSize() int {
	return p.Size
}

type RepositoriesApiGroup struct {
	c *Client
}

func (r *RepositoriesApiGroup) ListRepositoriesInWorkspace(workspace string) ([]Repository, error) {
	o := RequestOptions{
		Method:      "GET",
		Path:        r.c.requestPath("/repositories/%s", workspace),
		IsPageable:  true,
		HasNextPage: true,
		CurrentPage: 1,
	}
	var pages []RepositoriesPage
	for o.HasNextPage {
		var page RepositoriesPage
		err := r.c.execute(&o, &page)
		if err != nil {
			return nil, err
		}
		pages = append(pages, page)

		if r.c.Debug {
			fmt.Printf("page: %+v\n", page)
			fmt.Printf("Current page: %d\n", o.CurrentPage)
			fmt.Printf("Page length: %d\n", r.c.pagination.PageLength)
			fmt.Printf("Page size: %d\n", page.GetSize())
		}

		if o.CurrentPage*r.c.pagination.PageLength >= page.GetSize() {
			o.HasNextPage = false
		}
		o.CurrentPage++
	}

	var repositories []Repository
	for _, page := range pages {
		for _, v := range page.Values {
			repositories = append(repositories, v)
		}
	}

	return repositories, nil
}
