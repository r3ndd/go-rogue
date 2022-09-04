package go_rogue

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/r3ndd/go-rogue/engine"
)

func InitGame(f func()) {
	ebiten.SetWindowTitle("Urban Rogue")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.MaximizeWindow()

	go f()
	engine.Run()
}
