package helpers

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

func Choose(s []string) string {
	return s[rand.Intn(len(s))]
}

func WalkRouterGET(paths *[]string, top bool, exclude string) mux.WalkFunc {
	return func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		p, err := route.GetPathTemplate()
		if err != nil {
			return nil
		}

		if top && strings.Count(p, "/") > 1 {
			return nil
		}

		if exclude != "" && strings.Contains(p, exclude) {
			return nil
		}

		methods, err := route.GetMethods()
		if err != nil || len(methods) > 1 || methods[0] != "GET" {
			return nil
		}

		*paths = append(*paths, p)
		return nil
	}
}
