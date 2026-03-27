package helpers

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ParseEnvList(key string) []string {
	val, exists := os.LookupEnv(key)
	if !exists || val == "" {
		return []string{}
	}

	parts := strings.Split(val, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	return parts
}

func GetEnvBool(key string, defaultValue bool) bool {
	val, exists := os.LookupEnv(key)
	if !exists || val == "" {
		return defaultValue
	}

	parsedVal, err := strconv.ParseBool(strings.ToLower(val))
	if err != nil {
		log.Printf("Warning: Could not parse boolean value for %s: %v. Using default: %v", key, err, defaultValue)
		return defaultValue
	}

	return parsedVal
}
