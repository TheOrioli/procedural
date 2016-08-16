package maze

import (
	"testing"

	"github.com/Aorioli/procedural/concerns/direction"
)

func TestHas(t *testing.T) {
	c := Cell{
		Next: []direction.Direction{direction.North},
	}
	for i, test := range []struct {
		d   direction.Direction
		out bool
	}{
		{
			d:   direction.North,
			out: true,
		},
		{
			d:   direction.South,
			out: false,
		},
	} {
		out := c.Has(test.d)
		if out != test.out {
			t.Errorf("Test %d: Expected %t, got %t", i, test.out, out)
		}
	}
}
