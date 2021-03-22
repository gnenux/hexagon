package hexagon

import (
	"fmt"
	"image/color"
)

var (
	Wall = color.RGBA{
		R: 116,
		G: 105,
		B: 101,
		A: 255,
	}

	Road = color.RGBA{
		R: 243,
		G: 179,
		B: 14,
		A: 255,
	}
)

type HexagonMap struct {
	mapSize int
	blocks  map[string]*Block
}

func NewHexagonMap(mapSize int, sideLength float64) *HexagonMap {

	blocks := make(map[string]*Block, 0)

	for i := -mapSize; i <= mapSize; i++ { //x轴
		for j := -mapSize; j <= mapSize; j++ { //y轴
			k := 0 - i - j
			if k >= -mapSize && k <= mapSize {
				block := Block{
					X:     i,
					Y:     j,
					Z:     0 - i - j,
					Lenth: sideLength,
					C:     color.White,
				}

				blocks[locationToHashKey(i, j, 0-i-j)] = &block
			}
		}
	}

	return &HexagonMap{
		blocks:  blocks,
		mapSize: mapSize,
	}
}

func (hm HexagonMap) GetBlock(x, y, z int) *Block {
	key := locationToHashKey(x, y, z)
	return hm.blocks[key]
}

func (hm HexagonMap) AllBlocks() []*Block {
	bs := make([]*Block, 0, len(hm.blocks))
	for _, b := range hm.blocks {
		bs = append(bs, b)
	}

	return bs
}

// GetNeighborBlock 获取(x, y, z)的距离小于等于distance的邻居block，excludeSelf表示是否排除中心block
func (hm HexagonMap) GetNeighborBlocks(x, y, z, distance int, excludeSelf bool) []*Block {
	bs := make([]*Block, 0, 0)
	for i := -distance; i <= distance; i++ { //x轴
		for j := -distance; j <= distance; j++ { //y轴
			if -i-j >= -distance && -i-j <= distance {
				if excludeSelf && (x+i == x && y+j == y && z-i-j == z) {
					continue
				}

				tb := hm.GetBlock(x+i, y+j, z-i-j)
				if tb != nil {
					bs = append(bs, tb)
				}
			}

		}
	}

	return bs
}

func locationToHashKey(x, y, z int) string {
	return fmt.Sprintf("%d-%d-%d", x, y, z)
}
