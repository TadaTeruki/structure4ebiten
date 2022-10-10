package engine

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	World  *World
	Screen *Screen
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.Screen.Width, g.Screen.Height
}

func (g *Game) Draw(image *ebiten.Image) {
	for y := 0; y < g.Screen.Height; y++ {
		for x := 0; x < g.Screen.Width; x++ {
			g.Screen.PutPixel(x, y, color.RGBA{0xff, 0xff, 0xff, 0xff})
		}
	}
	image.WritePixels(g.Screen.pixels)

}
