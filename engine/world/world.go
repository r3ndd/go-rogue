package world

import (
	"log"

	"github.com/r3ndd/go-rogue/app/engine"
	"github.com/r3ndd/go-rogue/app/engine/entity"
)

// 8 bytes
type EntityTile struct {
	PassiveTypeId    entity.TypeId
	ActiveTypeId     entity.TypeId
	ActiveInstanceId entity.InstanceId
}
type WorldTiles [WorldSize][WorldSize]EntityTile

const WorldSize = 10e3

var Tiles = WorldTiles{}
var DepthTiles = map[int][][]EntityTile{}
var Generator func()

var numDepthTiles = 0

func init() {
	// fmt.Println(unsafe.Sizeof(EntityTile{}))
}

func CreateEntity(typeId entity.TypeId, x, y int, destroy bool) (entity.InstanceId, bool) {
	class, exists := entity.Classes[typeId]

	if !exists {
		log.Fatal("Attempted to create unregistered entity")
	}

	if x < 0 || y < 0 || x >= WorldSize || y >= WorldSize {
		return entity.InstanceId{}, false
	}

	tile := &Tiles[y][x]

	if class != entity.Passive {
		if tile.ActiveTypeId != 0 {
			if !destroy {
				return entity.InstanceId{}, false
			} else {
				DestroyEntityAt(x, y, true)
			}
		}

		tile.ActiveTypeId = typeId

		if class == entity.Active {
			id := entity.RegisterEntityInstance(typeId, x, y)
			tile.ActiveInstanceId = id
			return id, true
		}

		return entity.InstanceId{}, true
	} else {
		if tile.PassiveTypeId != 0 {
			if !destroy {
				return entity.InstanceId{}, false
			} else {
				DestroyEntityAt(x, y, false)
			}
		}

		tile.PassiveTypeId = typeId
		return entity.InstanceId{}, true
	}
}

func DestroyEntityAt(x, y int, active bool) {
	if x < 0 || y < 0 || x >= WorldSize || y >= WorldSize {
		return
	}

	tile := &Tiles[y][x]

	if active {
		id := tile.ActiveInstanceId

		if _, exists := entity.GetEntityState(id); exists {
			entity.DeleteEntityInstance(id)
		}

		tile.ActiveTypeId = 0
		tile.ActiveInstanceId = [4]byte{}
	} else {
		tile.PassiveTypeId = 0
	}
}

func DestroyActiveEntity(id entity.InstanceId) {
	state, exists := entity.GetEntityState(id)

	if !exists {
		return
	}

	x, y := state.GetPos()
	tile := &Tiles[y][x]

	entity.DeleteEntityInstance(id)
	tile.ActiveTypeId = 0
	tile.ActiveInstanceId = [4]byte{}
}

func GetPassiveEntityAt(x, y int) entity.TypeId {
	return Tiles[y][x].PassiveTypeId
}

func GetActiveEntityAt(x, y int) (*entity.EntityInfo, bool) {
	tile := Tiles[y][x]
	state, exists := entity.GetEntityState(tile.ActiveInstanceId)
	entity := entity.EntityInfo{
		TypeId: tile.ActiveTypeId,
		X:      x,
		Y:      y,
		State:  state,
	}

	return &entity, exists
}

func MoveEntity(fromX, fromY, toX, toY int, active bool) bool {
	if toX < 0 || toY < 0 || toX >= WorldSize || toY >= WorldSize {
		return false
	}

	fromTile := &Tiles[fromY][fromX]
	toTile := &Tiles[toY][toX]

	if active {
		state, exists := entity.GetEntityState(fromTile.ActiveInstanceId)
		var overlapping entity.Overlapping

		if exists {
			if state.GetOverlapped() {
				return false
			}

			overlapping = state.GetOverlapping()

			if toTile.ActiveTypeId != 0 {
				zIndex := state.GetZIndex()
				toState, toExists := entity.GetEntityState(toTile.ActiveInstanceId)
				var overlapable bool
				var toZIndex byte

				if toExists {
					overlapable = toState.GetOverlappable()
					toZIndex = toState.GetZIndex()
				} else {
					overlapable = entity.Overlapables[toTile.ActiveTypeId]
					toZIndex = entity.ZIndexes[toTile.ActiveTypeId]
				}

				if !overlapable || zIndex <= toZIndex {
					return false
				}

				overState, overExists := entity.GetEntityState(overlapping.InstanceId)

				if overExists {
					overState.SetOverlapped(false)
				}

				if toExists {
					toState.SetOverlapped(true)
				}
			}

			state.SetPos(toX, toY)
		} else if toTile.ActiveTypeId != 0 {
			return false
		}

		state.SetOverlapping(toTile.ActiveTypeId, toTile.ActiveInstanceId)
		toTile.ActiveTypeId = fromTile.ActiveTypeId
		toTile.ActiveInstanceId = fromTile.ActiveInstanceId
		fromTile.ActiveTypeId = overlapping.TypeId
		fromTile.ActiveInstanceId = overlapping.InstanceId
	} else {
		if toTile.PassiveTypeId != 0 {
			return false
		}

		toTile.PassiveTypeId = fromTile.PassiveTypeId
		fromTile.PassiveTypeId = 0
	}

	return true
}

func CreateDepthTiles(width, height int) int {
	numDepthTiles++
	DepthTiles[numDepthTiles] = make([][]EntityTile, height)

	for i := 0; i < height; i++ {
		DepthTiles[numDepthTiles][i] = make([]EntityTile, width)
	}

	return numDepthTiles
}

func Generate() {
	Generator()
	engine.AddView("world", view)
}
