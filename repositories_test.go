package bitbucket

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRepositoriesApiGroup_ListRepositoriesInWorkspace(t *testing.T) {
	type args struct {
		workspace string
	}
	type basicAuth struct {
		username string
		password string
	}
	type wantElement struct {
		type_     string
		name      string
		slug      string
		fullName  string
		isPrivate bool
	}
	type want struct {
		repositories []wantElement
	}
	tests := []struct {
		name    string
		auth    basicAuth
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "find one repository",
			auth: basicAuth{
				username: os.Getenv("BITBUCKET_USERNAME"),
				password: os.Getenv("BITBUCKET_APP_PASSWORD"),
			},
			args: args{
				workspace: os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
			},
			want: want{
				repositories: []wantElement{
					{
						type_:     "repository",
						name:      os.Getenv("BITBUCKET_REPO_SLUG"),
						slug:      os.Getenv("BITBUCKET_REPO_SLUG"),
						fullName:  fmt.Sprintf("%s/%s", os.Getenv("BITBUCKET_WORKSPACE_SLUG"), os.Getenv("BITBUCKET_REPO_SLUG")),
						isPrivate: true,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "not existing workspace",
			auth: basicAuth{
				username: os.Getenv("BITBUCKET_USERNAME"),
				password: os.Getenv("BITBUCKET_APP_PASSWORD"),
			},
			args: args{
				workspace: "not-existing-workspace",
			},
			want:    want{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClientWithBasicAuth(tt.auth.username, tt.auth.password)
			c.Debug = true

			g, err := c.Repositories.ListRepositoriesInWorkspace(tt.args.workspace)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			spew.Dump(g)
			got := want{
				repositories: []wantElement{},
			}
			for _, v := range g {
				got.repositories = append(got.repositories, wantElement{
					type_:     v.Type,
					name:      v.Name,
					slug:      v.Slug,
					fullName:  v.FullName,
					isPrivate: v.IsPrivate,
				})
			}
			for _, w := range tt.want.repositories {
				assert.Contains(t, got.repositories, w)
			}
		})
	}
}
