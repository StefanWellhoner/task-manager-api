package main

import "github.com/StefanWellhoner/task-manager-api/internal/app"

// main is the entry point of the application.
// It sets up the router and starts the server.
func main() {
	router := app.SetupRouter()
	router.Run()
}
