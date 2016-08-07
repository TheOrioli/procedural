package maze

type direction uint8

const (
	north direction = iota + 1
	east
	south
	west
)

var opposite = map[direction]direction{
	north: south,
	east:  west,
	west:  east,
	south: north,
}

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
	directions := []direction{north, east, south, west}

	x, y := (r.Int() % width), (r.Int() % height)
	entrance := point{
		X: x,
		Y: y,
	}

	ret := Maze{
		Width:    width,
		Height:   height,
		Entrance: entrance,
	}

	live := make([]point, 0, (width*height)/2)
	live = append(live, entrance)

	visited := make(map[point]struct{}, (width * height))
	grid := make(map[point]cell, (width * height))
	var exits []point

	for len(live) != 0 {
		i := chooser.Choose(len(live))
		p := live[i]
		visited[p] = struct{}{}

		dead := true
		for _, i := range r.Perm(4) {
			d := directions[i]
			n := p.AddDirection(d)

			if !n.inBounds(width-1, height-1) {
				continue
			}

			if _, ok := visited[n]; ok {
				continue
			}

			dead = false

			c, ok := grid[p]
			if !ok {
				c = cell{
					Next: []direction{},
				}
			}
			c.Next = append(c.Next, d)
			grid[p] = c

			grid[n] = cell{
				Next: []direction{opposite[d]},
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
