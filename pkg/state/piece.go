package state

import (
	"chess/pkg/sprite"
	"github.com/faiface/pixel"
)

type PieceName string
type PieceColour string

const (
	Pawn   PieceName = "pawn"
	Rook             = "rook"
	Knight           = "knight"
	Bishop           = "bishop"
	Queen            = "queen"
	King             = "king"
)

const (
	White PieceColour = "white"
	Black             = "black"
)

type MoveFinder interface {
	Find(coor Coordinate, tiles [8][8]*Piece) map[Coordinate]bool
}

type Piece struct {
	Name       PieceName
	Colour     PieceColour
	Sprite     *pixel.Sprite
	ValidMove  map[Coordinate]bool
	moveFinder MoveFinder
}

func (p *Piece) UpdateValidMoves(coord Coordinate, tiles [8][8]*Piece) {
	p.ValidMove = p.moveFinder.Find(coord, tiles)

}

type SelectedPiece struct {
	Piece       *Piece
	OldPosition Coordinate
}

type PieceFactory struct {
	spriteFactory *sprite.Factory
	spriteNames   map[colourName]sprite.Name
	moveFinders   map[colourName]MoveFinder
}

type colourName struct {
	colour PieceColour
	name   PieceName
}

func (f *PieceFactory) Create(name PieceName, colour PieceColour) *Piece {
	return &Piece{
		Name:       name,
		Colour:     colour,
		Sprite:     f.spriteFactory.Get(f.spriteNames[colourName{colour, name}]),
		moveFinder: f.moveFinders[colourName{colour, name}],
	}
}

func NewPieceFactory(spriteFactory *sprite.Factory) *PieceFactory {
	spriteNames := map[colourName]sprite.Name{
		{White, Pawn}:   sprite.WhitePawn,
		{White, Rook}:   sprite.WhiteRook,
		{White, Knight}: sprite.WhiteKnight,
		{White, Bishop}: sprite.WhiteBishop,
		{White, Queen}:  sprite.WhiteQueen,
		{White, King}:   sprite.WhiteKing,
		{Black, Pawn}:   sprite.BlackPawn,
		{Black, Rook}:   sprite.BlackRook,
		{Black, Knight}: sprite.BlackKnight,
		{Black, Bishop}: sprite.BlackBishop,
		{Black, Queen}:  sprite.BlackQueen,
		{Black, King}:   sprite.BlackKing,
	}

	moveFinders := map[colourName]MoveFinder{
		{White, Pawn}: &WhitePawnMoveFinder{},
	}

	return &PieceFactory{
		spriteFactory: spriteFactory,
		spriteNames:   spriteNames,
		moveFinders:   moveFinders,
	}
}
