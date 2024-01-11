package main

import (
	"fmt"

	"github.com/StefanWellhoner/task-manager-api/internal/app"
	"github.com/StefanWellhoner/task-manager-api/internal/config"
	"github.com/StefanWellhoner/task-manager-api/internal/mode"
	model "github.com/StefanWellhoner/task-manager-api/internal/models"
	"github.com/StefanWellhoner/task-manager-api/internal/services"
)

var (
	// Version is the version of the application.
	Version = "unknown"
	// Commit is the commit of the application.
	Commit = "unknown"
	// BuildTime is the build time of the application.
	BuildTime = "unknown"
	// Mode the build mode of the application.
	Mode = mode.Dev
)

// main is the entry point of the application.
// It sets up the router and starts the server.

func main() {
	vInfo := &model.VersionInfo{Version: Version, Commit: Commit, BuildTime: BuildTime}
	mode.SetEnv(Mode)

	fmt.Println("Starting Task Manager API version", vInfo.Version, "in", mode.GetEnv(), "mode.")
	conf := config.Get()

	db, err := services.New(conf.Database.Database)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := app.SetupRouter(db)
	router.Run()
}
