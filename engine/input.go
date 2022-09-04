package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type InputType int

type KeyboardHandler struct {
	Key   ebiten.Key
	Phase string
	Run   func()
}

type MouseButtonHandler struct {
	Button ebiten.MouseButton
	Phase  string
	Run    func()
}

type MouseScrollHandler struct {
	Phase string
	Run   func()
}

const (
	Keyboard    InputType = 0
	MouseButton           = 1
	MouseScroll           = 2
)

var keyboardHandlers = map[ebiten.Key]map[string][]KeyboardHandler{}
var mouseButtonHandlers = map[ebiten.MouseButton]map[string][]MouseButtonHandler{}
var mouseScrollHandlers = map[string][]MouseScrollHandler{}

func AddKeyboardListener(key ebiten.Key, phase string, run func()) {
	handler := KeyboardHandler{
		Key:   key,
		Phase: phase,
		Run:   run,
	}

	_, exists := keyboardHandlers[key]

	if !exists {
		keyboardHandlers[key] = map[string][]KeyboardHandler{}
	}

	keyboardHandlers[key][phase] = append(keyboardHandlers[key][phase], handler)
}

func AddMouseButtonListener(button ebiten.MouseButton, phase string, run func()) {
	handler := MouseButtonHandler{
		Button: button,
		Phase:  phase,
		Run:    run,
	}

	_, exists := mouseButtonHandlers[button]

	if !exists {
		mouseButtonHandlers[button] = map[string][]MouseButtonHandler{}
	}

	mouseButtonHandlers[button][phase] = append(mouseButtonHandlers[button][phase], handler)
}

func AddMouseScrollListener(phase string, run func()) {
	handler := MouseScrollHandler{
		Phase: phase,
		Run:   run,
	}

	mouseScrollHandlers[phase] = append(mouseScrollHandlers[phase], handler)
}

func handleInput() error {
	for key, handlersMap := range keyboardHandlers {
		if ebiten.IsKeyPressed(key) {
			for _, handler := range handlersMap["keypressed"] {
				handler.Run()
			}
		}

		if inpututil.IsKeyJustPressed(key) {
			for _, handler := range handlersMap["keydown"] {
				handler.Run()
			}
		}

		if inpututil.IsKeyJustReleased(key) {
			for _, handler := range handlersMap["keyup"] {
				handler.Run()
			}
		}
	}

	return nil
}
