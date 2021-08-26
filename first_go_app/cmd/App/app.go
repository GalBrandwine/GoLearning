package firstapp

import (
	"fmt"
	"time"

	"github.com/GalBrandwine/GoLearning/first_go_app/main.go/web"
)

// Init initiate app
func Init() {
	fmt.Println("Hello from my first GO app")
	fmt.Println("The time is: " + time.Now().String())
	fmt.Println("Initiating server...")
	web.Init()
}
