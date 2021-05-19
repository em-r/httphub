package structs

// HTTPMethodsResponse represents the response
// that will be sent back to clients requesting
// any of the http_methods endpoint.
type HTTPMethodsResponse struct {
	URL     string                 `json:"url"`
	Args    map[string][]string    `json:"args,omitempty"`
	Headers map[string][]string    `json:"headers"`
	Origin  string                 `json:"origin,omitempty"`
	JSON    map[string]interface{} `json:"json,omitempty"`
	Method  string                 `json:"method,omitempty"`
}