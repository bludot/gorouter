package service

type ServiceService struct {
	Name string
}

func (s *ServiceService) GetName() string {
	return s.Name
}
