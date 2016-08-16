package point

import (
	"testing"

	"github.com/Aorioli/procedural/concerns/direction"
)

func TestAddDirection(t *testing.T) {
	p := Point{}
	for i, test := range []struct {
		in  direction.Direction
		out Point
	}{
		{
			in:  direction.North,
			out: Point{X: 0, Y: -1},
		},
		{
			in:  direction.South,
			out: Point{X: 0, Y: 1},
		},
		{
			in:  direction.East,
			out: Point{X: 1, Y: 0},
		},
		{
			in:  direction.West,
			out: Point{X: -1, Y: 0},
		},
		{
			in:  direction.Direction(12),
			out: Point{},
		},
	} {
		out := p.AddDirection(test.in)
		if out != test.out {
			t.Errorf("Test %d: Expected %+v, got %+v", i, test.out, out)
		}
	}
}

func TestAdd(t *testing.T) {
	p := Point{}
	for i, test := range []struct {
		in  Point
		out Point
	}{
		{
			in:  Point{X: 0, Y: -1},
			out: Point{X: 0, Y: -1},
		},
		{
			in:  Point{X: 0, Y: 1},
			out: Point{X: 0, Y: 1},
		},
		{
			in:  Point{X: 1, Y: 0},
			out: Point{X: 1, Y: 0},
		},
		{
			in:  Point{X: -1, Y: 0},
			out: Point{X: -1, Y: 0},
		},
	} {
		out := p.Add(test.in)
		if out != test.out {
			t.Errorf("Test %d: Expected %+v, got %+v", i, test.out, out)
		}
	}
}

func TestInside(t *testing.T) {
	min := Point{}
	max := Point{X: 5, Y: 5}
	for i, test := range []struct {
		in  Point
		out bool
	}{
		{
			in:  Point{},
			out: true,
		},
		{
			in:  Point{X: 5, Y: 5},
			out: true,
		},
		{
			in:  Point{X: -1, Y: 5},
			out: false,
		},
		{
			in:  Point{X: 2, Y: 6},
			out: false,
		},
	} {
		out := test.in.Inside(min, max)
		if out != test.out {
			t.Errorf("Test %d: Expected %t, got  %t", i, test.out, out)
		}
	}
}

func TestDistance(t *testing.T) {
	for i, test := range []struct {
		a   Point
		b   Point
		out int
	}{
		{
			a:   Point{},
			b:   Point{X: 2, Y: 2},
			out: 4,
		},
		{
			a:   Point{X: 3, Y: 0},
			b:   Point{X: 2, Y: 2},
			out: 3,
		},
		{
			a:   Point{X: 0, Y: 3},
			b:   Point{X: 2, Y: 2},
			out: 3,
		},
		{
			a:   Point{X: 2, Y: 2},
			b:   Point{X: 2, Y: 2},
			out: 0,
		},
	} {
		out := test.a.Distance(test.b)
		if out != test.out {
			t.Errorf("Test %d: Expected %d, got  %d", i, test.out, out)
		}
	}
}
