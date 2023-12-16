package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Service struct{}

func (s *Service) Init() *echo.Echo {
	e := echo.New()

	// Map paths to handlers
	e.GET("/", s.GetRoot)

	return e
}

func (s *Service) GetRoot(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello World")
}