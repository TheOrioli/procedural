package karplus

import (
	"io"
	"math"
	"math/rand"

	"github.com/Aorioli/chopher/note"
	"github.com/Aorioli/chopher/song"
)

type Note struct {
	Note   song.SongNote
	Buffer []float64
}

func NewNote(n song.SongNote, samplingRate int) *Note {
	buf := make([]float64, int(
		math.Ceil(
			float64(samplingRate)/n.Note.Frequency(),
		),
	))

	for i := 0; i < len(buf); i++ {
		buf[i] = rand.Float64()*2.0 - 1.0
	}
	return &Note{
		Note:   n,
		Buffer: buf,
	}
}

// Sound pops the current buffer value and appends the new one
func (n *Note) Sound() float64 {
	if n.Note.Note.Note == note.Rest {
		return 0
	}
	sampleValue := n.Buffer[0]

	v := (n.Buffer[0] + n.Buffer[1]) * 0.5 * 0.9999
	n.Buffer = append(n.Buffer[1:], v)

	return sampleValue
}

type Song struct {
	Song         song.Song
	SamplingRate int
	CurrentNotes []*Note
}

func (s *Song) Sound(w io.Writer) {
	var lastNote int
	for i, n := range s.Song.Notes {
		if n.IsValid(0) {
			s.CurrentNotes = append(s.CurrentNotes, NewNote(n, s.SamplingRate))
			lastNote = i
		}
	}

	var (
		time      float64
		increment = float64(s.Song.Tempo) / float64(s.SamplingRate)
	)

	temp := make([]*Note, 0, len(s.Song.Notes)/10+1)
	for len(s.CurrentNotes) > 0 {
		var sample float64
		temp = make([]*Note, 0, len(s.Song.Notes)/10+1)
		for _, n := range s.CurrentNotes {
			sample += n.Sound()
			if n.Note.IsValid(time) {
				temp = append(temp, n)
			}
		}

		sampleValue := int16(sample * 2048)
		w.Write([]byte{byte(sampleValue), byte(sampleValue >> 8)})
		time += increment

		for i := lastNote + 1; i < len(s.Song.Notes); i++ {
			n := s.Song.Notes[i]
			if n.IsValid(time) {
				temp = append(temp, NewNote(n, s.SamplingRate))
				lastNote = i
			}
		}

		s.CurrentNotes = temp

	}
}
