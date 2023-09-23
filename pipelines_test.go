package bitbucket

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestPipelinesApiGroup_GetVariableForWorkspace(t *testing.T) {
	type args struct {
		workspace string
		uuid      string
	}
	type basicAuth struct {
		username string
		password string
	}
	type want struct {
		type_   string
		key     string
		value   string
		secured bool
	}
	tests := []struct {
		name    string
		auth    basicAuth
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "correct variable insecure",
			auth: basicAuth{
				username: os.Getenv("BITBUCKET_USERNAME"),
				password: os.Getenv("BITBUCKET_APP_PASSWORD"),
			},
			args: args{
				workspace: os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
				uuid:      os.Getenv("BITBUCKET_WORKSPACE_INSECURE_VARIABLE_UUID"),
			},
			want: want{
				type_:   "pipeline_variable",
				key:     os.Getenv("BITBUCKET_WORKSPACE_INSECURE_VARIABLE_KEY"),
				value:   os.Getenv("BITBUCKET_WORKSPACE_INSECURE_VARIABLE_VALUE"),
				secured: false,
			},
			wantErr: false,
		},
		{
			name: "correct variable secure",
			auth: basicAuth{
				username: os.Getenv("BITBUCKET_USERNAME"),
				password: os.Getenv("BITBUCKET_APP_PASSWORD"),
			},
			args: args{
				workspace: os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
				uuid:      os.Getenv("BITBUCKET_WORKSPACE_SECURE_VARIABLE_UUID"),
			},
			want: want{
				type_:   "pipeline_variable",
				key:     os.Getenv("BITBUCKET_WORKSPACE_SECURE_VARIABLE_KEY"),
				value:   "",
				secured: true,
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
				uuid:      os.Getenv("BITBUCKET_WORKSPACE_INSECURE_VARIABLE_UUID"),
			},
			want:    want{},
			wantErr: true,
		},
		{
			name: "incorrect variable",
			auth: basicAuth{
				username: os.Getenv("BITBUCKET_USERNAME"),
				password: os.Getenv("BITBUCKET_APP_PASSWORD"),
			},
			args: args{
				workspace: os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
				uuid:      "{7a7ea792-72fe-4768-9e6c-9a420b066658}",
			},
			want:    want{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClientWithBasicAuth(tt.auth.username, tt.auth.password)
			c.Debug = true

			g, err := c.Pipelines.GetVariableForWorkspace(tt.args.workspace, tt.args.uuid)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			spew.Dump(g)
			got := want{
				type_:   g.Type,
				key:     g.Key,
				secured: g.Secured,
				value:   g.Value,
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPipelinesApiGroup_ListVariablesForEnvironment(t *testing.T) {
	type args struct {
		workspace       string
		repoSlug        string
		environmentUuid string
	}
	type basicAuth struct {
		username string
		password string
	}
	type wantElement struct {
		type_   string
		key     string
		value   string
		secured bool
	}
	type want struct {
		variables []wantElement
	}
	tests := []struct {
		name    string
		auth    basicAuth
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "environment with no variables",
			auth: basicAuth{
				username: os.Getenv("BITBUCKET_USERNAME"),
				password: os.Getenv("BITBUCKET_APP_PASSWORD"),
			},
			args: args{
				workspace:       os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
				repoSlug:        os.Getenv("BITBUCKET_REPO_SLUG"),
				environmentUuid: os.Getenv("BITBUCKET_ENVIRONMENT_UUID_NO_VARIABLES"),
			},
			want: want{
				variables: []wantElement{},
			},
			wantErr: false,
		},
		{
			name: "environment with one variable",
			auth: basicAuth{
				username: os.Getenv("BITBUCKET_USERNAME"),
				password: os.Getenv("BITBUCKET_APP_PASSWORD"),
			},
			args: args{
				workspace:       os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
				repoSlug:        os.Getenv("BITBUCKET_REPO_SLUG"),
				environmentUuid: os.Getenv("BITBUCKET_ENVIRONMENT_UUID_ONE_VARIABLE"),
			},
			want: want{
				variables: []wantElement{
					{
						type_:   "pipeline_variable",
						key:     "variable_1_key",
						value:   "variable_1_value",
						secured: false,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "environment with three variables",
			auth: basicAuth{
				username: os.Getenv("BITBUCKET_USERNAME"),
				password: os.Getenv("BITBUCKET_APP_PASSWORD"),
			},
			args: args{
				workspace:       os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
				repoSlug:        os.Getenv("BITBUCKET_REPO_SLUG"),
				environmentUuid: os.Getenv("BITBUCKET_ENVIRONMENT_UUID_THREE_VARIABLES"),
			},
			want: want{
				variables: []wantElement{
					{
						type_:   "pipeline_variable",
						key:     "variable_1_key",
						value:   "variable_1_value",
						secured: false,
					},
					{
						type_:   "pipeline_variable",
						key:     "variable_2_key",
						value:   "variable_2_value",
						secured: false,
					},
					{
						type_:   "pipeline_variable",
						key:     "variable_3_key",
						value:   "variable_3_value",
						secured: false,
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
				workspace:       os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
				repoSlug:        os.Getenv("BITBUCKET_REPO_SLUG"),
				environmentUuid: os.Getenv("BITBUCKET_ENVIRONMENT_UUID_THREE_VARIABLES"),
			},
			want:    want{},
			wantErr: true,
		},
		{
			name: "incorrect environment uuid",
			auth: basicAuth{
				username: os.Getenv("BITBUCKET_USERNAME"),
				password: os.Getenv("BITBUCKET_APP_PASSWORD"),
			},
			args: args{
				workspace:       os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
				repoSlug:        os.Getenv("BITBUCKET_REPO_SLUG"),
				environmentUuid: "{2f98df5f-c320-4a71-b9f0-4d89ac6dda8c}", // this is random uuid
			},
			want: want{
				variables: []wantElement{},
			},
			wantErr: false, //unfortunately, BitBucket returns empty page instead of error
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClientWithBasicAuth(tt.auth.username, tt.auth.password)
			c.Debug = true
			c.pagination.PageLength = 2 //set this for test with 3 variables to test pagination

			g, err := c.Pipelines.ListVariablesForEnvironment(tt.args.workspace, tt.args.repoSlug, tt.args.environmentUuid)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			spew.Dump(g)
			got := want{
				variables: []wantElement{},
			}
			for _, v := range g {
				got.variables = append(got.variables, wantElement{
					type_:   v.Type,
					key:     v.Key,
					value:   v.Value,
					secured: v.Secured,
				})
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPipelinesApiGroup_CreateVariableForEnvironment(t *testing.T) {
	currentTime := time.Now()
	formattedTime := currentTime.Format("060102_1504")

	type args struct {
		workspace       string
		repoSlug        string
		environmentUuid string
		variable        Variable
	}
	type basicAuth struct {
		username string
		password string
	}
	type want struct {
		type_   string
		key     string
		value   string
		secured bool
	}
	tests := []struct {
		name    string
		auth    basicAuth
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "correct variable insecure",
			auth: basicAuth{
				username: os.Getenv("BITBUCKET_USERNAME"),
				password: os.Getenv("BITBUCKET_APP_PASSWORD"),
			},
			args: args{
				workspace:       os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
				repoSlug:        os.Getenv("BITBUCKET_REPO_SLUG"),
				environmentUuid: os.Getenv("BITBUCKET_ENVIRONMENT_UUID_CREATE_VARIABLES"),
				variable: Variable{
					Key:     fmt.Sprintf("variable_1_key_%s", formattedTime),
					Value:   "variable_1_value",
					Secured: false,
				},
			},
			want: want{
				key:     fmt.Sprintf("variable_1_key_%s", formattedTime),
				value:   "variable_1_value",
				secured: false,
			},
			wantErr: false,
		},
		{
			name: "correct variable secure",
			auth: basicAuth{
				username: os.Getenv("BITBUCKET_USERNAME"),
				password: os.Getenv("BITBUCKET_APP_PASSWORD"),
			},
			args: args{
				workspace:       os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
				repoSlug:        os.Getenv("BITBUCKET_REPO_SLUG"),
				environmentUuid: os.Getenv("BITBUCKET_ENVIRONMENT_UUID_CREATE_VARIABLES"),
				variable: Variable{
					Key:     fmt.Sprintf("variable_2_key_%s", formattedTime),
					Value:   "variable_2_value",
					Secured: true,
				},
			},
			want: want{
				key:     fmt.Sprintf("variable_2_key_%s", formattedTime),
				value:   "",
				secured: true,
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
				environmentUuid: os.Getenv("BITBUCKET_ENVIRONMENT_UUID_CREATE_VARIABLES"),
				variable: Variable{
					Key:     fmt.Sprintf("variable_3_key_%s", formattedTime),
					Value:   "variable_3_value",
					Secured: false,
				},
			},
			want:    want{},
			wantErr: true,
		},
		//unfortunately, BitBucket is allowing to create variables for non-existing environments
		//{
		//	name: "incorrect environment uuid",
		//	auth: basicAuth{
		//		username: os.Getenv("BITBUCKET_USERNAME"),
		//		password: os.Getenv("BITBUCKET_APP_PASSWORD"),
		//	},
		//	args: args{
		//		workspace:       os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
		//		repoSlug:        os.Getenv("BITBUCKET_REPO_SLUG"),
		//		environmentUuid: "{51f2d4f7-ba7b-4f03-8e27-1b962f39f627}", // this is random uuid
		//		variable: Variable{
		//			Key:     fmt.Sprintf("variable_4_key_%s", formattedTime),
		//			Value:   "variable_4_value",
		//			Secured: false,
		//		},
		//	},
		//	want:    want{},
		//	wantErr: true,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClientWithBasicAuth(tt.auth.username, tt.auth.password)
			c.Debug = true

			g, err := c.Pipelines.CreateVariableForEnvironment(tt.args.workspace, tt.args.repoSlug, tt.args.environmentUuid, tt.args.variable)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			spew.Dump(g)
			got := want{
				key:     g.Key,
				secured: g.Secured,
				value:   g.Value,
			}
			assert.Equal(t, tt.want, got)
			assert.NotEmpty(t, g.Uuid)
		})
	}
}

func TestPipelinesApiGroup_UpdateVariableForEnvironment(t *testing.T) {
	currentTime := time.Now()
	formattedTime := currentTime.Format("060102_1504")

	type args struct {
		workspace       string
		repoSlug        string
		environmentUuid string
		create          Variable
		update          Variable
	}
	type basicAuth struct {
		username string
		password string
	}
	type want struct {
		type_   string
		key     string
		value   string
		secured bool
	}
	tests := []struct {
		name    string
		auth    basicAuth
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "correct variable insecure edit value",
			auth: basicAuth{
				username: os.Getenv("BITBUCKET_USERNAME"),
				password: os.Getenv("BITBUCKET_APP_PASSWORD"),
			},
			args: args{
				workspace:       os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
				repoSlug:        os.Getenv("BITBUCKET_REPO_SLUG"),
				environmentUuid: os.Getenv("BITBUCKET_ENVIRONMENT_UUID_UPDATE_VARIABLES"),
				create: Variable{
					Key:     fmt.Sprintf("variable_1_key_%s", formattedTime),
					Value:   "variable_1_value",
					Secured: false,
				},
				update: Variable{
					Key:     fmt.Sprintf("variable_1_key_%s", formattedTime),
					Value:   "variable_1_value_updated",
					Secured: false,
				},
			},
			want: want{
				key:     fmt.Sprintf("variable_1_key_%s", formattedTime),
				value:   "variable_1_value_updated",
				secured: false,
			},
			wantErr: false,
		},
		{
			name: "correct variable insecure edit key",
			auth: basicAuth{
				username: os.Getenv("BITBUCKET_USERNAME"),
				password: os.Getenv("BITBUCKET_APP_PASSWORD"),
			},
			args: args{
				workspace:       os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
				repoSlug:        os.Getenv("BITBUCKET_REPO_SLUG"),
				environmentUuid: os.Getenv("BITBUCKET_ENVIRONMENT_UUID_UPDATE_VARIABLES"),
				create: Variable{
					Key:     fmt.Sprintf("variable_2_key_%s", formattedTime),
					Value:   "variable_2_value",
					Secured: false,
				},
				update: Variable{
					Key:     fmt.Sprintf("variable_2_key_%s_updated", formattedTime),
					Value:   "variable_2_value",
					Secured: false,
				},
			},
			want: want{
				key:     fmt.Sprintf("variable_2_key_%s_updated", formattedTime),
				value:   "variable_2_value",
				secured: false,
			},
			wantErr: false,
		},
		{
			name: "correct variable secure edit value",
			auth: basicAuth{
				username: os.Getenv("BITBUCKET_USERNAME"),
				password: os.Getenv("BITBUCKET_APP_PASSWORD"),
			},
			args: args{
				workspace:       os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
				repoSlug:        os.Getenv("BITBUCKET_REPO_SLUG"),
				environmentUuid: os.Getenv("BITBUCKET_ENVIRONMENT_UUID_UPDATE_VARIABLES"),
				create: Variable{
					Key:     fmt.Sprintf("variable_3_key_%s", formattedTime),
					Value:   "variable_3_value",
					Secured: true,
				},
				update: Variable{
					Key:     fmt.Sprintf("variable_3_key_%s", formattedTime),
					Value:   "variable_3_value_updated",
					Secured: true,
				},
			},
			want: want{
				key:     fmt.Sprintf("variable_3_key_%s", formattedTime),
				value:   "",
				secured: true,
			},
			wantErr: false,
		},
		{
			name: "correct variable secure edit key",
			auth: basicAuth{
				username: os.Getenv("BITBUCKET_USERNAME"),
				password: os.Getenv("BITBUCKET_APP_PASSWORD"),
			},
			args: args{
				workspace:       os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
				repoSlug:        os.Getenv("BITBUCKET_REPO_SLUG"),
				environmentUuid: os.Getenv("BITBUCKET_ENVIRONMENT_UUID_UPDATE_VARIABLES"),
				create: Variable{
					Key:     fmt.Sprintf("variable_4_key_%s", formattedTime),
					Value:   "variable_4_value",
					Secured: true,
				},
				update: Variable{
					Key:     fmt.Sprintf("variable_4_key_%s_updated", formattedTime),
					Value:   "variable_4_value",
					Secured: true,
				},
			},
			want: want{
				key:     fmt.Sprintf("variable_4_key_%s_updated", formattedTime),
				value:   "",
				secured: true,
			},
			wantErr: false,
		},
		{
			name: "correct variable insecure edit value to secure",
			auth: basicAuth{
				username: os.Getenv("BITBUCKET_USERNAME"),
				password: os.Getenv("BITBUCKET_APP_PASSWORD"),
			},
			args: args{
				workspace:       os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
				repoSlug:        os.Getenv("BITBUCKET_REPO_SLUG"),
				environmentUuid: os.Getenv("BITBUCKET_ENVIRONMENT_UUID_UPDATE_VARIABLES"),
				create: Variable{
					Key:     fmt.Sprintf("variable_5_key_%s", formattedTime),
					Value:   "variable_5_value",
					Secured: false,
				},
				update: Variable{
					Key:     fmt.Sprintf("variable_5_key_%s", formattedTime),
					Value:   "variable_5_value",
					Secured: true,
				},
			},
			want: want{
				key:     fmt.Sprintf("variable_5_key_%s", formattedTime),
				value:   "",
				secured: true,
			},
			wantErr: false,
		},
		{
			name: "correct variable secure edit value to insecure",
			auth: basicAuth{
				username: os.Getenv("BITBUCKET_USERNAME"),
				password: os.Getenv("BITBUCKET_APP_PASSWORD"),
			},
			args: args{
				workspace:       os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
				repoSlug:        os.Getenv("BITBUCKET_REPO_SLUG"),
				environmentUuid: os.Getenv("BITBUCKET_ENVIRONMENT_UUID_UPDATE_VARIABLES"),
				create: Variable{
					Key:     fmt.Sprintf("variable_6_key_%s", formattedTime),
					Value:   "variable_6_value",
					Secured: true,
				},
				update: Variable{
					Key:     fmt.Sprintf("variable_6_key_%s", formattedTime),
					Value:   "variable_6_value", //TODO that in fact should be omitted to really test behaviour of BitBucket
					Secured: false,
				},
			},
			want: want{
				key:     fmt.Sprintf("variable_6_key_%s", formattedTime),
				value:   "variable_6_value",
				secured: false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClientWithBasicAuth(tt.auth.username, tt.auth.password)
			c.Debug = true

			gt, err := c.Pipelines.CreateVariableForEnvironment(tt.args.workspace, tt.args.repoSlug, tt.args.environmentUuid, tt.args.create)
			assert.NoError(t, err)
			assert.NotEmpty(t, gt.Uuid)
			g, err := c.Pipelines.UpdateVariableForEnvironment(tt.args.workspace, tt.args.repoSlug, tt.args.environmentUuid, gt.Uuid, tt.args.update)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			spew.Dump(g)
			got := want{
				key:     g.Key,
				secured: g.Secured,
				value:   g.Value,
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPipelinesApiGroup_GetConfiguration(t *testing.T) {
	type args struct {
		workspace string
		repoSlug  string
	}
	type basicAuth struct {
		username string
		password string
	}
	type want struct {
		type_   string
		enabled bool
	}
	tests := []struct {
		name    string
		auth    basicAuth
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "correct configuration",
			auth: basicAuth{
				username: os.Getenv("BITBUCKET_USERNAME"),
				password: os.Getenv("BITBUCKET_APP_PASSWORD"),
			},
			args: args{
				workspace: os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
				repoSlug:  os.Getenv("BITBUCKET_REPO_SLUG"),
			},
			want: want{
				type_:   "repository_pipelines_configuration",
				enabled: true,
			},
			wantErr: false,
		},
		// TODO add more tests
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClientWithBasicAuth(tt.auth.username, tt.auth.password)
			c.Debug = true

			g, err := c.Pipelines.GetConfiguration(tt.args.workspace, tt.args.repoSlug)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			spew.Dump(g)
			got := want{
				type_:   g.Type,
				enabled: g.Enabled,
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPipelinesApiGroup_UpdateConfiguration(t *testing.T) {
	type args struct {
		workspace string
		repoSlug  string
		update    PipelinesConfiguration
	}
	type basicAuth struct {
		username string
		password string
	}
	type want struct {
		type_   string
		enabled bool
	}
	tests := []struct {
		name    string
		auth    basicAuth
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "correct configuration",
			auth: basicAuth{
				username: os.Getenv("BITBUCKET_USERNAME"),
				password: os.Getenv("BITBUCKET_APP_PASSWORD"),
			},
			args: args{
				workspace: os.Getenv("BITBUCKET_WORKSPACE_SLUG"),
				repoSlug:  os.Getenv("BITBUCKET_REPO_SLUG"),
				update: PipelinesConfiguration{
					Enabled: true,
				},
			},
			want: want{
				type_:   "repository_pipelines_configuration",
				enabled: true,
			},
			wantErr: false,
		},
		// TODO add more tests
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClientWithBasicAuth(tt.auth.username, tt.auth.password)
			c.Debug = true

			g, err := c.Pipelines.UpdateConfiguration(tt.args.workspace, tt.args.repoSlug, tt.args.update)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			spew.Dump(g)
			got := want{
				type_:   g.Type,
				enabled: g.Enabled,
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
