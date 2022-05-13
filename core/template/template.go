package template

import (
	"github.com/bludot/gorouter/templates"
	"io/fs"
	"log"
	"strings"
)

type TemplateService struct {
}

func (t *TemplateService) GenerateHTML(templateName string, data interface{}) (string, error) {
	sub, err := fs.Sub(templates.TemplateFiles, "root")
	_, err = sub.Open(templateName)
	if err != nil {
		log.Println("Error opening template: ", err)
		return "", err
	}
	rawTemplateByte, _ := fs.ReadFile(sub, templateName)
	rawTemplate := string(rawTemplateByte)
	rawString := ""
	dataMap := data.(map[string]string)
	for k, v := range dataMap {
		log.Println(k, v)
		rawString = strings.Replace(rawTemplate, "{"+k+"}", v, -1)
	}
	return rawString, nil
}

func (t *TemplateService) Render(name string, data interface{}) (string, string, error) {
	rawString, err := t.GenerateHTML(name, data)
	if err != nil {
		return "", "", err
	}
	return rawString, "text/html", nil
}

var templateEngine *TemplateService

func GetTemplateEngine() *TemplateService {
	if templateEngine == nil {
		templateEngine = &TemplateService{}
	}
	return templateEngine
}
