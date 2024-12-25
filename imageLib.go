package main

import (
	"log"
	"os"
)

type ImageLib struct {
	bladeFiles []string
	hiltFiles  []string
}

func getImageFiles(imageEntries []os.DirEntry) []string {
	var imageFiles []string
	for _, entry := range imageEntries {
		if !entry.IsDir() {
			imageFiles = append(imageFiles, entry.Name())
		}
	}
	return imageFiles
}

func GetImageLib() *ImageLib {
	bladeEntries, err := os.ReadDir("images/blades")
	if err != nil {
		log.Fatal(err)
	}
	var bladeFiles []string = getImageFiles(bladeEntries)

	hiltEntries, err := os.ReadDir("images/hilts")
	if err != nil {
		log.Fatal(err)
	}
	var hiltFiles []string = getImageFiles(hiltEntries)

	return &ImageLib{
		bladeFiles: bladeFiles,
		hiltFiles:  hiltFiles,
	}
}
