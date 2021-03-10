package hexagon

import (
	"image/color"
	"testing"

	"github.com/fogleman/gg"
)

func TestHexagonMap(t *testing.T) {
	mapSize := 5
	l := float64(100)
	hm := NewHexagonMap(mapSize, l)

	dc := gg.NewContext(2000, 2000)
	for _, block := range hm.AllBlocks() {
		block.Draw(dc)
	}

	bs := hm.GetNeighborBlocks(0, 0, 0, 1, true)
	for _, block := range bs {
		block.C = color.RGBA{
			R: 0,
			G: 128,
			B: 0,
			A: 255,
		}
		block.Draw(dc)
	}

	dc.SavePNG("hexagonmap.png")
}
