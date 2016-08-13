package music

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Aorioli/chopher/karplus"
	"github.com/Aorioli/chopher/note"
	"github.com/Aorioli/chopher/song"
	"github.com/Aorioli/procedural/endpoints"
)

func decodeRequest(sizeLimit int) func(r *http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		request := &generateRequest{
			length: 10,
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

		s = q.Get("size")
		if s != "" {
			d, err := strconv.Atoi(s)
			if err != nil || d <= 0 || d > sizeLimit {
				return endpoints.Err(errors.New("Invalid size value"), http.StatusBadRequest), nil
			}
			request.length = d
		}

		s = q.Get("smoke_on_the_water")
		if s != "" {
			request.smoke = (s == "true")
		}

		return request, nil
	}
}

type jsonResponse struct {
	Scale []int          `json:"scale"`
	Key   jsonNote       `json:"key"`
	Tempo song.Tempo     `json:"tempo"`
	Song  []jsonSongNote `json:"song"`
}

type jsonNote struct {
	N string  `json:"note"`
	O int     `json:"octave"`
	F float64 `json:"frequency"`
}

type jsonSongNote struct {
	jsonNote
	Duration note.Duration `json:"duration"`
	StartAt  float64       `json:"start_at"`
}

func encodeJSONResponse(w http.ResponseWriter, response interface{}) error {
	errored, err := endpoints.CheckError(w, response)
	if err != nil {
		return err
	} else if errored {
		return nil
	}

	v, ok := response.(karplus.Song)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return nil
	}

	resp := jsonResponse{
		Scale: v.Song.Scale.Pattern.Scale,
		Key:   toNote(v.Song.Scale.Key),
		Tempo: v.Song.Tempo,
		Song:  toSongNote(v.Song.Notes),
	}

	w.Header().Add(endpoints.ContentType, endpoints.ApplicationJSON)
	return json.NewEncoder(w).Encode(resp)
}

func toNote(n note.Note) jsonNote {
	return jsonNote{
		F: n.Frequency(),
		O: n.Octave,
		N: note.Notes[n.Note],
	}
}

func toSongNote(s []song.SongNote) []jsonSongNote {
	ret := make([]jsonSongNote, len(s))
	for i := 0; i < len(s); i++ {
		ret[i] = jsonSongNote{
			jsonNote: toNote(s[i].Note),
			Duration: s[i].Duration,
			StartAt:  s[i].Start,
		}
	}
	return ret
}
