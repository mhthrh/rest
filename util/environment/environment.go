package environment

import "os"

const (
	appRootDir = "myAppRootDir"
)

func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}
func SetEnv(key, value string) error {
	return os.Setenv(key, value)
}
func GetAppPath() string {
	path := os.Getenv(appRootDir)
	if path == "" {
		return os.Getenv("GOHOME")
	}
	return path
}
