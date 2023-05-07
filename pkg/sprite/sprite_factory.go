package sprite

import (
	"chess/pkg/util"
	"fmt"
	"github.com/faiface/pixel"
)

type Name int

const (
	Board = iota
	BlackPawn
	BlackRook
	BlackKnight
	BlackBishop
	BlackQueen
	BlackKing
	WhitePawn
	WhiteRook
	WhiteKnight
	WhiteBishop
	WhiteQueen
	WhiteKing
)

type Factory struct {
	sprites map[Name]*pixel.Sprite
}

func (f *Factory) Get(name Name) *pixel.Sprite {
	sprite, ok := f.sprites[name]

	if !ok {
		panic(fmt.Errorf("provided sprite does not exist"))
	}

	return sprite
}

func NewFactory() *Factory {
	width := float64(16)
	height := float64(32)

	boardPicture := util.LoadPicture("assets\\images\\board.png")
	whitePicture := util.LoadPicture("assets\\images\\white_pieces.png")
	blackPicture := util.LoadPicture("assets\\images\\black_pieces.png")

	return &Factory{
		sprites: map[Name]*pixel.Sprite{
			Board:       pixel.NewSprite(boardPicture, boardPicture.Bounds()),
			BlackPawn:   pixel.NewSprite(blackPicture, pixel.R(0*width, 0, 1*width, height)),
			BlackRook:   pixel.NewSprite(blackPicture, pixel.R(1*width, 0, 2*width, height)),
			BlackKnight: pixel.NewSprite(blackPicture, pixel.R(2*width, 0, 3*width, height)),
			BlackBishop: pixel.NewSprite(blackPicture, pixel.R(3*width, 0, 4*width, height)),
			BlackQueen:  pixel.NewSprite(blackPicture, pixel.R(4*width, 0, 5*width, height)),
			BlackKing:   pixel.NewSprite(blackPicture, pixel.R(5*width, 0, 6*width, height)),
			WhitePawn:   pixel.NewSprite(whitePicture, pixel.R(0*width, 0, 1*width, height)),
			WhiteRook:   pixel.NewSprite(whitePicture, pixel.R(1*width, 0, 2*width, height)),
			WhiteKnight: pixel.NewSprite(whitePicture, pixel.R(2*width, 0, 3*width, height)),
			WhiteBishop: pixel.NewSprite(whitePicture, pixel.R(3*width, 0, 4*width, height)),
			WhiteQueen:  pixel.NewSprite(whitePicture, pixel.R(4*width, 0, 5*width, height)),
			WhiteKing:   pixel.NewSprite(whitePicture, pixel.R(5*width, 0, 6*width, height)),
		},
	}
}
