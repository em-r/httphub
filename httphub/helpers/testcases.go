package helpers

import (
	"encoding/json"
	"net/url"

	"github.com/ElMehdi19/httphub/httphub/structs"
)

var HTTPMethodsBaseTc = structs.HTTPMethodsTestCase{
	Args: map[string][]string{"x": {"1", "2"}, "y": {"3"}},
	Headers: map[string][]string{
		"scranton": {"bears", "beats", "battlestar galactica"},
		"whomai":   {"mehdi"},
	},
}

var HTTPMethodsTcs = []structs.HTTPMethodsTestCase{
	{
		Name:    "with json",
		Args:    HTTPMethodsBaseTc.Args,
		Headers: HTTPMethodsBaseTc.Headers,
		JSON: map[string]interface{}{
			"bool": true,
			"int":  1,
			"str":  "whatever",
		},
		ContentType: "application/json",
	},
	{
		Name:    "with form",
		Args:    HTTPMethodsBaseTc.Args,
		Headers: HTTPMethodsBaseTc.Headers,
		Form: map[string][]string{
			"user":    {"me"},
			"message": {"whatever"},
		},
		ContentType: "application/x-www-form-urlencoded",
	},
	{
		Name:        "with binary",
		Args:        HTTPMethodsBaseTc.Args,
		Headers:     HTTPMethodsBaseTc.Headers,
		Data:        "\x01\x02\x03\x04\x05",
		ContentType: "application/octet-stream",
	},
	{
		Name:    "with text",
		Args:    HTTPMethodsBaseTc.Args,
		Headers: HTTPMethodsBaseTc.Headers,
		Data:    "some random stuff",
	},
}

func MakeBodyFromTestCase(tc structs.HTTPMethodsTestCase) []byte {
	var b []byte
	switch tc.ContentType {
	case "application/json":
		b, _ = json.Marshal(tc.JSON)
	case "application/x-www-form-urlencoded":
		b = []byte(url.Values(tc.Form).Encode())
	default:
		// plain/text
		b = []byte(tc.Data.(string))
	}
	return b
}
