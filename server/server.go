package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
)

func New() error {
	e := echo.New()

	err := setup(e)
	if err != nil {
		return fmt.Errorf("failed to setup HTTP server: %v", err)
	}

	err = e.Start(":1323")
	if err != nil {
		return fmt.Errorf("failed to start HTTP server: %v", err)
	}

	return nil
}

func setup(e *echo.Echo) error {
	initMiddlewares(e)
	initRoutes(e)

	return nil
}

func initMiddlewares(e *echo.Echo) {
	log.Info("Initializing middlewares...")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
}

func initRoutes(e *echo.Echo) {
	log.Info("Initializing routes...")
	e.GET("/", hello)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!")
}
