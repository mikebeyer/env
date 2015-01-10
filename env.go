package env

import "os"

func Get(key string, def string) string {
	val := os.Getenv(key)
	if val == "" {
		val = def
	}
	return val
}

func Set(key string, val string) error {
	return os.Setenv(key, val)
}
