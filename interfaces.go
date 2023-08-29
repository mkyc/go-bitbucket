package bitbucket

type userApiGroup interface {
	GetCurrentUser() (*User, error)
}

type workspacesApiGroup interface {
	GetWorkspace(name string) (*Workspace, error)
}

type deploymentsApiGroup interface {
	GetEnvironment(workspace, repoSlug, uuid string) (*Environment, error)
}

type pipelinesApiGroup interface {
	GetVariableForWorkspace(workspace, uuid string) (*Variable, error)
}
