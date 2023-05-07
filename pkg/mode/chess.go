package mode

import (
	"chess/pkg/command"
	"chess/pkg/config"
	"chess/pkg/screen"
	"chess/pkg/sprite"
	"chess/pkg/state"
	"github.com/faiface/pixel/pixelgl"
	_ "image/png"
)

func Run() {
	screenConfig := config.LoadScreen()
	windowConfig := config.LoadWindow(screenConfig)
	win, err := pixelgl.NewWindow(*windowConfig)
	if err != nil {
		panic(err)
	}

	spriteFactory := sprite.NewFactory()
	pieceFactory := state.NewPieceFactory(spriteFactory)
	gameState := state.NewGame(spriteFactory, pieceFactory)
	tilePositions := command.FindTilePositions(screenConfig)
	finder := command.NewFinder(tilePositions)
	drawer := screen.NewDrawer(screenConfig)

	for !win.Closed() {
		commands := finder.Find(win)
		command.Execute(commands, gameState)
		drawer.Draw(win, gameState)
	}
}
