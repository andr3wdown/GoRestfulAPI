package main

import (
	"image/color"
	"math/rand"
)

func GetMainColors() []color.RGBA {
	return []color.RGBA{
		{R: 255, G: 0, B: 0, A: 255}, // Red
		{R: 0, G: 255, B: 0, A: 255}, // Green
		{R: 0, G: 0, B: 255, A: 255}, // Blue
	}
}

type ColorLib struct {
	metalColors       []color.RGBA
	shadowMetalColors []color.RGBA
	wrapColors        []color.RGBA
	gemColors         []color.RGBA
}

func GetMaterialLib() *ColorLib {
	return &ColorLib{
		metalColors: []color.RGBA{
			{R: 192, G: 192, B: 192, A: 255}, // Silver
			{R: 255, G: 215, B: 0, A: 255},   // Gold
			{R: 205, G: 127, B: 50, A: 255},  // Bronze
		},
		shadowMetalColors: []color.RGBA{
			{R: 128, G: 128, B: 128, A: 255}, // Gray
			{R: 139, G: 69, B: 19, A: 255},   // Saddle Brown
			{R: 10, G: 10, B: 10, A: 255},    // Dark Grey
		},
		wrapColors: []color.RGBA{
			{R: 255, G: 0, B: 0, A: 255}, // Red
			{R: 0, G: 255, B: 0, A: 255}, // Green
			{R: 0, G: 0, B: 255, A: 255}, // Blue
		},
		gemColors: []color.RGBA{
			{R: 75, G: 0, B: 130, A: 255},    // Indigo
			{R: 238, G: 130, B: 238, A: 255}, // Violet
			{R: 255, G: 20, B: 147, A: 255},  // Deep Pink
		},
	}
}

func RandomColor() color.RGBA {
	return color.RGBA{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
		A: 255,
	}
}
