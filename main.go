package main

import (
	"log"
	"net/http"
)

func main() {
	router := makeRouter()
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatalln("Connot start server")
	}
}
