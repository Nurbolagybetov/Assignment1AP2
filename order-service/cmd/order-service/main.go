package main

import "order-service/internal/app"

func main() {
	app := app.NewApp()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
