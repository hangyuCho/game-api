// +build tools

package main

import (
	"fmt"

	_ "github.com/hangyuCho/game-api/docs"

	"github.com/hangyuCho/game-api/hello"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title GAME API
// @version 0.0.1
// @license.name Apache 2.0
// @description This is a game api server
// @contact.name Request permission of game API
// @contact.url http://github.com/hangyuCho
// @contact.email chg3376@gmail.com
// @host localhost
// @BasePath /api
func main() {
	// Echo 인스턴스 생성
	fmt.Println("hoge")
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//라우팅
	e.GET("/hello", hello.MainPage())
	e.GET("/api/hello", hello.ApiHelloGet())
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	//서버 기동
	//e.Start(":8080")
	e.Logger.Fatal(e.Start(":30000"))
}
