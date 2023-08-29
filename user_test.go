package bitbucket

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestUserApiGroup_GetCurrentUser(t *testing.T) {
	type basicAuth struct {
		username string
		password string
	}

	type want struct {
		username    string
		type_       string
		displayName string
	}

	tests := []struct {
		name    string
		auth    basicAuth
		want    want
		wantErr bool
	}{
		{
			name: "correct user",
			auth: basicAuth{
				username: os.Getenv("BITBUCKET_USERNAME"),
				password: os.Getenv("BITBUCKET_APP_PASSWORD"),
			},
			wantErr: false,
			want: want{
				username:    os.Getenv("BITBUCKET_USERNAME"),
				type_:       "user",
				displayName: os.Getenv("BITBUCKET_DISPLAY_NAME"),
			},
		},
		{
			name: "incorrect credentials",
			auth: basicAuth{
				username: "incorrect",
				password: "incorrect",
			},
			want:    want{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClientWithBasicAuth(tt.auth.username, tt.auth.password)

			c.Debug = true

			g, err := c.User.GetCurrentUser()
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			spew.Dump(g)
			got := want{
				username:    g.Username,
				type_:       g.Type,
				displayName: g.DisplayName,
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
