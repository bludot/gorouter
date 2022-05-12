package controller

type IController interface {
	Run(params map[string]string) (string, error)
}
