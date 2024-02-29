package main

import "johansolbakken.no/weatherdemo/pkg/app"

func main() {
	server := app.Server{}
	server.Run()
}
