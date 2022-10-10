package engine

import "image/color"

type Screen struct {
	Width  int
	Height int
	pixels []byte
}

func NewScreen(width, height int) *Screen {
	screen := new(Screen)
	screen.Width = width
	screen.Height = height
	screen.pixels = make([]byte, width*height*4)
	return screen
}

func (sc *Screen) PutPixel(x, y int, rgba color.RGBA) {

	i := (y*sc.Width + x) * 4

	if i > sc.Width*sc.Height*4 {
		return
	}

	sc.pixels[i+0] = rgba.R
	sc.pixels[i+1] = rgba.G
	sc.pixels[i+2] = rgba.B
	sc.pixels[i+3] = rgba.A
}
