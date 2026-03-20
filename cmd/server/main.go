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
		log.Fatal("error running the app")
	}
	fmt.Print("app started")
}
