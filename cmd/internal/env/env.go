package env

import "os"

func GetString(key, fallback string) string {
	str, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return str
}
