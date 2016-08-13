package song

import (
	"github.com/Aorioli/chopher/note"
	"github.com/Aorioli/chopher/scale"
)

const (
	Fast   Tempo = 1.0
	Medium Tempo = 0.5
	Slow   Tempo = 0.33
)

// Tempo is the duration of the Full note in seconds
//
// Not my tempo!!
type Tempo float64

type SongNote struct {
	Note      note.Note
	Duration  note.Duration
	Start     float64
	ChordBase bool
}

type Song struct {
	Tempo Tempo
	Notes []SongNote
	Scale scale.Scale
}

func New(t Tempo) Song {
	return Song{
		Tempo: t,
	}
}

func (s *Song) add(note note.Note, duration note.Duration, start float64, base bool) *Song {
	s.Notes = append(s.Notes, SongNote{
		Note:      note,
		Duration:  duration,
		Start:     start,
		ChordBase: base,
	})
	return s
}

func (s *Song) AddAfter(note note.Note, duration note.Duration) *Song {
	lastNote := SongNote{}
	if len(s.Notes) > 0 {
		lastNote = s.Notes[len(s.Notes)-1]
	}
	return s.add(note, duration, lastNote.Start+float64(lastNote.Duration), true)
}

func (s *Song) AddWith(note note.Note, duration note.Duration) *Song {
	lastNote := SongNote{}
	if len(s.Notes) > 0 {
		lastNote = s.Notes[len(s.Notes)-1]
	}
	return s.add(note, duration, lastNote.Start, false)
}

func (s *Song) Add(addNote note.Note, duration note.Duration) *Song {
	notel := len(s.Notes)
	if notel == 0 {
		return s.add(addNote, duration, 0, true)
	}

	for _, c := range s.Scale.Chords {
		if len(c) > notel+1 {
			c = c[:notel+1]
		}

		var (
			base note.Note
			pos  = notel
		)

		for i := notel - 1; i >= 0; i-- {
			if s.Notes[i].ChordBase {
				base = s.Notes[i].Note
				pos = notel - i
				break
			}
		}
		if c.NotesInChord(base, addNote, pos) {
			return s.AddWith(addNote, duration)
		}
	}

	return s.AddAfter(addNote, duration)
}

func (s *SongNote) IsValid(time float64) bool {
	// log.Println(s.Note, time, s.Start, s.Start+float64(s.Duration))
	if time < s.Start {
		return false
	}
	return time < (s.Start + float64(s.Duration))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
