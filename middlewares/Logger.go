package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Logger(e *echo.Echo) {
	e.Use(middleware.Logger())
}
