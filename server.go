package main

import (
	"server/controllers"

	"github.com/labstack/echo/v4"
)


func main() {
    e := echo.New()
    e.GET("/", controllers.Index)

    e.Logger.Fatal(e.Start(":3000"))
}