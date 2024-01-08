package service

import (
	"net/http"

	"github.com/0x4445565a/golang-service/pkg/api"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Service struct {
	Redis   *redis.Client
	ReadDB  *sqlx.DB
	WriteDB *sqlx.DB
}

func (s *Service) Init() *echo.Echo {
	e := echo.New()

	// Map paths to handlers
	e.GET("/", s.GetRoot)
	e.POST("/", s.PostRoot)

	return e
}

func (s *Service) GetRoot(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello World")
}

func (s *Service) PostRoot(c echo.Context) error {
	exampleRequest := new(api.ExampleRequest)
	if err := c.Bind(exampleRequest); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, exampleRequest)
}
