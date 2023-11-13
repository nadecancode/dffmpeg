package main

import (
	"controller/handler"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	fmt.Println("Hello World")
	listen()
}

func listen() {
	server := echo.New()

	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	server.GET("/job/create", func(c echo.Context) error {
		return handler.CreateJob(c)
	})

	server.Logger.Fatal(server.Start(":8080"))
}
