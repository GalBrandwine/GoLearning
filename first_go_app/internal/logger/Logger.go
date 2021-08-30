package logger

import "fmt"

var loggingSessionsCounter int

func printLoggingSessions() {
	fmt.Printf("Logging sessions: %d\n", loggingSessionsCounter)
}
func format() {
	fmt.Println("This is un-exported format function")
}

// LogInfo used for loggin INFO events
func LogInfo(message string) {
	fmt.Println(message)
	format()
	loggingSessionsCounter++
	printLoggingSessions()
}
