package engine

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type View func(*ebiten.Image) error

var views = map[string]View{}

func AddView(name string, view View) {
	views[name] = view
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, view := range views {
		if view == nil {
			continue
		}

		err := view(screen)

		if err != nil {
			log.Fatal(err)
		}
	}
}
