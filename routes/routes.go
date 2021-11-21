package routes

import (
	"github.com/ghynmo/MVC_Learning/controllers"
	"github.com/labstack/echo/v4"
)

func NewRoutes() *echo.Echo {
	echoInit := echo.New()

	echoInit.GET("/books", controllers.GetBookController)
	echoInit.GET("/book/:id", controllers.GetBookbyIDController)
	echoInit.POST("/book", controllers.SaveBookController)
	echoInit.PUT("/book/:id", controllers.UpdateBookbyIDController)
	echoInit.DELETE("/book/:id", controllers.DeleteBookbyIDController)

	return echoInit
}
