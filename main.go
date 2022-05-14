package main

import (
	"github.com/bludot/gorouter/controllers/root_controller"
	"github.com/bludot/gorouter/controllers/test_controller"
	"github.com/bludot/gorouter/controllers/test_params_controller"
	"github.com/bludot/gorouter/core/renderer"
	"github.com/bludot/gorouter/core/router"
	"github.com/bludot/gorouter/core/router/entities"
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
	mainRouter.AddRoute(entities.Route{
		Handler: root_controller.Controller().Root,
		Path:    "/",
		Method:  http.MethodGet,
	})
	mainRouter.AddRoute(entities.Route{
		Handler: test_controller.NewTestController().TestRoute,
		Path:    "/test/$id",
		Method:  http.MethodGet,
	})
	mainRouter.AddRoute(entities.Route{
		Handler: test_params_controller.NewTestParamsController().TestParams,
		Path:    "/test/test",
		Method:  http.MethodPost,
	})
	http.Handle("/", mainRouter)
	http.ListenAndServe(":8080", nil)
}
