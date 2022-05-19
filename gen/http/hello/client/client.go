// Code generated by goa v3.7.4, DO NOT EDIT.
//
// hello client HTTP transport
//
// Command:
// $ goa gen hello-goakit/design

package client

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	goahttp "goa.design/goa/v3/http"
)

// Client lists the hello service endpoint HTTP clients.
type Client struct {
	// Hello Doer is the HTTP client used to make requests to the hello endpoint.
	HelloDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the hello service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		HelloDoer:           doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Hello returns an endpoint that makes HTTP requests to the hello service
// hello server.
func (c *Client) Hello() endpoint.Endpoint {
	var (
		decodeResponse = DecodeHelloResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildHelloRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.HelloDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("hello", "hello", err)
		}
		return decodeResponse(resp)
	}
}