package actor

import (
	"github.com/r3ndd/go-rogue/app/engine/entity"
	"github.com/r3ndd/go-rogue/app/engine/world"
)

func MoveSelf(self entity.InstanceId, args ...interface{}) bool {
	state, exists := entity.GetEntityState(self)

	if !exists || state == nil {
		return false
	}

	x, y := state.GetPos()
	dir := args[0].(string)

	switch dir {
	case "left":
		return MoveLeft(x, y)
	case "right":
		return MoveRight(x, y)
	case "up":
		return MoveUp(x, y)
	case "down":
		return MoveDown(x, y)
	default:
		return false
	}
}

func MoveLeft(x, y int) bool {
	return world.MoveEntity(x, y, x-1, y, true)
}

func MoveRight(x, y int) bool {
	return world.MoveEntity(x, y, x+1, y, true)
}

func MoveUp(x, y int) bool {
	return world.MoveEntity(x, y, x, y-1, true)
}

func MoveDown(x, y int) bool {
	return world.MoveEntity(x, y, x, y+1, true)
}
