package structs

type HTTPMethodsTestCase struct {
	Name        string
	Args        map[string][]string
	Headers     map[string][]string
	JSON        interface{}
	Form        map[string][]string
	ContentType string
	Data        interface{}
}
