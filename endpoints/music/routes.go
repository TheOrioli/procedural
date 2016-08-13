package music

import (
	"net/http"

	"github.com/Aorioli/procedural/concerns/version"
	"github.com/Aorioli/procedural/endpoints"
	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

var (
	serviceIntro   = "Music-as-a-service"
	serviceVersion = version.Version{
		Major: 0,
		Minor: 0,
		Patch: 0,
	}
)

// HTTP returns the created routes
func HTTP(ctx context.Context) []endpoints.Route {
	return []endpoints.Route{
		{
			Path:    "/",
			Method:  http.MethodGet,
			Handler: endpoints.Description(serviceIntro, serviceVersion),
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
			Path:   "/generate/wave",
			Method: http.MethodGet,
			Handler: httptransport.NewServer(
				ctx,
				makeGenerateEndpoint(),
				decodeRequest(500),
				encodeWaveResponse,
			),
		},
	}
}
