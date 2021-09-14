package firstapp

import (
	"time"

	"github.com/GalBrandwine/GoLearning/first_go_app/main.go/internal/logger"
	web "github.com/GalBrandwine/GoLearning/first_go_app/main.go/web"
)

// Init initiate app
func Init() {
	logger.LogInfo("Hello from my first GO app")
	logger.LogInfo("The time is: " + time.Now().String())
	logger.LogInfo("Initiating server...")
	go web.Init()
	logger.LogInfo("Server initiated...")

	logger.LogInfo("Entering maintainance loop...")
	count := 100
	for i := 0; i < count; i++ {
		logger.LogInfo("Loop...")
		time.Sleep(time.Duration(time.Second))
	}

	// p1 := &web.Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	// p1.Save()
	// p2, _ := web.LoadPage("TestPage")
	// logger.LogInfo(p2.Title)
}
