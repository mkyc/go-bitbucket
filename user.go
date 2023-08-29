package bitbucket

import (
	"encoding/json"
)

type UserLinks struct {
	Avatar Link `json:"avatar"`
}

type User struct {
	Object
	Username    string    `json:"username"`
	DisplayName string    `json:"display_name"`
	Uuid        string    `json:"uuid"`
	Links       UserLinks `json:"links"`
	CreatedOn   string    `json:"created_on"`
}

type UserApiGroup struct {
	c *Client
}

func (u *UserApiGroup) GetCurrentUser() (*User, error) {
	o := RequestOptions{Method: "GET", Path: u.c.requestPath("/user")}
	req, err := u.c.newRequest(o)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := u.c.do(req)
	if err != nil {
		return nil, err
	}
	if u.c.Debug {
		u.c.logPrettyBody(bodyBytes)
	}
	var user User
	err = json.Unmarshal(bodyBytes, &user)
	return &user, err
}
