package actor

import (
	"github.com/r3ndd/go-rogue/app/engine/entity"
	"github.com/r3ndd/go-rogue/app/engine/world"
)

// Self is active (not hybrid)
func ActSelf(self entity.InstanceId, actionId entity.ActionId, args ...interface{}) bool {
	state, exists := entity.GetEntityState(self)

	if !exists {
		return false
	}

	typeId := state.GetTypeId()
	action, exists := entity.EntitySelfActions[typeId][actionId]

	if !exists {
		return false
	}

	return action(self, args...)
}

// Self is active (not hybrid), target is active or hybrid
func ActTarget(self entity.InstanceId, actionId entity.ActionId, targetX, targetY int, args ...interface{}) bool {
	state, exists := entity.GetEntityState(self)

	if !exists {
		return false
	}

	typeId := state.GetTypeId()
	action, exists := entity.EntityTargetActions[typeId][actionId]

	if !exists {
		return false
	}

	targetTile := &world.Tiles[targetY][targetX]

	if targetTile.ActiveTypeId == 0 {
		return false
	}

	_, exists = entity.GetEntityState(targetTile.ActiveInstanceId)

	if !exists {
		targetTile.ActiveInstanceId = entity.RegisterEntityInstance(targetTile.ActiveTypeId, targetX, targetY)
	}

	target := targetTile.ActiveInstanceId
	res := true

	if action != nil {
		res = action(self, target, args...)
	}

	react(target, self, actionId, args...)
	return res
}

// Self is active or hybrid
func react(self entity.InstanceId, actor entity.InstanceId, actionId entity.ActionId, args ...interface{}) {
	state, _ := entity.GetEntityState(self)

	typeId := state.GetTypeId()
	reaction, exists := entity.EntityReactions[typeId][actionId]

	if !exists {
		return
	}

	reaction(self, actor, args...)
}
