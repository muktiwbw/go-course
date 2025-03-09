package env

import (
	"os"
	"strconv"
)

func GetString(key string, fallback string) string {
	value, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	return value
}

func GetInt(key string, fallback int) int {
	value, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	valInt, err := strconv.Atoi(value)

	if err != nil {
		return fallback
	}

	return valInt
}

func GetBool(key string, fallback bool) bool {
	value, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	valBool, err := strconv.ParseBool(value)

	if err != nil {
		return fallback
	}

	return valBool
}
