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

func EnvironmentIsProd() bool {
	if GetEnvironment() == prodEnvName {
		return true
	}
	return false
}

func EnvironmentIsDev() bool {
	if GetEnvironment() == devEnvName {
		return true
	}
	return false
}

func CheckEnvironmentIsProd(environment string) bool {
	if environment == prodEnvName {
		return true
	}
	return false
}

func CheckEnvironmentIsDev(environment string) bool {
	if environment == devEnvName {
		return true
	}
	return false
}
