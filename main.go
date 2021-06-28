package main

import (
	"net/http"

	"exponea.com/controller"
	"github.com/labstack/echo/v4"
)

func main() {
	e := newRouter()
	e.Logger.Fatal(e.Start(":8080"))
}

func newRouter() *echo.Echo {
	e := echo.New()
	g := e.Group("/api")
	g.GET("/all", getAll)
	g.GET("/first", getFirst)
	return e
}

// e.GET("/api/first", getFirst)
func getFirst(c echo.Context) error {
	works, _ := controller.BuildWorks(c)
	return c.JSON(http.StatusOK, works[0])
}

// e.GET("/api/all", getAll)
func getAll(c echo.Context) error {
	works, err := controller.BuildWorks(c)
	if works == nil {
		return err
	}
	return c.JSON(http.StatusOK, works)
}
