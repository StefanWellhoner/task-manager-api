package main

import "github.com/StefanWellhoner/task-manager-api/internal/app"

func main() {
	router := app.SetupRouter()
	router.Run(":80")
}
