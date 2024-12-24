package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math/rand"
	"os"
)

func processPart(img image.Image, colors []color.RGBA, indexes []int, materialLib [][]color.RGBA) (image.Image, error) {
	var width int = img.Bounds().Dx()
	var height int = img.Bounds().Dy()
	var newImg *image.RGBA = image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			var r, g, b, a = img.At(x, y).RGBA()
			var color color.RGBA = color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}
			switch {
			case color == colors[0]:
				newImg.Set(x, y, materialLib[0][indexes[0]])
			case color == colors[1]:
				newImg.Set(x, y, materialLib[1][indexes[1]])
			case color == colors[2]:
				newImg.Set(x, y, materialLib[2][indexes[2]])
			default:
				newImg.Set(x, y, color)
			}
		}
	}
	return newImg, nil
}

func GenerateSwordImage() string {

	bladeFile, err := os.Open("images/blades/blade-test.png")
	if err != nil {
		panic(err)
	}
	defer bladeFile.Close()

	bladeImg, err := png.Decode(bladeFile)
	if err != nil {
		panic(err)
	}

	hiltFile, err := os.Open("images/hilts/hilt-test.png")
	if err != nil {
		panic(err)
	}
	defer hiltFile.Close()

	hiltImg, err := png.Decode(hiltFile)
	if err != nil {
		panic(err)
	}
	var mainColors []color.RGBA = GetMainColors()
	var materialLib *ColorLib = GetMaterialLib()

	bladeImg, err = processPart(
		bladeImg,
		mainColors,
		[]int{rand.Intn(len(materialLib.metalColors)), rand.Intn(len(materialLib.shadowMetalColors)), rand.Intn(len(materialLib.metalColors))},
		[][]color.RGBA{materialLib.metalColors, materialLib.shadowMetalColors, materialLib.metalColors},
	)
	if err != nil {
		log.Fatal(err)
	}

	hiltImg, err = processPart(
		hiltImg,
		mainColors,
		[]int{rand.Intn(len(materialLib.wrapColors)), rand.Intn(len(materialLib.metalColors)), rand.Intn(len(materialLib.gemColors))},
		[][]color.RGBA{materialLib.wrapColors, materialLib.metalColors, materialLib.gemColors},
	)
	if err != nil {
		log.Fatal(err)
	}

	var width int = bladeImg.Bounds().Dx()
	var height int = bladeImg.Bounds().Dy() + hiltImg.Bounds().Dy()
	swordImg := image.NewRGBA(image.Rect(0, 0, width, height))

	draw.Draw(swordImg, bladeImg.Bounds(), bladeImg, image.Point{0, 0}, draw.Src)
	draw.Draw(swordImg, hiltImg.Bounds().Add(image.Point{0, bladeImg.Bounds().Dy()}), hiltImg, image.Point{0, 0}, draw.Src)

	outFile, err := os.Create("images/output/sword.png")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	png.Encode(outFile, swordImg)

	return "images/output/sword.png"
}
