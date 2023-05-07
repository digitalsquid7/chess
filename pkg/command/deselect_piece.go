package command

import (
	"chess/pkg/state"
)

type DeselectPiece struct {
	TilePositions  *TilePositions
	MousePositionX float64
	MousePositionY float64
}

func (c *DeselectPiece) Execute(gameState *state.Game) {
	if gameState.Selected != nil {
		row, col, ok := c.TilePositions.GetTile(c.MousePositionX, c.MousePositionY)
		validMove := gameState.Selected.Piece.ValidMove[state.Coordinate{X: col, Y: row}]
		if ok && validMove {
			gameState.Board.Tiles[row][col] = gameState.Selected.Piece
			c.updateValidMoves(gameState)
		} else {
			gameState.Board.Tiles[gameState.Selected.OldPosition.Y][gameState.Selected.OldPosition.X] = gameState.Selected.Piece
		}

		gameState.Selected = nil
	}
}

func (c *DeselectPiece) updateValidMoves(gameState *state.Game) {
	for _, row := range gameState.Board.Tiles {
		for _, tile := range row {
			if tile != nil && tile.Name == state.Pawn && tile.Colour == state.White {
				//tile.UpdateValidMoves(gameState.Board.Tiles)
			}
		}
	}
}
