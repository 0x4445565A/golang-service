package main

import (
	"github.com/0x4445565a/golang-service/pkg/api"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	service := api.Service{}
	e := service.Init()

	e.Use(middleware.Logger())

	e.Start(":8080")
}
