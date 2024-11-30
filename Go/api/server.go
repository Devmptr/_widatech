package api

import (
	"fmt"
	"net/http"
	"widatech_interview/golang/config"

	"github.com/labstack/echo/v4"
)

func Create() *echo.Echo {
	return echo.New()
}

func Register(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}

func Start(e *echo.Echo, serverConfig *config.ServerConfig) {
	e.Logger.Fatal(
		e.Start(fmt.Sprintf("%s:%d", serverConfig.Address, serverConfig.Port)),
	)
}
