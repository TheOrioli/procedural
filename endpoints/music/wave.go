package music

import (
	"io"
	"net/http"

	"github.com/Aorioli/chopher/karplus"
	"github.com/Aorioli/chopher/wave"
	"github.com/Aorioli/procedural/endpoints"
)

func encodeWaveResponse(w http.ResponseWriter, response interface{}) error {
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

	wav := wave.New(wave.Stereo, 22000)
	v.Sound(&wav)
	w.Header().Add(endpoints.ContentType, "audio/wav")
	io.Copy(w, wav.Reader())

	return nil
}
