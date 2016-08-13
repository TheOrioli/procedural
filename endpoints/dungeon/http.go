package dungeon

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Aorioli/procedural/concerns/direction"
	"github.com/Aorioli/procedural/concerns/point"
	"github.com/Aorioli/procedural/endpoints"
	"github.com/meshiest/go-dungeon/dungeon"
)

func decodeRequest(sizeLimit int) func(r *http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		request := &generateRequest{
			Size:  5,
			Rooms: 1,
			Seed:  0,
		}

		q := r.URL.Query()

		s := q.Get("seed")
		if s != "" {
			d, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				return endpoints.Err(errors.New("Invalid seed value"), http.StatusBadRequest), nil
			}
			request.Seed = d
		}

		s = q.Get("size")
		if s != "" {
			d, err := strconv.Atoi(s)
			if err != nil || d < 3 || d > sizeLimit {
				return endpoints.Err(errors.New("Invalid size value"), http.StatusBadRequest), nil
			}
			request.Size = d
		}

		s = q.Get("rooms")
		if s != "" {
			d, err := strconv.Atoi(s)
			if err != nil || d < 1 || d > 500 {
				return endpoints.Err(errors.New("Invalid room number value"), http.StatusBadRequest), nil
			}
			request.Rooms = d
		}
		return request, nil
	}
}

type generateResponse struct {
	Width  int         `json:"width"`
	Height int         `json:"height"`
	Grid   []gridPoint `json:"grid"`
}

type gridPoint struct {
	P point.Point   `json:"point"`
	N []point.Point `json:"next"`
}

func encodeJSONResponse(w http.ResponseWriter, response interface{}) error {
	w.Header().Add(endpoints.ContentType, endpoints.ApplicationJSON)
	switch v := response.(type) {
	case endpoints.Error:
		w.WriteHeader(v.Status)
		return json.NewEncoder(w).Encode(v)
	case error:
		w.WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(w).Encode(v)
	}

	v, ok := response.(dungeon.Dungeon)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return nil
	}

	resp := generateResponse{
		Width:  v.Size,
		Height: v.Size,
		Grid:   dungeonToGridPoints(v.Grid),
	}

	return json.NewEncoder(w).Encode(resp)
}

func dungeonToGridPoints(dungeon [][]int) []gridPoint {
	size := len(dungeon)
	ret := make([]gridPoint, 0, size*size)
	var directions = []direction.Direction{
		direction.North,
		direction.South,
		direction.East,
		direction.West,
	}

	min := point.Point{
		X: 0,
		Y: 0,
	}

	max := point.Point{
		X: size - 1,
		Y: size - 1,
	}

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			p := dungeon[y][x]
			if p == 0 {
				continue
			}

			gp := point.Point{
				X: x,
				Y: y,
			}

			next := make([]point.Point, 0, 4)
			for _, d := range directions {
				pn := gp.AddDirection(d)

				if pn.Inside(min, max) && dungeon[pn.Y][pn.X] == 1 {
					next = append(next, pn)
				}
			}

			ret = append(ret, gridPoint{
				P: gp,
				N: next,
			})
		}
	}
	return ret
}
