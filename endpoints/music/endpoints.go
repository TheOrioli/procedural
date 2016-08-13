package music

import (
	"io"
	"math/rand"
	"net/http"

	"github.com/Aorioli/chopher/hasher"
	"github.com/Aorioli/chopher/karplus"
	"github.com/Aorioli/chopher/note"
	"github.com/Aorioli/chopher/scale"
	"github.com/Aorioli/chopher/song"

	"github.com/Aorioli/procedural/endpoints"
	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

// generateRequest struct
type generateRequest struct {
	length int
	seed   int64
	smoke  bool
}

type byteGenerator struct {
	length    int64
	processed int64
	rnd       *rand.Rand
}

func (b *byteGenerator) Read(d []byte) (int, error) {
	requested := len(d)
	var eof error
	if int64(requested) > b.length {
		requested = int(b.length)
		eof = io.EOF
	}

	for i := 0; i < requested; i++ {
		d[i] = byte(b.rnd.Uint32() % 256)
	}

	return requested, eof
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
		var sng *song.Song
		if req.smoke {
			s := song.New(song.Medium * 1.5)
			sng = s.Add(note.New(note.D, 2), note.Half+note.Quarter).
				AddWith(note.New(note.G, 2), note.Half+note.Quarter).
				AddAfter(note.New(note.F, 2), note.Half+note.Quarter).
				AddWith(note.New(note.AIS, 2), note.Half+note.Quarter).
				AddAfter(note.New(note.G, 2), note.Full).
				AddWith(note.New(note.C, 2), note.Full).
				AddAfter(note.New(note.D, 2), note.Half+note.Quarter).
				AddWith(note.New(note.G, 2), note.Half+note.Quarter).
				AddAfter(note.New(note.F, 2), note.Half+note.Quarter).
				AddWith(note.New(note.AIS, 2), note.Half+note.Quarter).
				AddAfter(note.New(note.GIS, 2), note.Half).
				AddWith(note.New(note.CIS, 2), note.Half).
				AddAfter(note.New(note.G, 2), note.Full).
				AddWith(note.New(note.C, 2), note.Full).
				AddAfter(note.New(note.D, 2), note.Half+note.Quarter).
				AddWith(note.New(note.G, 2), note.Half+note.Quarter).
				AddAfter(note.New(note.F, 2), note.Half+note.Quarter).
				AddWith(note.New(note.AIS, 2), note.Half+note.Quarter).
				AddAfter(note.New(note.G, 2), note.Full).
				AddWith(note.New(note.C, 2), note.Full).
				AddAfter(note.New(note.F, 2), note.Half+note.Quarter).
				AddWith(note.New(note.AIS, 2), note.Half+note.Quarter).
				AddAfter(note.New(note.D, 2), note.Full*2).
				AddWith(note.New(note.G, 2), note.Full*2)

			sng.Scale = scale.Minor.New(note.New(note.G, 2), false)
		} else {
			b := &byteGenerator{
				length: 64 + int64(req.length),
				rnd:    rand.New(rand.NewSource(req.seed)),
			}
			h := hasher.New(b)

			sng = h.Hash()
		}

		ks := karplus.Song{
			Song:         *sng,
			SamplingRate: 22000,
		}

		return ks, nil
	}
}
