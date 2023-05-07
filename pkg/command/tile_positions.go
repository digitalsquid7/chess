package command

import "chess/pkg/config"

type TilePosition struct {
	Min float64
	Max float64
}

type TilePositions struct {
	X [8]TilePosition
	Y [8]TilePosition
}

func (t *TilePositions) GetTile(x float64, y float64) (int, int, bool) {
	row, ok := t.getTilePosition(y, t.Y)
	if !ok {
		return 0, 0, false
	}

	col, ok := t.getTilePosition(x, t.X)
	if !ok {
		return 0, 0, false
	}

	return row, col, true
}

func (t *TilePositions) getTilePosition(ordinate float64, tilePositions [8]TilePosition) (int, bool) {
	for i, tilePosition := range tilePositions {
		if ordinate >= tilePosition.Min && ordinate <= tilePosition.Max {
			return i, true
		}
	}
	return 0, false
}

func FindTilePositions(screenConfig *config.Screen) *TilePositions {
	var tilePositionsX, tilePositionsY [8]TilePosition

	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			minY := screenConfig.BoardOffset + screenConfig.EdgeWidth + screenConfig.TileWidth*float64(row)
			minX := screenConfig.BoardOffset + screenConfig.EdgeWidth + screenConfig.TileWidth*float64(col)
			maxY := screenConfig.BoardOffset + screenConfig.EdgeWidth + screenConfig.TileWidth*float64(row+1)
			maxX := screenConfig.BoardOffset + screenConfig.EdgeWidth + screenConfig.TileWidth*float64(col+1)

			tilePositionsX[col] = TilePosition{Min: minX, Max: maxX}
			tilePositionsY[row] = TilePosition{Min: minY, Max: maxY}
		}
	}

	return &TilePositions{
		X: tilePositionsX,
		Y: tilePositionsY,
	}
}
