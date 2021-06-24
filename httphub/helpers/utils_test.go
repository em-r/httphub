package helpers

import (
	"math/rand"
	"net/http"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestFlatten(t *testing.T) {
	assert := assert.New(t)
	type mapString map[string][]string
	type mapGeneric map[string]interface{}
	type testCase struct {
		desc      string
		baseMap   mapString
		flattened map[string]interface{}
	}

	tcs := []testCase{
		{
			desc:      "one element to be flattened",
			baseMap:   mapString{"x": {"a"}, "y": {"b", "c"}},
			flattened: mapGeneric{"x": "a", "y": []string{"b", "c"}},
		},
		{
			desc:      "all elements to be flattened",
			baseMap:   mapString{"x": {"a"}, "y": {"b"}},
			flattened: mapGeneric{"x": "a", "y": "b"},
		},
		{
			desc:      "shouldn't be flattened",
			baseMap:   mapString{"x": {"a", "b"}, "y": {"c", "d"}},
			flattened: mapGeneric{"x": []string{"a", "b"}, "y": []string{"c", "d"}},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			assert.True(reflect.DeepEqual(Flatten(tc.baseMap), tc.flattened))
		})
	}

}

func TestRandomStr(t *testing.T) {
	assert := assert.New(t)
	r := rand.Intn(30)
	assert.Len(RandomStr(r), r)
}

func TestWalkRouter(t *testing.T) {
	assert := assert.New(t)
	handler := mux.NewRouter()
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {}
	handler.HandleFunc("/valid", handlerFunc).Methods("GET")
	handler.HandleFunc("/root/sub", handlerFunc).Methods("GET")
	handler.HandleFunc("/exclude", handlerFunc).Methods("GET")
	handler.HandleFunc("/post", handlerFunc).Methods("POST")

	var paths []string
	handler.Walk(WalkRouterGET(&paths, true, "exclude"))
	assert.Len(paths, 1)
	assert.Equal(paths[0], "/valid")
}
