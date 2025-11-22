package main

import "fmt"

type Tile struct {
	Passable bool
}

type Map struct {
	Width  int
	Height int
	Tiles  [][]Tile
}

func NewMap(width int, height int) *Map {
	tiles := make([][]Tile, height)
	for i := range tiles {
		tiles[i] = make([]Tile, width)
		for j := range tiles[i] {
			tiles[i][j] = Tile{Passable: true}
		}
	}
	return &Map{Width: width, Height: height, Tiles: tiles}
}

func (m Map) SetTile(x, y int, passable bool) {
	if x >= 0 && x < m.Width && y >= 0 && y < m.Height {
		m.Tiles[x][y].Passable = passable
	} else {
		fmt.Println("坐标超出地图范围")
	}
}

func (m Map) IsTilePassable(x, y int) bool {
	if x >= 0 && x < m.Width && y >= 0 && y < m.Height {
		return m.Tiles[x][y].Passable
	}
	fmt.Println("坐标超出地图范围")
	return false
}

func main() {
	gameMap := NewMap(10, 10)

	gameMap.SetTile(5, 8, false)

	if gameMap.IsTilePassable(5, 8) {
		fmt.Println("(5, 8)位置可通行")
	} else {
		fmt.Println("(5, 8)位置不可通行")
	}
	if gameMap.IsTilePassable(1, 7) {
		fmt.Println("(1, 7)位置可通行")
	} else {
		fmt.Println("(1, 7)位置不可通行")
	}

}
