package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rachmankamil/kampus-merdeka-b/lib/database"
	"github.com/rachmankamil/kampus-merdeka-b/models"
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
