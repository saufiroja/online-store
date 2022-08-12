package auth

import (
	"net/http"
	"project/online-store/entity"
	"project/online-store/service/auth"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	S auth.AuthService
}

func (c *Controller) Register(ctx echo.Context) error {
	user := entity.User{
		Id:       uuid.New().String(),
		Name:     ctx.FormValue("name"),
		Email:    ctx.FormValue("email"),
		Password: ctx.FormValue("password"),
		RoleId:   ctx.FormValue("role"),
	}

	err := ctx.Bind(&user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"error":   err.Error(),
			"message": "Invalid request body",
		})
	}

	data, er := c.S.Register(user)
	if er != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"error":   er.Error(),
			"message": "Internal server error",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "User registered successfully",
		"data":    data,
	})
}

func (c *Controller) Login(ctx echo.Context) error {
	user := entity.User{
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

	token, er := c.S.Login(user.Email, user.Password)
	if er != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"error":   er.Error(),
			"message": "Internal server error",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "User logged in successfully",
		"data":    token,
	})
}
