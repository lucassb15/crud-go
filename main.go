package main

import (
	"log"
	"modulo/handlers"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	http.HandleFunc("/", handlers.UserHandler)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}

}
