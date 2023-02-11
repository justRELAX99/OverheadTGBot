package pkg

import "os"

const (
	globalEnvVariableName = "ENV_NAME"
	devEnvName            = "dev"
	prodEnvName           = "prod"
)

func GetEnvironment() string {
	return os.Getenv(globalEnvVariableName)
}

func IsProd() bool {
	return GetEnvironment() == prodEnvName
}

func IsDev() bool {
	return GetEnvironment() == devEnvName
}
