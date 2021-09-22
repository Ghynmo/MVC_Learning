package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func logMiddlewareInit(e *echo.Echo) {
	logger := middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}
		${path} ${latency_human}` + "\n",
	}
	e.Use(middleware.LoggerWithConfig(logger))
}