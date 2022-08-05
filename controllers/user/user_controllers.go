package user

import (
	"net/http"
	"project/online-store/service/user"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	S user.UserService
}

func (c *Controller) FindAllUsers(ctx echo.Context) error {
	users, err := c.S.FindAllUsers()

	if len(users) == 0 {
		return ctx.JSON(http.StatusNotFound, map[string]any{
			"message": "No users found",
		})
	}

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "Successfully found all users",
		"data":    users,
	})
}
