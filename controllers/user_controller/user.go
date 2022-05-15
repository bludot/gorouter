package user_controller

import (
	"context"
	"encoding/json"
	"github.com/bludot/gorouter/core/controller"
	"github.com/bludot/gorouter/core/renderer"
	"github.com/bludot/gorouter/core/router/entities"
	"github.com/bludot/gorouter/core/service"
	"github.com/bludot/gorouter/services/user_service"
	"log"
	"net/http"
)

type UserDTO struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserController struct {
	controller.Controller
}

func (c *UserController) GetUser(ctx context.Context, r entities.HTTPRequest) error {
	log.Println("Handler:", c.Name)
	log.Println("Params:", r.Params)
	service := *(c.Controller.GetService("user_service"))
	user := service.(*user_service.UserService).GetUser(r.Params.Get("name"))

	if user == nil {
		renderer.GetRender().ToJSON(map[string]string{"error": "User not found"}, http.StatusNotFound)
		return nil
	}

	renderer.GetRender().ToJSON(user, http.StatusOK)
	return nil
}
func (c *UserController) AddUser(ctx context.Context, r entities.HTTPRequest) error {
	log.Println("Handler:", c.Name)
	log.Println("Params:", r.Params)
	var userDTO UserDTO
	if r.Body == nil {
		renderer.GetRender().ToJSON(map[string]string{"error": "Body is empty"}, http.StatusBadRequest)
		return nil
	}
	err := json.Unmarshal(*r.Body, &userDTO)
	if err != nil {
		renderer.GetRender().ToJSON(map[string]string{"error": "Invalid request"}, http.StatusBadRequest)
		return nil
	}
	service := *(c.Controller.GetService("user_service"))
	service.(*user_service.UserService).AddUser(userDTO.Name, userDTO.Age)

	log.Println(service.(*user_service.UserService).GetUsers())
	renderer.GetRender().ToJSON(map[string]interface{}{
		"message": "User added",
	}, http.StatusOK)
	return nil
}

var userController *UserController

func Controller(services ...service.Service) *UserController {
	if userController == nil {

		thisController := &controller.Controller{
			Name:     "UserController",
			Services: services,
		}
		userController = &UserController{
			Controller: *thisController,
		}
	}
	return userController
}
