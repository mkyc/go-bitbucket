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
	// CreateEnvironment is described at:
	// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-deployments/#api-repositories-workspace-repo-slug-environments-post
	CreateEnvironment(workspace, repoSlug string, environment Environment) (*Environment, error)
	// DeleteEnvironment is described at:
	// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-deployments/#api-repositories-workspace-repo-slug-environments-environment-uuid-delete
	DeleteEnvironment(workspace, repoSlug, environmentUuid string) error
}

type pipelinesApiGroup interface {
	// GetVariableForWorkspace is described at:
	// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pipelines/#api-workspaces-workspace-pipelines-config-variables-variable-uuid-get
	GetVariableForWorkspace(workspace, variableUuid string) (*Variable, error)
	// ListVariablesForEnvironment is described at:
	// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pipelines/#api-repositories-workspace-repo-slug-deployments-config-environments-environment-uuid-variables-get
	ListVariablesForEnvironment(workspace, repoSlug, environmentUuid string) ([]Variable, error)
	// CreateVariableForEnvironment is described at:
	// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pipelines/#api-repositories-workspace-repo-slug-deployments-config-environments-environment-uuid-variables-post
	CreateVariableForEnvironment(workspace, repoSlug, environmentUuid string, variable Variable) (*Variable, error)
	// UpdateVariableForEnvironment is described at:
	// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pipelines/#api-repositories-workspace-repo-slug-deployments-config-environments-environment-uuid-variables-variable-uuid-put
	UpdateVariableForEnvironment(workspace, repoSlug, environmentUuid, variableUuid string, variable Variable) (*Variable, error)
}
