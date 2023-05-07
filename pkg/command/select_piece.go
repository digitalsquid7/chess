package command

import "chess/pkg/state"

type SelectPiece struct {
	TilePositions  *TilePositions
	MousePositionX float64
	MousePositionY float64
}

func (c *SelectPiece) Execute(gameState *state.Game) {
	row, col, ok := c.TilePositions.GetTile(c.MousePositionX, c.MousePositionY)

	if ok && gameState.Board.Tiles[row][col] != nil {
		gameState.Selected = &state.SelectedPiece{Piece: gameState.Board.Tiles[row][col], OldPosition: state.Coordinate{X: col, Y: row}}
		gameState.Board.Tiles[row][col] = nil
	}
}
