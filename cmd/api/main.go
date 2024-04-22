package main

import (
	"fmt"

	"github.com/StefanWellhoner/task-manager-api/internal/app"
	"github.com/StefanWellhoner/task-manager-api/internal/config"
	"github.com/StefanWellhoner/task-manager-api/internal/mode"
	model "github.com/StefanWellhoner/task-manager-api/internal/models"
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

	fmt.Println("Config is set in mode", mode.GetEnv())
	fmt.Println("Listening on", conf.Server.ListenAddr, "port", conf.Server.Port)

	router := app.SetupRouter()
	router.Run(fmt.Sprintf("%s:%d", conf.Server.ListenAddr, conf.Server.Port))
}
