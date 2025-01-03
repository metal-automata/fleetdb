package fleetdbapi

import (
	"context"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/pkg/errors"
)

var apiVersion = "v1"

// Client has the ability to talk to a fleetdb api server running at the given URI
type Client struct {
	url        string
	authToken  string
	httpClient Doer
	// dumper writes http request, responses to the given writer for debugging.
	dumper io.Writer
}

// Doer is an interface for an HTTP client that can make requests
type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

// NewClientWithToken will initialize a new hollow client with the given auth token and URL
func NewClientWithToken(authToken, endpoint string, doerClient Doer) (*Client, error) {
	if authToken == "" {
		return nil, newClientError("failed to initialize: no auth token provided")
	}

	c, err := NewClient(endpoint, doerClient)
	if err != nil {
		return nil, err
	}

	c.authToken = authToken

	return c, nil
}

// NewClient will return a fleetdb client configured to talk to the given URL.
// This client will not set the authorization header for you automatically and is left to be handled by the Doer that is provided.
//
// Example:
//
//	ctx := context.TODO()
//	provider, _ := oidc.NewProvider(ctx, "https://OIDC_ISSUER.COM")
//
//	oauthConfig := clientcredentials.Config{
//		ClientID:       "CLIENT_ID",
//		ClientSecret:   "CLIENT_SECRET",
//		TokenURL:       provider.Endpoint().TokenURL,
//		Scopes:         []string{"SCOPE", "SCOPE2"},
//		EndpointParams: url.Values{"audience": []string{"HOLLOW_AUDIENCE_VALUE"}},
//	}
//
//	c, _ := fleetdbapi.NewClient("HOLLOW_URI", oauthConfig.Client(ctx))
func NewClient(endpoint string, doerClient Doer) (*Client, error) {
	if endpoint == "" {
		return nil, newClientError("failed to initialize: no hollow api url provided")
	}

	endpoint = strings.TrimSuffix(endpoint, "/")

	c := &Client{
		url: endpoint,
	}

	c.httpClient = doerClient
	if c.httpClient == nil {
		// Use the default client as a fallback if one isn't passed
		c.httpClient = http.DefaultClient
	}

	if os.Getenv("DEBUG_CLIENT") != "" {
		c.dumper = os.Stdout
	}

	return c, nil
}

// SetToken allows you to change the token of a client
func (c *Client) SetToken(token string) {
	c.authToken = token
}

// NextPage will update the server response with the next page of results
func (c *Client) NextPage(ctx context.Context, resp ServerResponse, recs interface{}) (*ServerResponse, error) {
	if !resp.HasNextPage() {
		return nil, ErrNoNextPage
	}

	uri := resp.Links.Next.Href

	// for some reason in production the links are only the path
	if strings.HasPrefix(uri, "/api") {
		uri = c.url + uri
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, http.NoBody)
	if err != nil {
		return nil, err
	}

	r := ServerResponse{Records: &recs}
	err = c.do(req, &r)

	return &r, err
}

// post provides a reusable method for a standard POST to a fleetdbapi server
func (c *Client) post(ctx context.Context, path string, body interface{}) (*ServerResponse, error) {
	request, err := newPostRequest(ctx, c.url, path, body)
	if err != nil {
		return nil, err
	}

	r := ServerResponse{}

	if err := c.do(request, &r); err != nil {
		if r.Error != "" {
			return nil, errors.Wrap(err, r.Error)
		}

		return nil, err
	}

	return &r, nil
}

// postWithReciever provides a reusable method for a standard POST to a fleetdbapi server
func (c *Client) postWithReciever(ctx context.Context, path string, body, resp interface{}) error {
	request, err := newPostRequest(ctx, c.url, path, body)
	if err != nil {
		return err
	}

	return c.do(request, &resp)
}

// put provides a reusable method for a standard PUT to a fleetdbapi server
func (c *Client) put(ctx context.Context, path string, body interface{}) (*ServerResponse, error) {
	request, err := newPutRequest(ctx, c.url, path, body)
	if err != nil {
		return nil, err
	}

	r := ServerResponse{}

	if err := c.do(request, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type queryParams interface {
	setQuery(url.Values)
}

// list provides a reusable method for a standard list to a fleetdbapi server
func (c *Client) list(ctx context.Context, path string, params queryParams, resp interface{}) error {
	request, err := newGetRequest(ctx, c.url, path)
	if err != nil {
		return err
	}

	if params != nil {
		q := request.URL.Query()
		params.setQuery(q)
		request.URL.RawQuery = q.Encode()
	}

	return c.do(request, &resp)
}

// get provides a reusable method for a standard GET of a single item
func (c *Client) get(ctx context.Context, path string, resp interface{}) error {
	request, err := newGetRequest(ctx, c.url, path)
	if err != nil {
		return err
	}

	return c.do(request, &resp)
}

// getWithParams provides a reusable method that accepts query params for a standard GET of a single item
func (c *Client) getWithParams(ctx context.Context, path string, params queryParams, resp interface{}) error {
	request, err := newGetRequest(ctx, c.url, path)
	if err != nil {
		return err
	}

	if params != nil {
		q := request.URL.Query()
		params.setQuery(q)
		request.URL.RawQuery = q.Encode()
	}

	return c.do(request, &resp)
}

// post provides a reusable method for a standard post to a fleetdbapi server
func (c *Client) delete(ctx context.Context, path string) (*ServerResponse, error) {
	request, err := newDeleteRequest(ctx, c.url, path)
	if err != nil {
		return nil, err
	}

	var r ServerResponse

	return &r, c.do(request, &r)
}

// SetDumper sets requests and responses to be written to the given writer (os.Stdout for example) to aid debugging.
func (c *Client) SetDumper(w io.Writer) {
	c.dumper = w
}

// if c.dumper is set, dumpRequest writes outgoing client requests to dumper
func (c *Client) dumpRequest(req *http.Request) error {
	if c.dumper == nil {
		return nil
	}

	d, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		return err
	}

	d = append(d, '\n')

	_, err = c.dumper.Write(d)
	if err != nil {
		return err
	}

	return nil
}

// if c.dumper is set, dumpRequest writes incoming responses to dumper
func (c *Client) dumpResponse(resp *http.Response) error {
	if c.dumper == nil {
		return nil
	}

	d, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return err
	}

	d = append(d, '\n')

	_, err = c.dumper.Write(d)
	if err != nil {
		return err
	}

	return nil
}
