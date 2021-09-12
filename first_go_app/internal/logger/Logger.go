package logger

import (
	"fmt"
	"log"
)

var loggingSessionsCounter int

func printLoggingSessions() {
	fmt.Printf("Logging sessions: %d\n", loggingSessionsCounter)
}
func format() {
	fmt.Println("This is un-exported format function")
}

// LogInfo used for loggin INFO events
func LogInfo(message string) {
	log.Print(message)
	// fmt.Println(message)
	// format()
	// loggingSessionsCounter++
	// printLoggingSessions()
}

// LogFatalString logs fatal messages
func LogFatalString(message string) {
	log.Fatal(message)
}

// LogFatalError logs fatal messages
func LogFatalError(error1 error) {
	log.Fatal(error1.Error())
}
