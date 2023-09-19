package bitbucket

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDeploymentsApiGroup_GetEnvironment(t *testing.T) {
	type args struct {
		workspace       string
		repoSlug        string
		environmentUuid string
	}
	type basicAuth struct {
		username string
		password string
	}
	type want struct {
		type_ string
		name  string
	}
	tests := []struct {
		name    string
		auth    basicAuth
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "correct environment",
			auth: basicAuth{
				username: os.Getenv("BITBUCKET_USERNAME"),
				password: os.Getenv("BITBUCKET_APP_PASSWORD"),
			},
			args: args{
				workspace:       os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
				repoSlug:        os.Getenv("BITBUCKET_REPO_SLUG"),
				environmentUuid: os.Getenv("BITBUCKET_ENVIRONMENT_UUID"),
			},
			want: want{
				type_: "deployment_environment",
				name:  os.Getenv("BITBUCKET_ENVIRONMENT_NAME"),
			},
			wantErr: false,
		},
		{
			name: "incorrect credentials",
			auth: basicAuth{
				username: "incorrect",
				password: "incorrect",
			},
			args: args{
				workspace:       os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
				repoSlug:        os.Getenv("BITBUCKET_REPO_SLUG"),
				environmentUuid: os.Getenv("BITBUCKET_ENVIRONMENT_UUID"),
			},
			want:    want{},
			wantErr: true,
		},
		{
			name: "incorrect environment",
			auth: basicAuth{
				username: os.Getenv("BITBUCKET_USERNAME"),
				password: os.Getenv("BITBUCKET_APP_PASSWORD"),
			},
			args: args{
				workspace:       os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
				repoSlug:        os.Getenv("BITBUCKET_REPO_SLUG"),
				environmentUuid: "{0f0e5143-3edc-4b99-81a3-2866c3eee216}", //this is random uuid
			},
			want:    want{},
			wantErr: true,
		},
		// TODO: add test for repository with no environments
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClientWithBasicAuth(tt.auth.username, tt.auth.password)
			c.Debug = true

			g, err := c.Deployments.GetEnvironment(tt.args.workspace, tt.args.repoSlug, tt.args.environmentUuid)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			spew.Dump(g)
			got := want{
				type_: g.Type,
				name:  g.Name,
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDeploymentsApiGroup_ListEnvironments(t *testing.T) {
	type args struct {
		workspace string
		repoSlug  string
	}
	type basicAuth struct {
		username string
		password string
	}
	type wantElement struct {
		type_               string
		name                string
		slug                string
		environmentTypeName EnvironmentTypeName
		environmentTypeRank EnvironmentTypeRank
	}
	type want struct {
		environments []wantElement
	}
	tests := []struct {
		name    string
		auth    basicAuth
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "correct environments",
			auth: basicAuth{
				username: os.Getenv("BITBUCKET_USERNAME"),
				password: os.Getenv("BITBUCKET_APP_PASSWORD"),
			},
			args: args{
				workspace: os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
				repoSlug:  os.Getenv("BITBUCKET_REPO_SLUG"),
			},
			want: want{
				environments: []wantElement{
					{
						type_:               "deployment_environment",
						name:                os.Getenv("BITBUCKET_ENVIRONMENT_NAME"),
						slug:                os.Getenv("BITBUCKET_ENVIRONMENT_NAME"),
						environmentTypeName: EnvironmentTypeTest,
						environmentTypeRank: EnvironmentTypeRankTest,
					},
					{
						type_:               "deployment_environment",
						name:                os.Getenv("BITBUCKET_ENVIRONMENT_NAME_CREATE_VARIABLES"),
						slug:                os.Getenv("BITBUCKET_ENVIRONMENT_NAME_CREATE_VARIABLES"),
						environmentTypeName: EnvironmentTypeTest,
						environmentTypeRank: EnvironmentTypeRankTest,
					},
					{
						type_:               "deployment_environment",
						name:                os.Getenv("BITBUCKET_ENVIRONMENT_NAME_NO_VARIABLES"),
						slug:                os.Getenv("BITBUCKET_ENVIRONMENT_NAME_NO_VARIABLES"),
						environmentTypeName: EnvironmentTypeTest,
						environmentTypeRank: EnvironmentTypeRankTest,
					},
					{
						type_:               "deployment_environment",
						name:                os.Getenv("BITBUCKET_ENVIRONMENT_NAME_ONE_VARIABLE"),
						slug:                os.Getenv("BITBUCKET_ENVIRONMENT_NAME_ONE_VARIABLE"),
						environmentTypeName: EnvironmentTypeTest,
						environmentTypeRank: EnvironmentTypeRankTest,
					},
					{
						type_:               "deployment_environment",
						name:                os.Getenv("BITBUCKET_ENVIRONMENT_NAME_THREE_VARIABLES"),
						slug:                os.Getenv("BITBUCKET_ENVIRONMENT_NAME_THREE_VARIABLES"),
						environmentTypeName: EnvironmentTypeTest,
						environmentTypeRank: EnvironmentTypeRankTest,
					},
					{
						type_:               "deployment_environment",
						name:                os.Getenv("BITBUCKET_ENVIRONMENT_NAME_UPDATE_VARIABLES"),
						slug:                os.Getenv("BITBUCKET_ENVIRONMENT_NAME_UPDATE_VARIABLES"),
						environmentTypeName: EnvironmentTypeTest,
						environmentTypeRank: EnvironmentTypeRankTest,
					},
					{
						type_:               "deployment_environment",
						name:                os.Getenv("BITBUCKET_ENVIRONMENT_NAME_STAGING"),
						slug:                os.Getenv("BITBUCKET_ENVIRONMENT_NAME_STAGING"),
						environmentTypeName: EnvironmentTypeStaging,
						environmentTypeRank: EnvironmentTypeRankStaging,
					},
					{
						type_:               "deployment_environment",
						name:                os.Getenv("BITBUCKET_ENVIRONMENT_NAME_PRODUCTION"),
						slug:                os.Getenv("BITBUCKET_ENVIRONMENT_NAME_PRODUCTION"),
						environmentTypeName: EnvironmentTypeProduction,
						environmentTypeRank: EnvironmentTypeRankProduction,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "incorrect credentials",
			auth: basicAuth{
				username: "incorrect",
				password: "incorrect",
			},
			args: args{
				workspace: os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
				repoSlug:  os.Getenv("BITBUCKET_REPO_SLUG"),
			},
			want:    want{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClientWithBasicAuth(tt.auth.username, tt.auth.password)
			c.Debug = true
			// It looks to me that BB API ignores this parameter at all. In BitBucket Web UI there is following info:
			// "We now support a limit of up to 100 environments across test, staging and production."
			// So I think that this is a bug in BB API, or they are using strict pagelen of 100 and ignoring this
			// parameter.
			c.pagination.PageLength = 100

			g, err := c.Deployments.ListEnvironments(tt.args.workspace, tt.args.repoSlug)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			spew.Dump(g)
			got := want{
				environments: []wantElement{},
			}
			for _, v := range g {
				got.environments = append(got.environments, wantElement{
					type_:               v.Type,
					name:                v.Name,
					slug:                v.Slug,
					environmentTypeName: v.EnvironmentType.Name,
					environmentTypeRank: v.EnvironmentType.Rank,
				})
			}
			for _, w := range tt.want.environments {
				assert.Contains(t, got.environments, w)
			}
			//assert.Equal(t, tt.want, got)
		})
	}
}
