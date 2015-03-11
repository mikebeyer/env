package env

import (
	"os"
	"strconv"
)

func Get(key string, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}

func Int(key string, def int) int {
	val := os.Getenv(key)
	if val == "" {
		return def
	}

	i, err := strconv.Atoi(val)
	if err != nil {
		return def
	}
	return i
}

func Bool(key string, def bool) bool {
	val := os.Getenv(key)
	if val == "" {
		return def
	}

	b, err := strconv.ParseBool(val)
	if err != nil {
		return def
	}
	return b
}

func Float(key string, def float64, bit int) float64 {
	val := os.Getenv(key)
	if val == "" {
		return def
	}

	f, err := strconv.ParseFloat(val, bit)
	if err != nil {
		return def
	}
	return f
}
