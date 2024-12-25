package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", HandleMain)
	http.HandleFunc("/profile", HandleProfile)
	http.HandleFunc("/sword", HandleSword)
	http.HandleFunc("/test", HandleTest)

	log.Println("Starting server on :4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
