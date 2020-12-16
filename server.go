package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./assets")))
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}
