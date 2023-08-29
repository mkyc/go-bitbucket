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
