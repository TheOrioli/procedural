package direction

type Direction uint8

const (
	North Direction = iota + 1
	East
	South
	West
)

func Opposite(d Direction) Direction {
	switch d {
	case North:
		return South
	case East:
		return West
	case South:
		return North
	case West:
		return East
	default:
		return d
	}
}
