package bitbucket

type userApiGroup interface {
	GetCurrentUser() (*User, error)
}

type workspacesApiGroup interface {
	GetWorkspace(name string) (*Workspace, error)
}
