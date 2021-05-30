package structs

// HTTPMethodsResponse represents the response
// that will be sent back to clients requesting
// any of the http_methods endpoint.
type HTTPMethodsResponse struct {
	URL     string                 `json:"url"`
	Args    map[string]interface{} `json:"args,omitempty"`
	Headers map[string]interface{} `json:"headers"`
	Origin  string                 `json:"origin,omitempty"`
	Form    map[string]interface{} `json:"form,omitempty"`
	JSON    interface{}            `json:"json,omitempty"`
	Data    interface{}            `json:"data,omitempty"`
	Method  string                 `json:"method,omitempty"`
}
