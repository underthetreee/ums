package main

import (
	"log"

	"github.com/underthetreee/ums/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
