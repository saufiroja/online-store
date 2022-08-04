package main

import (
	"project/online-store/config"
	"project/online-store/routers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	conf := config.Config{}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	routers.AuthRoutes(e, conf)

	err := e.Start("127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
}
