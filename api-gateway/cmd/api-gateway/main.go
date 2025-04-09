package main

import "api-gateway/internal/app"

func main() {
	app := app.NewApp()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
