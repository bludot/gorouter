package renderer

import (
	"github.com/bludot/gorouter/core/http_handler"
	"github.com/bludot/gorouter/core/template"
	"github.com/bludot/gorouter/core/transformer"
	"net/http"
)

type Render struct {
	templateEngine template.TemplateEngine
	transformer    transformer.Transformer
	Response       *http_handler.HTTPResponse
}

var renderer *Render

func (r *Render) Render(templateName string, data interface{}, statusCode int) (*http_handler.HTTPResponse, error) {

	rawString, RenderType, err := r.templateEngine.Render(templateName, data)
	r.Response = &http_handler.HTTPResponse{
		Header: map[string][]string{
			"Content-Type": {RenderType},
		},
		Body:       []byte(rawString),
		StatusCode: statusCode,
	}
	return r.Response, err
}

func (r *Render) ToJSON(data interface{}, statusCode int) (*http_handler.HTTPResponse, error) {
	jsonString, err := r.transformer.ToJson(data)
	if err != nil {
		return nil, err
	}
	r.Response = &http_handler.HTTPResponse{
		Header: map[string][]string{
			"Content-Type": {"application/json"},
		},
		Body:       []byte(jsonString),
		StatusCode: statusCode,
	}
	return r.Response, nil

}

func (r *Render) SetTemplateEngine(templateEngine template.TemplateEngine) *Render {
	r.templateEngine = templateEngine
	return r
}

func (r *Render) SetTransformer(transformer transformer.Transformer) *Render {
	r.transformer = transformer
	return r
}

func (r *Render) Respond(w http.ResponseWriter, req *http.Request) {
	if r.Response == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(r.Response.Body)
		return
	}
	for k, v := range r.Response.Header {
		w.Header().Set(k, v[0])
	}
	w.WriteHeader(r.Response.StatusCode)
	w.Write(r.Response.Body)
}

func (r *Render) Clear() {
	r.Response = nil
}

func GetRender() *Render {
	if renderer == nil {
		renderer = &Render{}
	}
	return renderer
}
