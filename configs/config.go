package configs

import (
	"os"
)

// AppConfig holds all application-wide configurations
var AppConfig = struct {
	ServerPort string
}{
	ServerPort: getPort(),
}

// getPort checks for the "PORT" environment variable and returns it.
// If not found, it returns the default port ":8080".
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return ":8081"
	}
	return ":" + port
}
