package server

import (
	"log"
	"net/http"
)

func Run() {
	log.Println("-~- starting API v0.1 -~-")
	log.Printf("[ ~ ] starting server 127.0.0.1: on port 8080")
	log.Panic(http.ListenAndServe(":8080", nil))
	http.HandleFunc("/", Shortener)
}
