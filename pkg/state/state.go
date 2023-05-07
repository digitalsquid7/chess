package state

import (
	"chess/pkg/sprite"
	"github.com/faiface/pixel"
)

type Coordinate struct {
	X int
	Y int
}

type Board struct {
	Tiles  [8][8]*Piece
	Sprite *pixel.Sprite
}

type Game struct {
	Board    *Board
	Selected *SelectedPiece
}

func NewGame(spriteFactory *sprite.Factory, pieceFactory *PieceFactory) *Game {
	var tiles [8][8]*Piece
	pieceOrder := [8]PieceName{Rook, Knight, Bishop, Queen, King, Bishop, Knight, Rook}

	for i, piece := range pieceOrder {
		tiles[0][i] = pieceFactory.Create(piece, White)
	}

	for i, piece := range pieceOrder {
		tiles[7][i] = pieceFactory.Create(piece, Black)
	}

	for i := 0; i < 8; i++ {
		tiles[1][i] = pieceFactory.Create(Pawn, White)
	}

	for i := 0; i < 8; i++ {
		tiles[6][i] = pieceFactory.Create(Pawn, Black)
	}

	return &Game{
		Board: &Board{
			Tiles:  tiles,
			Sprite: spriteFactory.Get(sprite.Board),
		},
	}
}
