package main

import (
	"chess/pkg/mode"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(mode.Run)
}
