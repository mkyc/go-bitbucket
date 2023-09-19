package bitbucket

type userApiGroup interface {
	GetCurrentUser() (*User, error)
}

type workspacesApiGroup interface {
	GetWorkspace(name string) (*Workspace, error)
}

type deploymentsApiGroup interface {
	// ListEnvironments is described at:
	// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-deployments/#api-repositories-workspace-repo-slug-environments-get
	ListEnvironments(workspace, repoSlug string) ([]Environment, error)
	// GetEnvironment is described at:
	// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-deployments/#api-repositories-workspace-repo-slug-environments-environment-uuid-get
	GetEnvironment(workspace, repoSlug, environmentUuid string) (*Environment, error)
}

type pipelinesApiGroup interface {
	GetVariableForWorkspace(workspace, variableUuid string) (*Variable, error)
	ListVariablesForEnvironment(workspace, repoSlug, environmentUuid string) ([]Variable, error)
	CreateVariableForEnvironment(workspace, repoSlug, environmentUuid string, variable Variable) (*Variable, error)
	UpdateVariableForEnvironment(workspace, repoSlug, environmentUuid, variableUuid string, variable Variable) (*Variable, error)
}
