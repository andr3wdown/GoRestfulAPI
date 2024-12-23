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

func processBlade(img image.Image, colors []color.RGBA, materials *ColorLib) (image.Image, error) {
	var width int = img.Bounds().Dx()
	var height int = img.Bounds().Dy()
	var newImg *image.RGBA = image.NewRGBA(image.Rect(0, 0, width, height))

	var index1, index2, index3 int = rand.Intn(len(materials.metalColors)), rand.Intn(len(materials.shadowMetalColors)), rand.Intn(len(materials.metalColors))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			var r, g, b, a = img.At(x, y).RGBA()
			var color color.RGBA = color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}
			switch {
			case color == colors[0]:
				newImg.Set(x, y, materials.metalColors[index1])
			case color == colors[1]:
				newImg.Set(x, y, materials.shadowMetalColors[index2])
			case color == colors[2]:
				newImg.Set(x, y, materials.metalColors[index3])
			default:
				newImg.Set(x, y, color)
			}
		}
	}
	return newImg, nil
}
func processHilt(img image.Image, colors []color.RGBA, materials *ColorLib) (image.Image, error) {
	var width int = img.Bounds().Dx()
	var height int = img.Bounds().Dy()
	var newImg *image.RGBA = image.NewRGBA(image.Rect(0, 0, width, height))

	var index1, index2, index3 int = rand.Intn(len(materials.wrapColors)), rand.Intn(len(materials.metalColors)), rand.Intn(len(materials.gemColors))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			var r, g, b, a = img.At(x, y).RGBA()
			var color color.RGBA = color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}
			switch {
			case color == colors[0]:
				newImg.Set(x, y, materials.wrapColors[index1])
			case color == colors[1]:
				newImg.Set(x, y, materials.metalColors[index2])
			case color == colors[2]:
				newImg.Set(x, y, materials.gemColors[index3])
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

	bladeImg, err = processBlade(bladeImg, GetMainColors(), GetMaterialLib())
	if err != nil {
		log.Fatal(err)
	}

	hiltImg, err = processHilt(hiltImg, GetMainColors(), GetMaterialLib())
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
