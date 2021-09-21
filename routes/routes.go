package routes

import (
	"github.com/ghynmo/MVC_Learning/controllers"
	"github.com/labstack/echo/v4"
)

func NewRoutes() *echo.Echo {
	echoInit := echo.New()

	echoInit.GET("/article", controllers.GetArticleController)
	echoInit.GET("/article/:id", controllers.GetArticlebyIDController)
	echoInit.POST("/article", controllers.SaveArticleController)
	echoInit.PUT("/article/:id", controllers.UpdateArticlebyIDController)
	echoInit.DELETE("/article/:id", controllers.DeleteArticlebyIDController)
	return echoInit
}
