package dungeon

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/Aorioli/procedural/endpoints"
	"github.com/meshiest/go-dungeon/dungeon"
)

func decodeRequest(sizeLimit int) func(r *http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		request := &generateRequest{
			Size:  5,
			Rooms: 1,
			Seed:  time.Now().Unix(),
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
	Size    int         `json:"size"`
	Dungeon [][]int     `json:"dungeon"`
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
		Size:     v.Size,
		Dungeon:  v.Grid,
	}

	return json.NewEncoder(w).Encode(resp)
}
