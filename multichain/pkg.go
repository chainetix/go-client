package chainetix

import (
	"golang.org/x/net/context"
	//
	"github.com/golangdaddy/tarantula/httpclient"
)

type Client struct {
	*httpclient.Client
	headers map[string]string
}

func newClient(token string) *Client {
	return &Client{
		headers: map[string]string{
			"User-Agent": "github.com/chainetix/go-client",
			"Content-Type": "application/json",
			"Authorization": token,
		},
	}
}

// Passing the httpClient allows us to use appengine/urlfetch client when needed.
func NewClient(token string) *Client {
	c := newClient(token)
	c.Client = httpclient.NewClient()
	return c
}

// Passing the httpClient allows us to use appengine/urlfetch client when needed.
func NewUrlfetchClient(ctx context.Context, token string) *Client {
	c := newClient(token)
	c.Client = httpclient.NewUrlfetchClient(ctx)
	return c
}
