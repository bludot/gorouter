package main

import (
	"github.com/bludot/gorouter/controllers/root"
	"github.com/bludot/gorouter/controllers/test"
	"github.com/bludot/gorouter/controllers/test_params"
	"github.com/bludot/gorouter/core/renderer"
	"github.com/bludot/gorouter/core/router"
	"github.com/bludot/gorouter/core/template"
	"github.com/bludot/gorouter/core/transformer"
	"net/http"
)

func main() {
	renderer.
		GetRender().
		SetTemplateEngine(template.GetTemplateEngine()).
		SetTransformer(transformer.GetTransformer())
	mainRouter := router.NewRouter()
	mainRouter.AddRoute(router.Route{
		Controller: root.NewRootController(),
		Path:       "/",
	})
	mainRouter.AddRoute(router.Route{
		Controller: test.NewTestController(),
		Path:       "/test/$id",
	})
	mainRouter.AddRoute(router.Route{
		Controller: test_params.NewTestParamsController(),
		Path:       "/test/test",
	})
	http.Handle("/", mainRouter)
	http.ListenAndServe(":8080", nil)
}
