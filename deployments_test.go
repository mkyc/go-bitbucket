package bitbucket

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
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
		type_               string
		name                string
		slug                string
		environmentTypeName EnvironmentTypeName
		environmentTypeRank EnvironmentTypeRank
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
				type_:               "deployment_environment",
				name:                os.Getenv("BITBUCKET_ENVIRONMENT_NAME"),
				slug:                os.Getenv("BITBUCKET_ENVIRONMENT_NAME"),
				environmentTypeName: EnvironmentTypeTest,
				environmentTypeRank: EnvironmentTypeRankTest,
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
				type_:               g.Type,
				name:                g.Name,
				slug:                g.Slug,
				environmentTypeName: g.EnvironmentType.Name,
				environmentTypeRank: g.EnvironmentType.Rank,
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
		// TODO: add test for repository with no environments
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
		})
	}
}

func TestDeploymentsApiGroup_CreateEnvironment(t *testing.T) {
	currentTime := time.Now()
	formattedTime := currentTime.Format("060102_1504")

	type args struct {
		workspace   string
		repoSlug    string
		environment Environment
	}
	type basicAuth struct {
		username string
		password string
	}
	type want struct {
		type_               string
		name                string
		slug                string
		environmentTypeName EnvironmentTypeName
		environmentTypeRank EnvironmentTypeRank
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
				workspace: os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
				repoSlug:  os.Getenv("BITBUCKET_REPO_SLUG"),
				environment: Environment{
					Name: fmt.Sprintf("create-test-1-%s", formattedTime),
					EnvironmentType: EnvironmentType{
						Name: EnvironmentTypeTest,
						Rank: EnvironmentTypeRankTest,
					},
				},
			},
			want: want{
				type_:               "deployment_environment",
				name:                fmt.Sprintf("create-test-1-%s", formattedTime),
				slug:                fmt.Sprintf("create-test-1-%s", formattedTime),
				environmentTypeName: EnvironmentTypeTest,
				environmentTypeRank: EnvironmentTypeRankTest,
			},
			wantErr: false,
		},
		{
			name: "just environment type name",
			auth: basicAuth{
				username: os.Getenv("BITBUCKET_USERNAME"),
				password: os.Getenv("BITBUCKET_APP_PASSWORD"),
			},
			args: args{
				workspace: os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
				repoSlug:  os.Getenv("BITBUCKET_REPO_SLUG"),
				environment: Environment{
					Name: fmt.Sprintf("create-test-2-%s", formattedTime),
					EnvironmentType: EnvironmentType{
						Name: EnvironmentTypeStaging,
					},
				},
			},
			want: want{
				type_:               "deployment_environment",
				name:                fmt.Sprintf("create-test-2-%s", formattedTime),
				slug:                fmt.Sprintf("create-test-2-%s", formattedTime),
				environmentTypeName: EnvironmentTypeStaging,
				environmentTypeRank: EnvironmentTypeRankStaging,
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
				environment: Environment{
					Name: fmt.Sprintf("create-test-%s", formattedTime),
					EnvironmentType: EnvironmentType{
						Name: EnvironmentTypeTest,
						Rank: EnvironmentTypeRankTest,
					},
				},
			},
			want:    want{},
			wantErr: true,
		},
		{
			name: "incorrect environment type",
			auth: basicAuth{
				username: os.Getenv("BITBUCKET_USERNAME"),
				password: os.Getenv("BITBUCKET_APP_PASSWORD"),
			},
			args: args{
				workspace: os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
				repoSlug:  os.Getenv("BITBUCKET_REPO_SLUG"),
				environment: Environment{
					Name: fmt.Sprintf("create-test-%s", formattedTime),
					EnvironmentType: EnvironmentType{
						Name: "incorrect",
						Rank: EnvironmentTypeRankTest,
					},
				},
			},
			want:    want{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClientWithBasicAuth(tt.auth.username, tt.auth.password)
			c.Debug = true

			g, err := c.Deployments.CreateEnvironment(tt.args.workspace, tt.args.repoSlug, tt.args.environment)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			spew.Dump(g)
			got := want{
				type_:               g.Type,
				name:                g.Name,
				slug:                g.Slug,
				environmentTypeName: g.EnvironmentType.Name,
				environmentTypeRank: g.EnvironmentType.Rank,
			}
			assert.Equal(t, tt.want, got)
			assert.NotEmpty(t, g.Uuid)
		})
	}
}

func TestDeploymentsApiGroup_DeleteEnvironment(t *testing.T) {
	currentTime := time.Now()
	formattedTime := currentTime.Format("060102_1504")

	type args struct {
		workspace string
		repoSlug  string
	}
	type basicAuth struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		auth    basicAuth
		args    args
		create  bool
		wantErr bool
	}{
		{
			name: "correct environment",
			auth: basicAuth{
				username: os.Getenv("BITBUCKET_USERNAME"),
				password: os.Getenv("BITBUCKET_APP_PASSWORD"),
			},
			args: args{
				workspace: os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
				repoSlug:  os.Getenv("BITBUCKET_REPO_SLUG"),
			},
			create:  true,
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
			create:  false,
			wantErr: true,
		},
		{
			name: "not existing environment",
			auth: basicAuth{
				username: os.Getenv("BITBUCKET_USERNAME"),
				password: os.Getenv("BITBUCKET_APP_PASSWORD"),
			},
			args: args{
				workspace: os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
				repoSlug:  os.Getenv("BITBUCKET_REPO_SLUG"),
			},
			create:  false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClientWithBasicAuth(tt.auth.username, tt.auth.password)
			c.Debug = true

			environmentUuid := ""
			if tt.create {
				environment := Environment{
					Name: fmt.Sprintf("delete-test-%s", formattedTime),
					EnvironmentType: EnvironmentType{
						Name: EnvironmentTypeTest,
						Rank: EnvironmentTypeRankTest,
					},
				}
				g, err := c.Deployments.CreateEnvironment(tt.args.workspace, tt.args.repoSlug, environment)
				assert.NoError(t, err)
				spew.Dump(g)
				environmentUuid = g.Uuid
			} else {
				environmentUuid = "{73868aea-d679-4588-a7e1-5dcc7b4ecebc}" //this is random uuid
			}

			err := c.Deployments.DeleteEnvironment(tt.args.workspace, tt.args.repoSlug, environmentUuid)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
