package engine

import (
	"embed"
	"fmt"
	"log"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

//go:embed fonts/*
var fonts embed.FS

// //go:embed images/*
// var images embed.FS

var faces = map[string]*font.Face{}

func LoadFont(name string, size float64) *font.Face {
	bytes, err := fonts.ReadFile("fonts/" + name + ".ttf")

	if err != nil {
		log.Fatal(err)
	}

	font, err := truetype.Parse(bytes)

	if err != nil {
		log.Fatal(err)
	}

	face := truetype.NewFace(font, &truetype.Options{Size: size})
	name += fmt.Sprintf("_%f", size)
	faces[name] = &face
	return &face
}

func GetFont(name string, size float64) *font.Face {
	name += fmt.Sprintf("_%f", size)
	return faces[name]
}
