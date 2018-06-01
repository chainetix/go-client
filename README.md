# go-client

This client works for App Engine or vanilla Go!

### General Usage

Sign up at Chainetix.com and find a public network. Create a new project on that network, and use the API token when creating a new client instance.

```

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

// Vanilla client
func NewClient(token string) *Client {
	c := newClient(token)
	c.Client = httpclient.NewClient()
	return c
}

// App Engine client
func NewUrlfetchClient(ctx context.Context, token string) *Client {
	c := newClient(token)
	c.Client = httpclient.NewUrlfetchClient(ctx)
	return c
}

```

When using a Chainetix project you would typically take the following steps:

- Create a currency.
- Create an Agent for system operations.
- Issue a buffer of your currency to Agent.
- Create a user for 'Alice'.
- Transfer assets from Agent to Alice.
- Create a user for 'Bob'.
- Transfer assets from Agent to Bob.
- Transfer assets to/from Alice/Bob.
