package config

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Screen struct {
	Scale        float64
	BoardOffset  float64
	TileWidth    float64
	EdgeWidth    float64
	ScreenWidth  float64
	ScreenHeight float64
}

func LoadScreen() *Screen {
	scale := float64(4)
	boardOffset := float64(10)

	return &Screen{
		Scale:        scale,
		BoardOffset:  scale * boardOffset,
		TileWidth:    scale * float64(16),
		EdgeWidth:    scale * float64(7),
		ScreenWidth:  142*scale + 2*boardOffset*scale,
		ScreenHeight: 142*scale + 2*boardOffset*scale,
	}
}

func LoadWindow(screenConfig *Screen) *pixelgl.WindowConfig {
	return &pixelgl.WindowConfig{
		Title:  "Chess",
		Bounds: pixel.R(0, 0, screenConfig.ScreenWidth, screenConfig.ScreenHeight),
		VSync:  true,
	}
}
