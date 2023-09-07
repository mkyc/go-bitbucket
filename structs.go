package bitbucket

import "encoding/json"

type RequestOptions struct {
	Method      string
	Path        string
	IsPageable  bool
	CurrentPage int
	Data        interface{}
}
type Link struct {
	Href string `json:"href"`
}

type Typer interface {
	GetType() string
}
type Object struct {
	Type string `json:"type"`
}

func (o Object) GetType() string {
	return o.Type
}

type Pager interface {
	GetValues() []Typer
	GetSize() int
}

type Page struct {
	Size     int     `json:"size"`
	Page     int     `json:"page"`
	Pagelen  int     `json:"pagelen"`
	Next     string  `json:"next"`
	Previous string  `json:"previous"`
	Values   []Typer `json:"values"`
}

func (p *Page) GetValues() []Typer {
	return p.Values
}

func (p *Page) GetSize() int {
	return p.Size
}

func (p *Page) UnmarshalJSON(data []byte) error {
	type Alias Page
	aux := &struct {
		Values []json.RawMessage `json:"values"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	for _, value := range aux.Values {
		var v Variable
		if err := json.Unmarshal(value, &v); err != nil {
			return err
		}
		p.Values = append(p.Values, v)
	}
	return nil
}
