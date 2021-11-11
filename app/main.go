package main

import (

	"github.com/nerdcademy/restapi/controller"
	"github.com/nerdcademy/restapi/model"
)

func main() {
	model.Init()
	controller.Start()
}