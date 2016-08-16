package direction

import "testing"

func TestOpposite(t *testing.T) {
	for i, test := range []struct {
		in  Direction
		out Direction
	}{
		{
			in:  North,
			out: South,
		},
		{
			in:  South,
			out: North,
		},
		{
			in:  East,
			out: West,
		},
		{
			in:  West,
			out: East,
		},

		{
			in:  Direction(12),
			out: Direction(12),
		},
	} {
		out := Opposite(test.in)
		if out != test.out {
			t.Errorf("Test %d: Expected %v, got %v", i, test.out, out)
		}
	}
}
