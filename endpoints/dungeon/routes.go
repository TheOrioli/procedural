package dungeon

import (
	"net/http"

	"github.com/Aorioli/procedural/endpoints"
	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

// HTTP returns the created routes
func HTTP(ctx context.Context) []endpoints.Route {
	return []endpoints.Route{
		{
			Path:   "/",
			Method: http.MethodGet,
			Handler: endpoints.MessageServer{
				Message: "Dungeon-as-a-service",
			},
		},
		{
			Path:   "/generate",
			Method: http.MethodGet,
			Handler: httptransport.NewServer(
				ctx,
				makeGenerateEndpoint(),
				decodeRequest(500),
				encodeJSONResponse,
			),
		},
		{
			Path:   "/generate/image",
			Method: http.MethodGet,
			Handler: httptransport.NewServer(
				ctx,
				makeGenerateEndpoint(),
				decodeRequest(100),
				encodeImageResponse,
			),
		},
	}
}
