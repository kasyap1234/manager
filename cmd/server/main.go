package main

import (
	"fmt"
	"log"

	"manager/app"
)

func main() {
	app, err := app.New()
	if err != nil {
		panic(err)
	}

	if err := app.Run(); err != nil {
		log.Printf("error running the app: %v", err)
	}
	fmt.Print("app started")
}
