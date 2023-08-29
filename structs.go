package bitbucket

type RequestOptions struct {
	Method string
	Path   string
}
type Link struct {
	Href string `json:"href"`
}
type Object struct {
	Type string `json:"type"`
}

type Page struct {
	Size     int    `json:"size"`
	Page     int    `json:"page"`
	Pagelen  int    `json:"pagelen"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}
