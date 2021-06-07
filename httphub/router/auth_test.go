package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ElMehdi19/httphub/httphub/structs"
	"github.com/stretchr/testify/assert"
)

func TestViewAuth(t *testing.T) {
	assert := assert.New(t)
	user, passwd := "mehdi", "whatever"

	type testCase struct {
		id           string
		user, passwd string
		shouldFail   bool
	}

	tcs := []testCase{
		{
			id:     "valid auth",
			user:   user,
			passwd: passwd,
		},
		{
			id:         "wrong creds",
			user:       "wrong-username",
			passwd:     "wrong-passwd",
			shouldFail: true,
		},
		{
			id:         "creds not provided",
			shouldFail: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.id, func(t *testing.T) {
			req, err := http.NewRequest("GET", "http://127.0.0.1:5000", nil)
			assert.NoError(err)
			if tc.user != "" && tc.passwd != "" {
				req.SetBasicAuth(tc.user, tc.passwd)
			}

			rec := httptest.NewRecorder()
			viewBasicAuth(rec, req, user, passwd)

			if tc.shouldFail {
				assert.Equal(http.StatusUnauthorized, rec.Result().StatusCode)
				return
			}

			assert.Equal(http.StatusOK, rec.Result().StatusCode)

			var body structs.AuthResponse
			err = json.NewDecoder(rec.Body).Decode(&body)

			assert.NoError(err)
			assert.True(body.Authorized)
			assert.Equal(body.User, user)
		})
	}
}
