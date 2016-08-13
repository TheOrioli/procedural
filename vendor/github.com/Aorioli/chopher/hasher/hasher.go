package hasher

import (
	"io"

	"github.com/Aorioli/chopher/note"
	"github.com/Aorioli/chopher/scale"
	"github.com/Aorioli/chopher/song"
)

var (
	scaleMap = [...]scale.Pattern{
		0: scale.Major,
		1: scale.Minor,
		2: scale.Blues,
		3: scale.Bebop,
	}

	durationMap = [...]note.Duration{
		0: note.Quarter,
		1: note.Half,
		2: note.Full,
		3: note.Double,
	}
)

type Hasher struct {
	r    io.Reader
	Song *song.Song
}

func New(r io.Reader) Hasher {
	buf := make([]byte, 16)
	r.Read(buf)

	speed := song.Slow + song.Tempo(float64(buf[0])/256)
	scl := scaleMap[int(buf[1])%len(scaleMap)]
	var sum int
	for i := 1; i < len(buf); i++ {
		sum += int(buf[i])
	}
	nt := note.Note{Note: note.C, Octave: 1}.AddHalfSteps(sum % 36)
	s := song.New(speed)
	s.Scale = scl.New(nt, buf[2]%2 == 0)

	return Hasher{
		Song: &s,
		r:    r,
	}
}

func (h *Hasher) Hash() *song.Song {
	io.Copy(h, h.r)
	songLength := len(h.Song.Notes)
	if songLength > 200 {
		h.Song.Notes = h.Song.Notes[:200]
	}
	songLength = len(h.Song.Notes)
	if songLength > 1 {
		h.Song.Notes[songLength-1].Duration = note.Full + 1.0
	}
	return h.Song
}

func (h *Hasher) Write(p []byte) (int, error) {
	for i := 1; i < len(p); i = i + 2 {
		add := h.Song.Scale.Notes[int(p[i-1])%len(h.Song.Scale.Notes)]
		h.Song = h.Song.Add(add, durationMap[int(p[i])%len(durationMap)])
	}
	return len(p), nil
}
