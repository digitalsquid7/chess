package screen

import (
	"chess/pkg/config"
	"chess/pkg/state"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"time"
)

type Drawer struct {
	ScreenConfig *config.Screen
	frames       int
	second       <-chan time.Time
}

func (d *Drawer) Draw(win *pixelgl.Window, gameState *state.Game) {
	//win.Clear(color.RGBA{R: 69, G: 77, B: 95, A: 255})
	win.Clear(color.RGBA{R: 230, G: 234, B: 215, A: 255})
	moved := pixel.V(d.ScreenConfig.ScreenWidth/2, d.ScreenConfig.ScreenWidth/2)
	gameState.Board.Sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, d.ScreenConfig.Scale).Moved(moved))

	for row := 7; row >= 0; row-- {
		for col := 0; col < 8; col++ {
			if piece := gameState.Board.Tiles[row][col]; piece != nil && (gameState.Selected == nil ||
				(gameState.Selected.Piece != gameState.Board.Tiles[row][col])) {
				x := (float64(col) * d.ScreenConfig.TileWidth) + d.ScreenConfig.BoardOffset + d.ScreenConfig.EdgeWidth + (.5 * d.ScreenConfig.TileWidth)
				y := (float64(row+1) * d.ScreenConfig.TileWidth) + d.ScreenConfig.BoardOffset + d.ScreenConfig.EdgeWidth + d.ScreenConfig.Scale
				gameState.Board.Tiles[row][col].Sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, d.ScreenConfig.Scale).Moved(pixel.V(x, y)))
			}
		}
	}

	if gameState.Selected != nil {
		gameState.Selected.Piece.Sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, d.ScreenConfig.Scale).Moved(win.MousePosition()))
	}

	win.Update()

	d.frames++
	select {
	case <-d.second:
		win.SetTitle(fmt.Sprintf("Chess | FPS: %d", d.frames))
		d.frames = 0
	default:
	}
}

func NewDrawer(config *config.Screen) *Drawer {
	return &Drawer{
		ScreenConfig: config,
		frames:       0,
		second:       time.Tick(time.Second),
	}
}
