package player

import (
	"github.com/r3ndd/go-rogue/app/engine/actor"
	"github.com/r3ndd/go-rogue/app/engine/entity"
	"github.com/r3ndd/go-rogue/app/engine/turn"
)

func MoveTurn(dir string) {
	if !isTurn {
		return
	}

	turn.ConsumeTurn(playerId, turn.TurnCap, func() {
		actor.ActSelf(playerId, "move_self", dir)
	})
}

func ToggleDoorTurn() {
	if !isTurn {
		return
	}

	state, _ := entity.GetEntityState(playerId)
	x, y := state.GetPos()

	switch actDir {
	case "left":
		x -= 1
	case "right":
		x += 1
	case "up":
		y -= 1
	case "down":
		y += 1
	default:
		return
	}

	turn.ConsumeTurn(playerId, turn.TurnCap, func() {
		actor.ActTarget(playerId, "toggle", x, y)
	})
}
