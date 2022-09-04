package world

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/r3ndd/go-rogue/app/engine"
	"github.com/r3ndd/go-rogue/app/engine/entity"
	"github.com/r3ndd/go-rogue/app/utils"
	"golang.org/x/image/font"
)

const fontName string = "agave"
const baseGridSize int = 80
const topMargin int = 16
const bottomMargin int = 32
const fontRatioNum int = 3
const fontRatioDenom int = 2

var fontFace font.Face
var viewX int = 0
var viewY int = 0
var gridXSize int
var gridYSize int

func init() {
	fontFace = *engine.LoadFont(fontName, 22)
}

func view(screen *ebiten.Image) error {
	screenWidth, screenHeight := screen.Size()
	viewSize := screenHeight - topMargin - bottomMargin
	sideMargin := (screenWidth - viewSize) / 2
	gridXSize = baseGridSize
	gridYSize = baseGridSize / fontRatioNum * fontRatioDenom
	tileXSize := viewSize / gridXSize
	tileYSize := viewSize / gridYSize
	yExcess := viewSize % gridYSize

	if yExcess >= tileYSize {
		gridYSize += yExcess / tileYSize
	}

	for gridY, row := range Tiles[viewY:utils.IntMin(viewY+gridYSize, WorldSize-1)] {
		for gridX, tile := range row[viewX:utils.IntMin(viewX+gridXSize, WorldSize-1)] {
			var tileRune rune
			var tileColor color.Color

			state, exists := entity.GetEntityState(tile.ActiveInstanceId)

			if exists {
				tileRune = state.GetRune()
				tileColor = state.GetColor()
			} else {
				var typeId entity.TypeId

				if tile.ActiveTypeId != 0 {
					typeId = tile.ActiveTypeId
				} else {
					typeId = tile.PassiveTypeId
				}

				tileRune = entity.Runes[typeId]
				tileColor = entity.Colors[typeId]
			}

			x := gridX*tileXSize + sideMargin
			y := gridY*tileYSize + topMargin
			text.Draw(screen, string(tileRune), fontFace, x, y, tileColor)
		}
	}

	return nil
}

func MoveView(dir string) {
	shiftScale := utils.IntMax(baseGridSize/10, 1)

	switch dir {
	case "left":
		if viewX >= shiftScale {
			viewX -= shiftScale
		} else {
			viewX = 0
		}
	case "right":
		if viewX < WorldSize-gridXSize-shiftScale {
			viewX += shiftScale
		} else {
			viewX = WorldSize - gridXSize
		}
	case "up":
		if viewY >= shiftScale {
			viewY -= shiftScale
		} else {
			viewY = 0
		}
	case "down":
		if viewY < WorldSize-gridYSize-shiftScale {
			viewY += shiftScale
		} else {
			viewY = WorldSize - gridYSize
		}
	}

	fmt.Printf("%d, %d\n", viewX, viewY)
}
