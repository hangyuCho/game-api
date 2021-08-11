package main

import (
	"fmt"

	"github.com/hangyuCho/game-api/hello"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echo 인스턴스 생성
	fmt.Println("hoge")
	e := echo.New()

	//라우팅
	e.GET("/hello", hello.MainPage())
	e.GET("/api/hello", hello.ApiHelloGet())

	//서버 기동
	e.Start(":8080")
}
