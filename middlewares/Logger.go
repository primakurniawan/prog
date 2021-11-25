package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Logger(e *echo.Echo) {
	e.Use(middleware.Logger())
}
