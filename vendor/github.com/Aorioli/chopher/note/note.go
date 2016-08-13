package note

import (
	"fmt"
	"math"
)

const (
	a4Frequency float64  = 440.0
	magic       float64  = 1.0594630943592952645618252949463417007792043174941856
	Double      Duration = 2
	Full        Duration = 1
	Half        Duration = 0.5
	Quarter     Duration = 0.25
)

// Duration value
type Duration float64

// Musical scale halfstep values
const (
	C = iota
	CIS
	D
	DIS
	E
	F
	FIS
	G
	GIS
	A
	AIS
	B
)

const (
	Rest = 13
)

var (
	a4 = Note{
		Note:   A,
		Octave: 4,
	}
)

// Note represents a piano key note with a combination of a note and it's Octave
type Note struct {
	Note   int
	Octave int
}

func New(note, octave int) Note {
	return Note{
		Note:   note,
		Octave: octave,
	}
}

// AddHalfSteps adds a number of half steps and returns a new note
func (n Note) AddHalfSteps(hs int) Note {
	t := n.Note + hs
	if t >= 0 {
		n.Note = t % 12
		n.Octave = n.Octave + t/12
	} else {
		n.Note = B + t + 1
		n.Octave = n.Octave - 1
	}
	return n
}

// Frequency of the note
//
// https://en.wikipedia.org/wiki/Piano_key_frequencies
func (n Note) Frequency() float64 {
	if n.Note == Rest {
		return 1.0
	}
	return a4Frequency * math.Pow(magic, float64(HalfstepDistance(a4, n)))
}

func HalfstepDistance(from, to Note) int {
	return (to.Octave-from.Octave)*12 + (to.Note - from.Note)
}

var Notes = [...]string{
	"C", "C#", "D", "D#",
	"E", "F", "F#", "G",
	"G#", "A", "A#", "B",
}

func (n Note) String() string {
	return fmt.Sprintf("%s%d", Notes[n.Note], n.Octave)
}
