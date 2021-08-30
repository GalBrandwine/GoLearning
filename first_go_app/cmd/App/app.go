package firstapp

import (
	"time"

	"github.com/GalBrandwine/GoLearning/first_go_app/main.go/internal/logger"
	"github.com/GalBrandwine/GoLearning/first_go_app/main.go/web"
)

// Init initiate app
func Init() {
	logger.LogInfo("Hello from my first GO app")
	logger.LogInfo("The time is: " + time.Now().String())
	logger.LogInfo("Initiating server...")
	web.Init()
}
