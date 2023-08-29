package bitbucket

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestWorkspacesApiGroup_GetWorkspace(t *testing.T) {
	type basicAuth struct {
		username string
		password string
	}

	type want struct {
		type_ string
		slug  string
	}

	tests := []struct {
		name      string
		auth      basicAuth
		workspace string
		want      want
		wantErr   bool
	}{
		{
			name: "correct workspace",
			auth: basicAuth{
				username: os.Getenv("BITBUCKET_USERNAME"),
				password: os.Getenv("BITBUCKET_APP_PASSWORD"),
			},
			workspace: os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
			wantErr:   false,
			want: want{
				type_: "workspace",
				slug:  os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
			},
		},
		{
			name: "incorrect credentials",
			auth: basicAuth{
				username: "incorrect",
				password: "incorrect",
			},
			workspace: os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
			want:      want{},
			wantErr:   true,
		},
		{
			name: "incorrect workspace",
			auth: basicAuth{
				username: os.Getenv("BITBUCKET_USERNAME"),
				password: os.Getenv("BITBUCKET_APP_PASSWORD"),
			},
			workspace: "foobarbaz",
			want:      want{},
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClientWithBasicAuth(tt.auth.username, tt.auth.password)

			c.Debug = true

			g, err := c.Workspaces.GetWorkspace(tt.workspace)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			spew.Dump(g)
			got := want{
				type_: g.Type,
				slug:  g.Slug,
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
