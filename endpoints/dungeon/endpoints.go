package dungeon

import (
	"net/http"

	"github.com/Aorioli/procedural/endpoints"
	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"github.com/meshiest/go-dungeon/dungeon"
	"math/rand"
)

// generateRequest struct
type generateRequest struct {
	Size    int
	Rooms   int
	Seed    int64
}

func makeGenerateEndpoint() endpoint.Endpoint {
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

		grid := make([][]int, req.Size)
	  for i := 0; i < req.Size; i++ {
	    grid[i] = make([]int, req.Size)
	  }
	  
	  rand.Seed(req.Seed)

	  dngn := &dungeon.Dungeon{
	    Size:     req.Size,
	    NumRooms: req.Rooms,
	    Grid:     grid,
	    NumTries: 30,
	    MinSize:  3,
	    MaxSize:  12,
	    Rooms:    []dungeon.Rectangle{},
	    Regions:  []int{},
	    Bounds:   dungeon.Rectangle{X: 1, Y: 1, Width: req.Size - 2, Height: req.Size - 2},
	  }
	  dngn.Generate()

		return *dngn, nil
	}
}
