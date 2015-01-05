package env

import "os"

func Get(env string, def string) string {
	val := os.Getenv(env)
	if val == "" {
		val = def
	}
	return val
}
