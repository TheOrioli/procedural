package maze

import (
	"encoding/json"
	"errors"
	"image/png"
	"net/http"
	"strconv"

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

func encodeImageResponse(w http.ResponseWriter, response interface{}) error {
	switch v := response.(type) {
	case endpoints.Error:
		w.Header().Add(endpoints.ContentType, "application/json")
		w.WriteHeader(v.Status)
		return json.NewEncoder(w).Encode(v)
	case error:
		w.Header().Add(endpoints.ContentType, "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(w).Encode(v)
	}
	v, ok := response.(maze.Maze)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return nil
	}

	w.Header().Add(endpoints.ContentType, "image/png")
	return png.Encode(w, v.Image(20))
}
