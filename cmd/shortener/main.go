package main

import (
	"github.com/lupprl/goshortener/internal/app/server"
	"log"
)

func main() {
	log.Println("-~- increment -~-")
	server.Run()
}
