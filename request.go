package gateway

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/pkg/errors"
	"net/http"
)

// NewRequest returns a new http.Request from the given Lambda event.
func NewRequest(ctx context.Context, version string, payload []byte) (*http.Request, error) {

	if version == "1.0" {
		var e events.APIGatewayProxyRequest

		err := json.Unmarshal(payload, &e)
		if err != nil {
			return nil, errors.Wrap(err, "Unmarshal request")
		}

		return decodeV1Request(ctx, e)
	}

	return nil, errors.Errorf("unsupported payload version: %s", version)
}
