package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/r3ndd/go-rogue/app/engine"
	"github.com/r3ndd/go-rogue/app/engine/world"
)

var dirSelActive bool = false
var shiftActive bool = false
var ctrlActive bool = false
var actDir string

func AddInputListeners() {
	engine.AddKeyboardListener(ebiten.KeyShift, "keydown", func() {
		shiftActive = true
	})
	engine.AddKeyboardListener(ebiten.KeyControl, "keydown", func() {
		ctrlActive = true
	})
	engine.AddKeyboardListener(ebiten.KeyShift, "keyup", func() {
		shiftActive = false
	})
	engine.AddKeyboardListener(ebiten.KeyControl, "keyup", func() {
		ctrlActive = false
	})

	engine.AddKeyboardListener(ebiten.KeySpace, "keydown", func() {
		if !dirSelActive {
			dirSelActive = true
		} else {
			HandleDir("self")
		}
	})

	engine.AddKeyboardListener(ebiten.KeyH, "keydown", func() { HandleDir("left") })
	engine.AddKeyboardListener(ebiten.KeyL, "keydown", func() { HandleDir("right") })
	engine.AddKeyboardListener(ebiten.KeyK, "keydown", func() { HandleDir("up") })
	engine.AddKeyboardListener(ebiten.KeyJ, "keydown", func() { HandleDir("down") })

	engine.AddKeyboardListener(ebiten.KeyD, "keydown", ToggleDoorTurn)
}

func HandleDir(dir string) {
	if shiftActive {
		world.MoveView(dir)
	} else if !dirSelActive {
		MoveTurn(dir)
	} else {
		actDir = dir
		dirSelActive = false
	}
}
