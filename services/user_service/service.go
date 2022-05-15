package user_service

import "github.com/bludot/gorouter/core/service"

type User struct {
	Name string
	Age  int
}

type UserService struct {
	service.ServiceService
	Users map[string]*User
}

var UserServiceInstance *UserService

func (s *UserService) GetUser(name string) *User {
	return s.Users[name]
}

func (s *UserService) GetUsers() map[string]*User {
	return s.Users
}

func (s *UserService) AddUser(name string, age int) {
	s.Users[name] = &User{
		Name: name,
		Age:  age,
	}
}

func NewUserService() *UserService {
	userService := service.ServiceService{
		Name: "user_service",
	}
	return &UserService{
		userService,
		make(map[string]*User),
	}
}

func GetUserService() *UserService {
	if UserServiceInstance == nil {
		UserServiceInstance = NewUserService()
	}
	return UserServiceInstance
}
