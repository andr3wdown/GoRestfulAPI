package main

import (
	"fmt"
	"net/http"
)

func HandleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func HandleProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Profile Page")
}

func HandleSword(w http.ResponseWriter, r *http.Request) {
	GenerateSwordImage()
	fmt.Fprintf(w, "Sword Page")
}
