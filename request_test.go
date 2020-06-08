package gateway

import (
	"context"
	"github.com/tj/assert"
	"testing"
)

func TestDecodeRequest_v1(t *testing.T) {
	e := []byte(`{"version": "1.0", "path": "/pets/luna", "httpMethod": "GET"}`)
	r, err := NewRequest(context.Background(), "1.0", e)
	assert.NoError(t, err)

	assert.Equal(t, "GET", r.Method)
	assert.Equal(t, `/pets/luna`, r.URL.Path)
	assert.Equal(t, `/pets/luna`, r.URL.String())
}

func TestDecodeRequest_v2(t *testing.T) {
	e := []byte(`{"version": "2.0", "rawPath": "/pets/luna", "requestContext": {"http": {"method": "POST"}}}`)
	r, err := NewRequest(context.Background(), "2.0", e)
	assert.NoError(t, err)

	assert.Equal(t, "POST", r.Method)
	assert.Equal(t, `/pets/luna`, r.URL.Path)
	assert.Equal(t, `/pets/luna`, r.URL.String())
}
