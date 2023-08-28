package bitbucket

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

const (
	DEFAULT_BITBUCKET_API_BASE_URL = "https://api.bitbucket.org/2.0"
	DEFAULT_HEADER_ACCEPT          = "application/json"
)

type Client struct {
	Auth       *auth
	apiBaseURL *url.URL
	HttpClient *http.Client

	User       userApiGroup
	Workspaces workspacesApiGroup
	Debug      bool
}

type auth struct {
	user, password string
	bearerToken    string
}

func NewClientWithBearerToken(token string) *Client {
	a := &auth{bearerToken: token}
	return newClient(a)
}

func NewClientWithBasicAuth(user, password string) *Client {
	a := &auth{user: user, password: password}
	return newClient(a)
}

func (c *Client) prepareRequest(o RequestOptions) (*http.Request, error) {
	req, err := http.NewRequest(o.Method, c.apiBaseURL.String()+o.Path, nil)
	if err != nil {
		return nil, err
	}
	c.addDefaultHeaders(req)
	c.authenticateRequest(req)
	return req, nil
}

func (c *Client) do(req *http.Request) ([]byte, error) {
	if c.Debug {
		dumpReq, _ := httputil.DumpRequestOut(req, true)
		fmt.Printf("Request: %s\n\n", dumpReq)
	}

	resp, err := c.HttpClient.Do(req)
	defer resp.Body.Close()

	if c.Debug {
		dumpResp, _ := httputil.DumpResponse(resp, true)
		fmt.Printf("Response: %s\n\n", dumpResp)
	}

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, errorNotFound
	}
	return io.ReadAll(resp.Body)
}

func (c *Client) addDefaultHeaders(req *http.Request) {
	req.Header.Add("Accept", DEFAULT_HEADER_ACCEPT)
}

func (c *Client) authenticateRequest(req *http.Request) {
	if c.Auth.bearerToken != "" {
		req.Header.Add("Authorization", "Bearer "+c.Auth.bearerToken)
	} else if c.Auth.user != "" && c.Auth.password != "" {
		req.SetBasicAuth(c.Auth.user, c.Auth.password)
	}
}

func newClient(a *auth) *Client {
	bitbucketUrl, err := setApiBaseUrl()
	if err != nil {
		log.Fatalf("invalid bitbucket url")
	}
	c := &Client{
		Auth:       a,
		apiBaseURL: bitbucketUrl,
	}
	c.User = &UserApiGroup{c: c}
	c.Workspaces = &WorkspacesApiGroup{c: c}

	c.HttpClient = &http.Client{
		Timeout: 5 * time.Second,
	}
	return c
}

func setApiBaseUrl() (*url.URL, error) {
	e := os.Getenv("BITBUCKET_API_BASE_URL")
	if e == "" {
		e = DEFAULT_BITBUCKET_API_BASE_URL
	}

	return url.Parse(e)
}
