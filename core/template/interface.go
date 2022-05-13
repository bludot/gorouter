package template

type TemplateEngine interface {
	Render(name string, data interface{}) (string, string, error)
}
