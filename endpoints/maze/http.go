package maze

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Aorioli/procedural/concerns/point"
	"github.com/Aorioli/procedural/endpoints"
	"github.com/Aorioli/procedural/services/maze"
)

func decodeRequest(sizeLimit int) func(r *http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		request := &generateRequest{
			width:  10,
			height: 10,
		}

		q := r.URL.Query()

		s := q.Get("seed")
		if s != "" {
			d, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				return endpoints.Err(errors.New("Invalid seed value"), http.StatusBadRequest), nil
			}
			request.seed = d
		}

		s = q.Get("w")
		if s != "" {
			d, err := strconv.Atoi(s)
			if err != nil || d <= 0 || d > sizeLimit {
				return endpoints.Err(errors.New("Invalid width value"), http.StatusBadRequest), nil
			}
			request.width = d
		}

		s = q.Get("h")
		if s != "" {
			d, err := strconv.Atoi(s)
			if err != nil || d <= 0 || d > sizeLimit {
				return endpoints.Err(errors.New("Invalid height value"), http.StatusBadRequest), nil
			}
			request.height = d
		}
		return request, nil
	}
}

type gridPoint struct {
	P point.Point   `json:"point"`
	N []point.Point `json:"next"`
}

type generateResponse struct {
	Width    int         `json:"width"`
	Height   int         `json:"height"`
	Entrance point.Point `json:"entrance"`
	Exit     point.Point `json:"exit"`
	Grid     []gridPoint `json:"grid"`
}

func encodeJSONResponse(w http.ResponseWriter, response interface{}) error {
	errored, err := endpoints.CheckError(w, response)
	if err != nil {
		return err
	} else if errored {
		return nil
	}

	v, ok := response.(maze.Maze)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return nil
	}

	resp := generateResponse{
		Width:    v.Width,
		Height:   v.Height,
		Entrance: v.Entrance,
		Exit:     v.Exit,
	}

	g := make([]gridPoint, 0, len(v.Grid))
	for p, c := range v.Grid {
		gp := gridPoint{
			P: p,
		}

		n := make([]point.Point, len(c.Next))
		for i := 0; i < len(c.Next); i++ {
			n[i] = p.AddDirection(c.Next[i])
		}

		gp.N = n
		g = append(g, gp)
	}
	resp.Grid = g

	w.Header().Add(endpoints.ContentType, endpoints.ApplicationJSON)
	return json.NewEncoder(w).Encode(resp)
}
