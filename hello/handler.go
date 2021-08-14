package hello

import (
	"net/http"

	_ "github.com/hangyuCho/game-api/hello/docs"
	"github.com/labstack/echo/v4"
)

// @description game API 상세
// @version 1.0
// @accept application/x-json-stream
// @param none query string false "필요 없습니다."
// @Success 200 {string} string "ok"
// @router /hello [get]
func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}
}

// @description game API 상세
// @version 1.0
// @accept application/x-json-stream
// @param none query string false "필요 없습니다."
// @Success 200 {string} string "ok"
// @router /api/hello [get]
func ApiHelloGet() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"hello": "world"})
	}
}
