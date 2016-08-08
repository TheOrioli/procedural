package maze

import (
	"net/http"

	"github.com/Aorioli/procedural/endpoints"
	"github.com/Aorioli/procedural/services/maze"
	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

// HTTP returns the created routes
func HTTP(svc maze.Service, ctx context.Context) []endpoints.Route {
	return []endpoints.Route{
		{
			Path:   "/",
			Method: http.MethodGet,
			Handler: endpoints.MessageServer{
				Message: "Maze-as-a-service",
			},
		},
		{
			Path:   "/generate/backtrack",
			Method: http.MethodGet,
			Handler: httptransport.NewServer(
				ctx,
				makeGenerateEndpoint(svc, maze.Backtrack()),
				decodeRequest(500),
				encodeJSONResponse,
			),
		},
		{
			Path:   "/generate/backtrack/image",
			Method: http.MethodGet,
			Handler: httptransport.NewServer(
				ctx,
				makeGenerateEndpoint(svc, maze.Backtrack()),
				decodeRequest(100),
				encodeImageResponse,
			),
		},
	}
}
