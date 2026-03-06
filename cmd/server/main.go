package main

import "manager/app"

func main() {
	app, err := app.New()
	if err != nil {
		panic(err)
	}
	app.Run()
}
