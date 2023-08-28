package bitbucket

import (
	"encoding/json"
)

type UserLinks struct {
	Avatar Link `json:"avatar"`
}

type User struct {
	Username    string    `json:"username"`
	DisplayName string    `json:"display_name"`
	Uuid        string    `json:"uuid"`
	Links       UserLinks `json:"links"`
	Type        string    `json:"type"`
	CreatedOn   string    `json:"created_on"`
}

type UserApiGroup struct {
	c *Client
}

func (u *UserApiGroup) GetCurrentUser() (*User, error) {
	o := RequestOptions{Method: "GET", Path: "/user"}
	req, err := u.c.prepareRequest(o)
	if err != nil {
		return nil, err
	}
	bytes, err := u.c.do(req)
	if err != nil {
		return nil, err
	}
	var user User
	err = json.Unmarshal(bytes, &user)
	return &user, err
}
