package maze

type point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (p point) AddDirection(d direction) point {
	switch d {
	case north:
		return p.Add(point{X: 0, Y: -1})
	case east:
		return p.Add(point{X: 1, Y: 0})
	case south:
		return p.Add(point{X: 0, Y: 1})
	case west:
		return p.Add(point{X: -1, Y: 0})
	default:
		return p
	}
}

func (p point) Add(a point) point {
	return point{X: p.X + a.X, Y: p.Y + a.Y}
}

func (p point) inBounds(width, height int) bool {
	if p.X < 0 || p.X > width {
		return false
	}

	if p.Y < 0 || p.Y > height {
		return false
	}

	return true
}

func (p point) Distance(other point) int {
	if p.X < other.X {
		p.X, other.X = other.X, p.X
	}

	if p.Y < other.Y {
		p.Y, other.Y = other.Y, p.Y
	}

	return (p.X - other.X) + (p.Y - other.Y)
}
