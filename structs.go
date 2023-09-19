package bitbucket

type RequestOptions struct {
	Method      string
	Path        string
	IsPageable  bool
	CurrentPage int
	HasNextPage bool
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
	Size     int    `json:"size"`
	Page     int    `json:"page"`
	PageLen  int    `json:"pagelen"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}
