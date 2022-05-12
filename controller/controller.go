package controller

import "log"

type Controller struct {
	Name string `json:"name"`
}

func (c *Controller) Run(params map[string]string) (string, error) {
	log.Println("Controller:", c.Name)
	log.Println("Params:", params)
	return "", nil
}
