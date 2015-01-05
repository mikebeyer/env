package env

import "os"

func Getenv(env string, def string) string {
	val := os.Getenv(env)
	if val == "" {
		val = def
	}
	return val
}
