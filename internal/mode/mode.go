package mode

import "github.com/gin-gonic/gin"

const (
	// Env is the production mode.
	Prod = "production"
	// Env is the development mode.
	Dev = "development"
	// Env is the test mode.
	Test = "test"
)

// Env is the current Env of the application.
var Env = Dev

// SetEnv sets the Env of the application.
func SetEnv(mode string) {
	Env = mode
	updateGinEnv()
}

// Get returns the current mode of the application.
func GetEnv() string {
	return Env
}

// IsProduction returns true if the application is in production mode.
func IsProduction() bool {
	return Env == Prod
}

// IsDevelopment returns true if the application is in development mode.
func IsDevelopment() bool {
	return Env == Dev
}

// IsTest returns true if the application is in test mode.
func IsTest() bool {
	return Env == Test
}

// IsDebug returns true if the application is in development or test mode.
func IsDebug() bool {
	return IsDevelopment() || IsTest()
}

func updateGinEnv() {
	switch GetEnv() {
	case Prod:
		gin.SetMode(gin.ReleaseMode)
	case Dev:
		gin.SetMode(gin.DebugMode)
	case Test:
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}
