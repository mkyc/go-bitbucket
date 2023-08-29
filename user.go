package bitbucket

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
	o := RequestOptions{
		Method: "GET",
		Path:   u.c.requestPath("/user"),
	}
	var user User
	err := u.c.execute(&o, &user)
	return &user, err
}
