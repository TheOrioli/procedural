package maze

import (
	"encoding/json"
	"image"
	"image/color"
)

type cell struct {
	Next []direction `json:"next"`
}

func (c cell) Has(d direction) bool {
	for _, n := range c.Next {
		if n == d {
			return true
		}
	}
	return false
}

func (c cell) Draw(min, max point, im *image.NRGBA, col color.Color) {
	for x := min.X; x <= max.X; x++ {
		for y := min.Y; y <= max.Y; y++ {
			im.Set(x, y, col)
			if x == min.X && !c.Has(west) {
				im.Set(x, y, color.Black)
			} else if x == max.X && !c.Has(east) {
				im.Set(x, y, color.Black)
			} else if y == max.Y && !c.Has(south) {
				im.Set(x, y, color.Black)
			} else if y == min.Y && !c.Has(north) {
				im.Set(x, y, color.Black)
			}
		}
	}
}

type grid map[point]cell

type jsonable struct {
	P point   `json:"point"`
	N []point `json:"next"`
}

func (g grid) MarshalJSON() ([]byte, error) {
	d := make([]jsonable, 0, len(g))

	for p, c := range g {
		j := jsonable{
			P: p,
		}
		n := make([]point, len(c.Next))
		for i := 0; i < len(c.Next); i++ {
			n[i] = p.AddDirection(c.Next[i])
		}

		j.N = n
		d = append(d, j)
	}
	return json.Marshal(d)
}

// Maze internal structure
type Maze struct {
	Width    int   `json:"width"`
	Height   int   `json:"height"`
	Grid     grid  `json:"grid"`
	Entrance point `json:"entrance"`
	Exit     point `json:"exit"`
}

// Image creates maze image of selected size
func (m Maze) Image(size int) image.Image {
	im := image.NewNRGBA(image.Rect(0, 0, m.Width*size+1, m.Height*size+1))
	for y := 0; y <= m.Height; y++ {
		for x := 0; x <= m.Width; x++ {
			p := point{X: x, Y: y}
			var c color.Color
			if p == m.Entrance {
				c = color.NRGBA{
					R: 0x00,
					G: 0x0a,
					B: 0xff,
					A: 0xff,
				}
			} else if p == m.Exit {
				c = color.NRGBA{
					R: 0xff,
					G: 0x00,
					B: 0x0a,
					A: 0xff,
				}
			} else {
				c = color.White
			}
			if _, ok := m.Grid[p]; ok {
				m.Grid[p].Draw(
					point{X: x * size, Y: y * size},
					point{X: (x + 1) * size, Y: (y + 1) * size},
					im,
					c,
				)
			}
		}
	}
	return im
}
