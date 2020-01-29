package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"gitlab.com/martha.sutopo/workshop/model"
)

type H map[string]interface{}

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {

		var user model.User
		c.Bind(&user)

		if user.Username == "" || user.Password == "" {
			return c.JSON(http.StatusInternalServerError, H{
				"message": "Error",
				"data":    user,
			})
		} else {
			return c.JSON(http.StatusOK, H{
				"message": "Success",
				"data":    user,
			})

		}

	}
}
