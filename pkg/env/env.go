package env

import (
	"os"
)

// EnvironmentName is the name of the environment variable which describes the environment
var EnvironmentName = "GO_ENV"

// IsProduction returns if the environment is in production
func IsProduction() bool {
	return os.Getenv(EnvironmentName) == "Production"
}
