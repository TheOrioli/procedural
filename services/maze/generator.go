package maze

import (
	"github.com/Aorioli/procedural/concerns/direction"
	"github.com/Aorioli/procedural/concerns/point"
)

// Chooser interface is the interface that the cell picking algorithm should implement
type Chooser interface {
	Choose(int) int
}

// Randomizer abstracts the rand.Rand functions needed, for easier testing
type Randomizer interface {
	Int() int
	Intn(int) int
	Perm(int) []int
}

func generate(width, height int, r Randomizer, chooser Chooser) Maze {
	if width == 1 && height == 1 {
		return Maze{
			Width:  1,
			Height: 1,
			Grid: map[point.Point]Cell{
				point.Point{}: Cell{},
			},
		}
	}

	directions := []direction.Direction{
		direction.North,
		direction.East,
		direction.South,
		direction.West,
	}

	minBounds := point.Point{X: 0, Y: 0}
	maxBounds := point.Point{X: width - 1, Y: height - 1}

	x, y := (r.Int() % width), (r.Int() % height)
	entrance := point.Point{
		X: x,
		Y: y,
	}

	ret := Maze{
		Width:    width,
		Height:   height,
		Entrance: entrance,
	}

	live := make([]point.Point, 0, (width*height)/2)
	live = append(live, entrance)

	visited := make(map[point.Point]struct{}, (width * height))
	grid := make(map[point.Point]Cell, (width * height))
	var exits []point.Point

	for len(live) != 0 {
		i := chooser.Choose(len(live))
		p := live[i]
		visited[p] = struct{}{}

		dead := true
		for _, i := range r.Perm(4) {
			d := directions[i]
			n := p.AddDirection(d)

			if !n.Inside(minBounds, maxBounds) {
				continue
			}

			if _, ok := visited[n]; ok {
				continue
			}

			dead = false

			c, ok := grid[p]
			if !ok {
				c = Cell{
					Next: []direction.Direction{},
				}
			}
			c.Next = append(c.Next, d)
			grid[p] = c

			grid[n] = Cell{
				Next: []direction.Direction{direction.Opposite(d)},
			}

			live = append(live, n)
			break
		}

		if dead {
			live = append(live[:i], live[i+1:]...)
			if len(grid[p].Next) == 1 {
				exits = append(exits, p)
			}
		}
	}

	ret.Exit = exits[r.Intn(len(exits))]
	ret.Grid = grid
	return ret
}
