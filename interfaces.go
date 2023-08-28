package bitbucket

type userApiGroup interface {
	GetCurrentUser() (*User, error)
}
