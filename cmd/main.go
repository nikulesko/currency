package main

import (
	"currency/internal/transport/rest/router"

	"log"
)

func main() {
	log.Println("Start...")

	r := router.New()

	r.Logger.Fatal(r.Start("127.0.0.1:8585"))
}
