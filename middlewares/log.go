package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LogMiddlewareInit(e *echo.Echo) {
	logger := middleware.LoggerConfig{
		Format: `[${status}] ${method} ${latency_human}` + "\n",
	}
	e.Use(middleware.LoggerWithConfig(logger))
}
