package hello

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Description for game-api
// @Accept json
// @produce json
// @Success 200 {string} string
func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}
}

// @Description for game-api
// @Accept json
// @produce json
// @Success 200 {json} json
func ApiHelloGet() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"hello": "world"})
	}
}
