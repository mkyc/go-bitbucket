package bitbucket

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
	"testing"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
)

//type testLoggingTransport struct{}
//
//func (s *testLoggingTransport) RoundTrip(r *http.Request) (*http.Response, error) {
//	bytes, _ := httputil.DumpRequestOut(r, true)
//
//	resp, err := http.DefaultTransport.RoundTrip(r)
//	// err is returned after dumping the response
//
//	respBytes, _ := httputil.DumpResponse(resp, true)
//	bytes = append(bytes, respBytes...)
//
//	fmt.Printf("%s\n", bytes)
//
//	return resp, err
//}

func TestGetUser(t *testing.T) {
	username := os.Getenv("BITBUCKET_USERNAME")
	password := os.Getenv("BITBUCKET_APP_PASSWORD")
	basicAuthProvider, basicAuthProviderErr := securityprovider.NewSecurityProviderBasicAuth(username, password)
	if basicAuthProviderErr != nil {
		panic(basicAuthProviderErr)
	}

	clientWithResponses, err := NewClientWithResponses("https://api.bitbucket.org/2.0", WithRequestEditorFn(basicAuthProvider.Intercept))
	if err != nil {
		t.Error(err)
	}
	pr, err := clientWithResponses.GetUserWithResponse(context.Background())
	if err != nil {
		t.Error(err)
	}

	if pr.StatusCode() == http.StatusOK {
		t.Logf("User: %+v", pr.JSON200)
	}

}

func TestGetRepositories(t *testing.T) {
	username := os.Getenv("BITBUCKET_USERNAME")
	password := os.Getenv("BITBUCKET_APP_PASSWORD")
	workspace := os.Getenv("BITBUCKET_WORKSPACE_SLUG")

	basicAuthProvider, basicAuthProviderErr := securityprovider.NewSecurityProviderBasicAuth(username, password)
	if basicAuthProviderErr != nil {
		panic(basicAuthProviderErr)
	}

	c, err := NewClientWithResponses(
		"https://api.bitbucket.org/2.0",
		WithRequestEditorFn(basicAuthProvider.Intercept),
		//WithHTTPClient(
		//	&http.Client{
		//		Transport: &testLoggingTransport{},
		//	},
		//),
	)
	if err != nil {
		panic(err)
	}

	repos := make([]Repository, 0)
	hasNext := true
	page := 1
	for hasNext {
		// create context and add page parameter
		ctx := context.WithValue(context.Background(), "page", page)
		t.Logf("Page: %d", page)
		pr, err := c.GetRepositoriesWorkspaceWithResponse(
			ctx,
			workspace,
			nil,
			func(ctx context.Context, req *http.Request) error {
				page := ctx.Value("page").(int)
				q := req.URL.Query()
				q.Add("page", fmt.Sprintf("%d", page))
				req.URL.RawQuery = q.Encode()
				return nil
			},
		)
		if err != nil {
			t.Error(err)
		}
		if pr.StatusCode() == http.StatusOK {
			repos = append(repos, *pr.JSON200.Values...)
		}
		if pr.JSON200.Next != nil && *pr.JSON200.Next != "" {
			page++
		} else {
			hasNext = false
		}
	}
	log.Info(len(repos))
}
