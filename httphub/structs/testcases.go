package structs

type HTTPMethodsTestCase struct {
	Name    string
	Args    map[string][]string
	Headers map[string][]string
	Body    interface{}
	Form    map[string][]string
}
