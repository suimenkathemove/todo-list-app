package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/health_check", healthCheck)

	e.Logger.Fatal(e.Start(":8080"))
}

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
