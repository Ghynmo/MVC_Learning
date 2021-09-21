package controllers

import (
	"net/http"
	// "strconv"

	"github.com/ghynmo/MVC_Learning/database"
	"github.com/ghynmo/MVC_Learning/models"
	"github.com/labstack/echo/v4"
)

func GetArticleController(echoContext echo.Context) error {

	articles, err := database.GetArticles()
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   articles,
	})
}

func SaveArticleController(echoContext echo.Context) error {
	var articleReq models.Article
	echoContext.Bind(&articleReq)

	result, err := database.StoreArticle(articleReq)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}
	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   result,
	})
}

// func UpdateArticleController(echoContext echo.Context) error {
// 	id, _ := strconv.Atoi(echoContext.Param("id"))

// 	var articleReq models.Article
// 	echoContext.Bind(&articleReq)

// 	result, err := database.Update(articleReq)
// 	if err != nil {
// 		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"status":   "err",
// 			"messages": err,
// 		})
// 	}
// 	return echoContext.JSON(http.StatusOK, map[string]interface{}{
// 		"status": "success",
// 		"data":   result,
// 	})
// }
