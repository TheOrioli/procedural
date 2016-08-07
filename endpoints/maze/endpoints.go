package maze

import (
	"net/http"

	"github.com/Aorioli/procedural/endpoints"
	"github.com/Aorioli/procedural/services/maze"
	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

// generateRequest struct
type generateRequest struct {
	width  int
	height int
	seed   int64
}

func makeGenerateEndpoint(svc maze.Service, algorithm maze.Chooser) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if req, ok := request.(error); ok {
			return req, nil
		}

		req, ok := request.(*generateRequest)
		if !ok {
			return endpoints.Err(
				errors.New("Invalid generate request"),
				http.StatusInternalServerError,
			), nil
		}
		m := svc.Generate(req.width, req.height, req.seed, algorithm)
		return m, nil
	}
}
