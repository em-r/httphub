package helpers

import (
	"os"
	"strconv"
)

func IsDevMode() bool {
	env, ok := os.LookupEnv("DEV_MODE")
	if !ok {
		return false
	}

	if env == "true" {
		return true
	}

	envInt, err := strconv.Atoi(env)
	if err != nil {
		return false
	}

	return envInt > 0
}

// Flatten returns a map based on the passed `multi` param. If the slice for a key
// has only one value the key in the returned map will have that plain value
// instead of the slice.
func Flatten(multi map[string][]string) map[string]interface{} {
	flattened := make(map[string]interface{})

	for key, val := range multi {
		if len(val) == 1 {
			flattened[key] = val[0]
		} else {
			flattened[key] = val
		}
	}

	return flattened
}
