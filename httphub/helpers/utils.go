package helpers

import (
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

var PORT int = 80

// checking if there's an already set PORT env variable, this is mainly for Heroku deployment.
func init() {
	p, ok := os.LookupEnv("PORT")
	if !ok {
		return
	}

	portInt, err := strconv.Atoi(p)
	if err != nil {
		return
	}

	PORT = portInt
}

// IsDevMode returns true if DEV_MODE is present in the environment variables, and set to a non
// false value.
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

// Choose selects an element randomly from a slice and returns it.
func Choose(s []string) string {
	return s[rand.Intn(len(s))]
}

// WalkRouterGET takes a pointer to an empty slice, a bool and string params returns a mux.WalkFunc
// function used to walk down the mux router.
// The paths param slice will be populated with routers paths that match a specific criterea.
// If top is true, only top level routes will be selected.
// If exclude is not empty, routes that include the exclude value will be skipped.
// WalkRouterGET currently selects only routes that accept just the GET method, hence the name.
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
