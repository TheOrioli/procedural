package dungeon

import (
	"image"
	"image/color"
	"image/png"
	"net/http"

	"github.com/Aorioli/procedural/concerns/point"
	"github.com/Aorioli/procedural/endpoints"
	"github.com/meshiest/go-dungeon/dungeon"
)

func encodeImageResponse(w http.ResponseWriter, response interface{}) error {
	errored, err := endpoints.CheckError(w, response)
	if err != nil {
		return err
	} else if errored {
		return nil
	}

	v, ok := response.(dungeon.Dungeon)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return nil
	}

	w.Header().Add(endpoints.ContentType, "image/png")
	return png.Encode(w, img(v, 20))
}

func drawCell(cell int, min, max point.Point, im *image.NRGBA, col color.Color) {
	if cell == 0 {
		return
	}

	for x := min.X; x < max.X; x++ {
		for y := min.Y; y < max.Y; y++ {
			im.Set(x, y, col)
			if x == min.X {
				im.Set(x, y, color.Black)
			} else if x == (max.X - 1) {
				im.Set(x, y, color.Black)
			} else if y == (max.Y - 1) {
				im.Set(x, y, color.Black)
			} else if y == min.Y {
				im.Set(x, y, color.Black)
			}
		}
	}
}

func img(d dungeon.Dungeon, size int) image.Image {
	im := image.NewNRGBA(image.Rect(0, 0, d.Size*size, d.Size*size))
	for y := 0; y < d.Size; y++ {
		for x := 0; x < d.Size; x++ {
			c := color.White
			drawCell(
				d.Grid[y][x],
				point.Point{X: x * size, Y: y * size},
				point.Point{X: (x + 1) * size, Y: (y + 1) * size},
				im,
				c,
			)
		}
	}
	return im
}
