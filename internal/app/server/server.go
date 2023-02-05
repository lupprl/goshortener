package server

import (
	"github.com/lupprl/goshortener/internal/app/handlers"
	"log"
	"net/http"
)

func Run() {
	log.Println("-~- starting -~-")
	log.Printf("[ ~ ] starting server 127.0.0.1: on port 8080")
	log.Panic(http.ListenAndServe(":8080", nil))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.Shortener()
	})
}
