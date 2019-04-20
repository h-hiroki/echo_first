package main

import (
	"github.com/labstack/echo"
	"github.com/h-hiroki/echo_first/handlers"
)

func main() {
	e := echo.New()

	e.File("/", "public/index.html")
	e.GET("/todos", handlers.GetTodos)

	e.Logger.Fatal(e.Start(":1323"))
}
