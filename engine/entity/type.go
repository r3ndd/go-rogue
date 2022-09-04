package entity

import (
	"image/color"
)

type EntityClass int

type RegData struct {
	Name        string
	Desc        string
	Rune        rune
	Color       color.Color
	Overlapable bool
	ZIndex      byte
	Class       EntityClass
	InitState   EntityStateBase
	SelfActions []struct {
		ActionId
		SelfAction
	}
	TargetActions []struct {
		ActionId
		TargetAction
	}
	Reactions []struct {
		ActionId
		Reaction
	}
}

const (
	Passive EntityClass = 0
	Active              = 1
	Hybrid              = 2
)

var numEntities = 0
var Runes = map[TypeId]rune{
	0: '·',
}
var Names = map[TypeId]string{
	0: "",
}
var Descs = map[TypeId]string{
	0: "",
}
var Colors = map[TypeId]color.Color{
	0: color.White,
}

var Classes = map[TypeId]EntityClass{}
var InitStates = map[TypeId]EntityStateBase{}
var Overlapables = map[TypeId]bool{}
var ZIndexes = map[TypeId]byte{}

func RegisterEntityType(
	data *RegData,
) TypeId {
	numEntities++
	typeId := TypeId(numEntities)

	RegisterName(typeId, data.Name)
	RegisterRune(typeId, data.Rune)
	RegisterDesc(typeId, data.Desc)
	RegisterColor(typeId, data.Color)
	RegisterOverlapable(typeId, data.Overlapable)
	RegisterZIndex(typeId, data.ZIndex)
	RegisterClass(typeId, data.Class)
	RegisterInitState(typeId, data.InitState)
	RegisterSelfActions(typeId, data.SelfActions)
	RegisterTargetActions(typeId, data.TargetActions)
	RegisterReactions(typeId, data.Reactions)

	return typeId
}

func RegisterName(id TypeId, name string) {
	Names[id] = name
}

func RegisterRune(id TypeId, rune rune) {
	Runes[id] = rune
}

func RegisterDesc(id TypeId, desc string) {
	Descs[id] = desc
}

func RegisterColor(id TypeId, color color.Color) {
	Colors[id] = color
}

func RegisterOverlapable(id TypeId, overlapable bool) {
	Overlapables[id] = overlapable
}

func RegisterZIndex(id TypeId, zIndex byte) {
	ZIndexes[id] = zIndex
}

func RegisterClass(id TypeId, class EntityClass) {
	Classes[id] = class
}

func RegisterInitState(id TypeId, initState EntityStateBase) {
	InitStates[id] = initState
}

func RegisterSelfActions(id TypeId, selfActions []struct {
	ActionId
	SelfAction
}) {
	if _, exists := EntitySelfActions[id]; !exists {
		EntitySelfActions[id] = map[ActionId]SelfAction{}
	}

	for _, data := range selfActions {
		EntitySelfActions[id][data.ActionId] = data.SelfAction
	}
}

func RegisterTargetActions(id TypeId, targetActions []struct {
	ActionId
	TargetAction
}) {
	if _, exists := EntityTargetActions[id]; !exists {
		EntityTargetActions[id] = map[ActionId]TargetAction{}
	}

	for _, data := range targetActions {
		EntityTargetActions[id][data.ActionId] = data.TargetAction
	}
}

func RegisterReactions(id TypeId, reactions []struct {
	ActionId
	Reaction
}) {
	if _, exists := EntityReactions[id]; !exists {
		EntityReactions[id] = map[ActionId]Reaction{}
	}

	for _, data := range reactions {
		EntityReactions[id][data.ActionId] = data.Reaction
	}
}
