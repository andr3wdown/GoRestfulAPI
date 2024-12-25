package main

import (
	"fmt"
	"image/color"
	"net/http"
)

var imageLib *ImageLib
var mainColors []color.RGBA
var materialLib *ColorLib

func InitLibs() {
	imageLib = GetImageLib()
	mainColors = GetMainColors()
	materialLib = GetMaterialLib()
}

func HandleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func HandleProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Profile Page")
}

func HandleSword(w http.ResponseWriter, r *http.Request) {
	if imageLib == nil {
		InitLibs()
	}
	GenerateSwordImage(imageLib, mainColors, materialLib)
	fmt.Fprintf(w, "Sword Page")
}

func HandleTest(w http.ResponseWriter, r *http.Request) {
	if imageLib == nil {
		InitLibs()
	}
	var composed string = ""
	for _, image := range imageLib.bladeFiles {
		composed += image + "\n"
	}

	fmt.Fprintf(w, "%s", composed)
}
