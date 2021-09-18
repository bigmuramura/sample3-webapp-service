package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomePage(c echo.Context) error {
	return c.String(http.StatusOK, "こんにちは、あばしり")
}

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Health Check OK")
}
