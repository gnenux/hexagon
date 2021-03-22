package hexagon

import (
	"fmt"
	"image/color"
	"math"

	"github.com/fogleman/gg"
)

type Block struct {
	X     int
	Y     int
	Z     int
	Lenth float64
	C     color.Color
}

func (b Block) Draw(ctx *gg.Context) {
	centerX, centerY := b.GetLocation()

	locs := make([][]float64, 0, 0)

	// block之间的缝隙
	b.Lenth = b.Lenth - 2
	locA := []float64{centerX, centerY + b.Lenth}
	locB := []float64{centerX + b.Lenth*math.Sqrt(3)/2, centerY + b.Lenth/2}
	locs = append(locs, locB)
	locC := []float64{centerX + b.Lenth*math.Sqrt(3)/2, centerY - b.Lenth/2}
	locs = append(locs, locC)
	locD := []float64{centerX, centerY - b.Lenth}
	locs = append(locs, locD)
	locE := []float64{centerX - b.Lenth*math.Sqrt(3)/2, centerY - b.Lenth/2}
	locs = append(locs, locE)
	locF := []float64{centerX - b.Lenth*math.Sqrt(3)/2, centerY + b.Lenth/2}
	locs = append(locs, locF)
	// 在最后appendlocA
	locs = append(locs, locA)

	// 将第一个坐标移到locA
	// 以canvas中心为offset起始点
	ctx.MoveTo(locA[0]+float64(ctx.Width())/2, locA[1]+float64(ctx.Height())/2)
	for _, loc := range locs {
		ctx.LineTo(loc[0]+float64(ctx.Width())/2, loc[1]+float64(ctx.Height())/2)
	}

	ctx.SetColor(b.C)
	ctx.Fill()
	ctx.SetColor(color.Black)
	ctx.DrawStringAnchored(fmt.Sprintf("(%d,%d,%d)", b.X, b.Y, b.Z), centerX+float64(ctx.Width())/2, centerY+float64(ctx.Height())/2, 0.5, 0.5)
}

// GetLocation 转成笛卡尔直角坐标系
func (b Block) GetLocation() (float64, float64) {
	y := b.Lenth * 3 / 2 * float64(b.Z)
	x := float64(b.Y-b.X) * b.Lenth * math.Sqrt(3) / 2
	return x, y
}
