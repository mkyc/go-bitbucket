package bitbucket

type userApiGroup interface {
	GetCurrentUser() (*User, error)
}

type workspacesApiGroup interface {
	GetWorkspace(name string) (*Workspace, error)
}

type deploymentsApiGroup interface {
	GetEnvironment(workspace, repoSlug, environmentUuid string) (*Environment, error)
}

type pipelinesApiGroup interface {
	GetVariableForWorkspace(workspace, variableUuid string) (*Variable, error)
	ListVariablesForEnvironment(workspace, repoSlug, environmentUuid string) ([]Variable, error)
	CreateVariableForEnvironment(workspace, repoSlug, environmentUuid string, variable Variable) (*Variable, error)
}
