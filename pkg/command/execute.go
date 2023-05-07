package command

import "chess/pkg/state"

func Execute(commands []Command, gameState *state.Game) {
	for _, cmd := range commands {
		cmd.Execute(gameState)
	}
}
