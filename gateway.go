// Package gateway provides a drop-in replacement for net/http.ListenAndServe for use in AWS Lambda & API Gateway.
package gateway

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

// ListenAndServe is a drop-in replacement for
// http.ListenAndServe for use within AWS Lambda.
//
// ListenAndServe always returns a non-nil error.
func ListenAndServe(addr string, h http.Handler) error {
	if h == nil {
		h = http.DefaultServeMux
	}

	lambda.StartHandler(handlerFunc(func(ctx context.Context, payload []byte) ([]byte, error) {
		apiGWEvent := struct {
			Version string `json:"version"`
		}{}

		if err := json.Unmarshal(payload, &apiGWEvent); err != nil {
			return []byte{}, err
		}

		r, err := NewRequest(ctx, apiGWEvent.Version, payload)
		if err != nil {
			return []byte{}, err
		}

		w := NewResponse()
		h.ServeHTTP(w, r)

		resp := w.End()

		return json.Marshal(&resp)

	}))

	return nil
}

// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as Lambda Handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler that calls f.
type handlerFunc func(ctx context.Context, payload []byte) ([]byte, error)

// Invoke calls f(ctx, payload).
func (f handlerFunc) Invoke(ctx context.Context, payload []byte) ([]byte, error) {
	return f(ctx, payload)
}
