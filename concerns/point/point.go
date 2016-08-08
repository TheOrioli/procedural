package point

import "github.com/Aorioli/procedural/concerns/direction"

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (p Point) AddDirection(d direction.Direction) Point {
	switch d {
	case direction.North:
		return p.Add(Point{X: 0, Y: -1})
	case direction.East:
		return p.Add(Point{X: 1, Y: 0})
	case direction.South:
		return p.Add(Point{X: 0, Y: 1})
	case direction.West:
		return p.Add(Point{X: -1, Y: 0})
	default:
		return p
	}
}

func (p Point) Add(a Point) Point {
	return Point{X: p.X + a.X, Y: p.Y + a.Y}
}

func (p Point) Inside(min Point, max Point) bool {
	if p.X < min.X || p.X > max.X {
		return false
	}

	if p.Y < min.Y || p.Y > max.Y {
		return false
	}

	return true
}

func (p Point) Distance(other Point) int {
	if p.X < other.X {
		p.X, other.X = other.X, p.X
	}

	if p.Y < other.Y {
		p.Y, other.Y = other.Y, p.Y
	}

	return (p.X - other.X) + (p.Y - other.Y)
}
