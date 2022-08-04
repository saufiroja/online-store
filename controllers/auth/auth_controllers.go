package auth

import (
	"net/http"
	"project/online-store/entity"
	"project/online-store/service/auth"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	S auth.AuthService
}

func (c *Controller) Register(ctx echo.Context) error {
	user := entity.User{
		Name:     ctx.FormValue("name"),
		Email:    ctx.FormValue("email"),
		Password: ctx.FormValue("password"),
	}
	err := ctx.Bind(&user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"error":   err.Error(),
			"message": "Invalid request body",
		})
	}

	er := c.S.Register(user)
	if er != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"error":   er.Error(),
			"message": "Internal server error",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "User registered successfully",
		"data":    user,
	})
}
