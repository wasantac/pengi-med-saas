package config

import (
	"os"
	"strconv"
	"strings"
)

func GetNumberEnv(env string) (int64, error) {
	return strconv.ParseInt(os.Getenv(env), 10, 64)
}

// GetEnv retrieves the value of the environment variable named by env.
func GetEnv(env string) string {
	return os.Getenv(env)
}

func GetEnvWithDefault(env, defaultValue string) string {
	val := os.Getenv(env)
	if val == "" {
		return defaultValue
	}
	return val
}

func GetBoolEnv(env string) (bool, error) {
	val := os.Getenv(env)
	if val == "" {
		return false, nil
	}
	return strconv.ParseBool(strings.ToLower(val))
}
