package host

import (
	"game/pkg/config"
	"game/pkg/engine"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func RunGame() {
	game := &engine.Game{
		World:  engine.NewWorld(config.WorldWidth, config.WorldHeight),
		Screen: engine.NewScreen(config.ScreenWidth, config.ScreenHeight),
	}
	ebiten.SetWindowSize(game.Screen.Width, game.Screen.Height)
	ebiten.SetWindowTitle("Tada")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
