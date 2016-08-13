package maze

import (
	"image"
	"image/color"
	"image/png"
	"net/http"

	"github.com/Aorioli/procedural/concerns/direction"
	"github.com/Aorioli/procedural/concerns/point"
	"github.com/Aorioli/procedural/endpoints"
	"github.com/Aorioli/procedural/services/maze"
)

func encodeImageResponse(w http.ResponseWriter, response interface{}) error {
	errored, err := endpoints.CheckError(w, response)
	if err != nil {
		return err
	} else if errored {
		return nil
	}

	v, ok := response.(maze.Maze)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return nil
	}

	w.Header().Add(endpoints.ContentType, "image/png")
	return png.Encode(w, img(v, 20))
}

func drawCell(c maze.Cell, min, max point.Point, im *image.NRGBA, col color.Color) {
	for x := min.X; x < max.X; x++ {
		for y := min.Y; y < max.Y; y++ {
			im.Set(x, y, col)
			if x == min.X && !c.Has(direction.West) {
				im.Set(x, y, color.Black)
			} else if x == (max.X-1) && !c.Has(direction.East) {
				im.Set(x, y, color.Black)
			} else if y == (max.Y-1) && !c.Has(direction.South) {
				im.Set(x, y, color.Black)
			} else if y == min.Y && !c.Has(direction.North) {
				im.Set(x, y, color.Black)
			}
		}
	}
}

func img(m maze.Maze, size int) image.Image {
	im := image.NewNRGBA(image.Rect(0, 0, m.Width*size, m.Height*size))
	for y := 0; y <= m.Height; y++ {
		for x := 0; x <= m.Width; x++ {
			p := point.Point{X: x, Y: y}
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
				drawCell(
					m.Grid[p],
					point.Point{X: x * size, Y: y * size},
					point.Point{X: (x + 1) * size, Y: (y + 1) * size},
					im,
					c,
				)
			}
		}
	}
	return im
}
