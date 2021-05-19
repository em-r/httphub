package helpers

import "net/url"

// CreateURL creates a url by concatinating
// the base to the query args and returns
// the resulting url as a string.
func CreateURL(base string, args map[string][]string) string {
	u, _ := url.Parse(base)
	q := u.Query()

	for arg, vals := range args {
		for _, val := range vals {
			q.Add(arg, val)
		}
	}

	u.RawQuery = q.Encode()
	return u.String()
}
