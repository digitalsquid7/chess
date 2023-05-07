package command

import (
	"chess/pkg/state"
	"github.com/faiface/pixel/pixelgl"
)

type Command interface {
	Execute(gameState *state.Game)
}

type Finder struct {
	TilePositions *TilePositions
}

func (f *Finder) Find(win *pixelgl.Window) []Command {
	commands := make([]Command, 0)

	if win.JustPressed(pixelgl.MouseButtonLeft) {
		commands = append(commands, &SelectPiece{
			TilePositions:  f.TilePositions,
			MousePositionX: win.MousePosition().X,
			MousePositionY: win.MousePosition().Y,
		})
	}

	if win.JustReleased(pixelgl.MouseButtonLeft) {
		commands = append(commands, &DeselectPiece{
			TilePositions:  f.TilePositions,
			MousePositionX: win.MousePosition().X,
			MousePositionY: win.MousePosition().Y,
		})
	}

	return commands
}

func NewFinder(tilePositions *TilePositions) *Finder {
	return &Finder{
		TilePositions: tilePositions,
	}
}
