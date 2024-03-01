package main

import (
	"log"

	"github.com/ansxy/nagabelajar-be-go/cmd/app"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
