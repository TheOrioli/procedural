package maze

import (
	"github.com/Aorioli/procedural/concerns/direction"
	"github.com/Aorioli/procedural/concerns/point"
)

type Cell struct {
	Next []direction.Direction `json:"next"`
}

func (c Cell) Has(d direction.Direction) bool {
	for _, n := range c.Next {
		if n == d {
			return true
		}
	}
	return false
}

type grid map[point.Point]Cell

// Maze internal structure
type Maze struct {
	Width    int         `json:"width"`
	Height   int         `json:"height"`
	Grid     grid        `json:"grid"`
	Entrance point.Point `json:"entrance"`
	Exit     point.Point `json:"exit"`
}
