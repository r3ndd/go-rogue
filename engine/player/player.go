package player

import (
	"image/color"

	"github.com/r3ndd/go-rogue/app/engine/actor"
	"github.com/r3ndd/go-rogue/app/engine/entity"
	"github.com/r3ndd/go-rogue/app/engine/turn"
	"github.com/r3ndd/go-rogue/app/engine/world"
)

type PlayerState struct {
	entity.EntityState
}

var playerTypeId entity.TypeId
var playerId entity.InstanceId
var playerState PlayerState = PlayerState{}
var isTurn = false

func init() {
	selfActions := []struct {
		entity.ActionId
		entity.SelfAction
	}{
		{"move_self", actor.MoveSelf},
	}

	targetActions := []struct {
		entity.ActionId
		entity.TargetAction
	}{
		{"open", nil},
		{"close", nil},
		{"toggle", nil},
	}

	regData := entity.RegData{
		Name:          "Youself",
		Rune:          '@',
		Color:         color.White,
		Class:         entity.Active,
		InitState:     &playerState,
		SelfActions:   selfActions,
		TargetActions: targetActions,
		ZIndex:        1,
	}
	playerTypeId = entity.RegisterEntityType(&regData)
}

func Spawn() {
	playerState = PlayerState{}
	playerId, _ = world.CreateEntity(playerTypeId, 0, 0, true)

	turn.RegisterActor(playerId, OnTurn, AfterTurn)
	AddInputListeners()
}

func OnTurn() {
	isTurn = true
}

func AfterTurn() {
	isTurn = false
}
