package bitbucket

import (
	"github.com/davecgh/go-spew/spew"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClientWithBearerToken(t *testing.T) {
	tests := []struct {
		name  string
		token string
		want  string
	}{
		{
			name:  "random token",
			token: "95z3l2D3GpfL0",
			want:  "95z3l2D3GpfL0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)
			got := NewClientWithBearerToken(tt.token)

			spew.Dump(got)

			a.Equal(tt.want, got.auth.bearerToken)
		})
	}
}
