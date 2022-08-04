package routers

import (
	"project/online-store/config"
	c "project/online-store/controllers/auth"
	r "project/online-store/repository/auth"
	s "project/online-store/service/auth"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Echo, conf config.Config) {
	db := config.IntiDB(conf)

	repo := r.NewAuthRepository(db)
	service := s.NewAuthService(repo, conf)
	controll := c.Controller{
		S: service,
	}

	g := e.Group("/api")

	g.POST("/register", controll.Register)
	g.POST("/login", controll.Login)
}
