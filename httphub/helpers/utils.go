package helpers

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var PORT int = 5000
var HOST string

func init() {
	HOST = fmt.Sprintf("127.0.0.1:%d", PORT)
	// if IsDevMode() {
	// 	HOST = fmt.Sprintf("127.0.0.1:%d", PORT)
	// } else {
	// 	HOST = "httphub.io"
	// }
}

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

// Random generates a random string of numbers, uppercase and lowercase letters with
// a fixed length.
func RandomStr(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(chars[rand.Intn(len(chars))])
	}
	return sb.String()
}
