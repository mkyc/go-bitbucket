package bitbucket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"time"
)

const (
	DEFAULT_BITBUCKET_API_BASE_URL = "https://api.bitbucket.org/2.0"
	DEFAULT_HEADER_ACCEPT          = "application/json"
	DEFAULT_PAGE_LENGTH            = 100
)

type Client struct {
	auth       *auth
	pagination *pagination
	apiBaseURL *url.URL
	HttpClient *http.Client

	User        userApiGroup
	Workspaces  workspacesApiGroup
	Deployments deploymentsApiGroup
	Pipelines   pipelinesApiGroup
	Debug       bool
}

type auth struct {
	user, password string
	bearerToken    string
}

type pagination struct {
	PageLength int
}

func NewClientWithBearerToken(token string) *Client {
	a := &auth{bearerToken: token}
	return newClient(a)
}

func NewClientWithBasicAuth(user, password string) *Client {
	a := &auth{user: user, password: password}
	return newClient(a)
}

func (c *Client) newRequest(o RequestOptions) (*http.Request, error) {
	req, err := http.NewRequest(o.Method, c.apiBaseURL.String()+o.Path, nil)
	if err != nil {
		return nil, err
	}
	c.addDefaultHeaders(req)
	c.authenticateRequest(req)
	c.addDefaultParams(req)
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
	if c.auth.bearerToken != "" {
		req.Header.Add("Authorization", "Bearer "+c.auth.bearerToken)
	} else if c.auth.user != "" && c.auth.password != "" {
		req.SetBasicAuth(c.auth.user, c.auth.password)
	}
}

func (c *Client) requestPath(template string, args ...interface{}) string {
	return fmt.Sprintf(template, args...)
}

func (c *Client) logPrettyBody(bodyBytes []byte) {
	var pretty bytes.Buffer
	err := json.Indent(&pretty, bodyBytes, "", "  ")
	if err != nil {
		log.Printf("JSON parse error: %s", err)
		// If it's not JSON, just print the original body text
		log.Println(string(bodyBytes))
	} else {
		log.Println(string(pretty.Bytes()))
	}
}

func (c *Client) addDefaultParams(req *http.Request) {
	q := req.URL.Query()
	q.Set("pagelen", strconv.Itoa(c.pagination.PageLength))
	req.URL.RawQuery = q.Encode()
}

func (c *Client) execute(o RequestOptions, target Typer) error {
	req, err := c.newRequest(o)
	if err != nil {
		return err
	}
	bodyBytes, err := c.do(req)
	if err != nil {
		return err
	}
	if c.Debug {
		c.logPrettyBody(bodyBytes)
	}
	return json.Unmarshal(bodyBytes, target)
}

func newClient(a *auth) *Client {
	bitbucketUrl, err := setApiBaseUrl()
	if err != nil {
		log.Fatalf("invalid bitbucket url")
	}
	c := &Client{
		auth:       a,
		apiBaseURL: bitbucketUrl,
		pagination: &pagination{
			PageLength: DEFAULT_PAGE_LENGTH,
		},
	}
	c.User = &UserApiGroup{c: c}
	c.Workspaces = &WorkspacesApiGroup{c: c}
	c.Deployments = &DeploymentsApiGroup{c: c}
	c.Pipelines = &PipelinesApiGroup{c: c}

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
