package configs

// AppConfig holds all application-wide configurations
var AppConfig = struct {
	ServerPort string
}{
	ServerPort: ":8080",
}
