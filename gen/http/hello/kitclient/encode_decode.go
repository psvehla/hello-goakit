// Code generated by goa v3.7.4, DO NOT EDIT.
//
// hello go-kit HTTP client encoders and decoders
//
// Command:
// $ goa gen hello-goakit/design

package client

import (
	"context"
	"hello-goakit/gen/http/hello/client"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	goahttp "goa.design/goa/v3/http"
)

// DecodeHelloResponse returns a go-kit DecodeResponseFunc suitable for
// decoding hello hello responses.
func DecodeHelloResponse(decoder func(*http.Response) goahttp.Decoder) kithttp.DecodeResponseFunc {
	dec := client.DecodeHelloResponse(decoder, false)
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		return dec(resp)
	}
}