package common

import (
	"os"
)

// Getenv get environtment variable, if not found use fallback
func Getenv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}
