package environment

import (
	"log"
	"os"
)

func GetEnvOrStop( key string) string {
	value, ok  := os.LookupEnv( key )
	if ok {
		return value
	} else {
		log.Fatalf( "Envirnment variable %s missing", key)
		return ""
	}
}

// Get env var or default
func GetEnvWithFallback(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

