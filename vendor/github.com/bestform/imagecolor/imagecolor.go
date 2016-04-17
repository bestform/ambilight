package imagecolor

import (
	"image"
	_ "image/jpeg" // @todo: currently only jpegs are supported. This is kind of awkward. But dynamic imports aren't possible
	"log"
	"os"
)

type Imagecolor struct {
	path      string // @todo: maybe it would be wiser to support generic reader instances instead of a file path. But for the given use case, this is the best way.
	precision int
}

func NewImagecolor(path string) Imagecolor {
	return Imagecolor{path, 10} // @todo: make precision more easily configurable. At the moment, 10 as a default is ok
}

func (i Imagecolor) AverageColor() (int, int, int) {
	reader, err := os.Open(i.path)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()

	var r, g, b, pixels uint32

	for y := bounds.Min.Y; y < bounds.Max.Y; y = y + i.precision {
		for x := bounds.Min.X; x < bounds.Max.X; x = x + i.precision {
			nr, ng, nb, _ := m.At(x, y).RGBA()
			r = r + nr
			g = g + ng
			b = b + nb
			pixels++
		}
	}

	// max: 65535

	xr := averageTo255(pixels, r)
	xg := averageTo255(pixels, g)
	xb := averageTo255(pixels, b)

	return int(xr), int(xg), int(xb)

}

func averageTo255(pixels, color uint32) uint32 {
	return uint32(float64(color/pixels) / 65535 * 255)
}
