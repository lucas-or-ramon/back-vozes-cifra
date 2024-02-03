package main

import (
	"back-vozes-cifras/internal/app"
	"back-vozes-cifras/internal/app/config"
)

func init() {
	config.Initialize()
}

func main() {
	app.StartServer()
}
