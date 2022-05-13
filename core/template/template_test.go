package template_test

import (
	"github.com/bludot/gorouter/core/template"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestTemplateService_GenerateHTML(t *testing.T) {
	t.Run("GenerateHTML", func(t *testing.T) {
		a := assert.New(t)
		tpl := template.GetTemplateEngine()
		rawString, err := tpl.GenerateHTML("index.html", map[string]string{
			"body": "Hello World",
		})
		a.NoError(err)
		a.Contains(rawString, "Hello World")
		log.Println(rawString)
	})
	t.Run("GenerateHTML - subfolder support", func(t *testing.T) {
		a := assert.New(t)
		tpl := template.GetTemplateEngine()
		rawString, err := tpl.GenerateHTML("test/index.html", map[string]string{
			"body": "Test",
		})
		a.NoError(err)
		a.Contains(rawString, "Hello Test")
		log.Println(rawString)
	})
}
