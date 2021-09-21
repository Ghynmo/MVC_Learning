package controllers

import (
	"net/http"
	"strconv"

	"github.com/ghynmo/MVC_Learning/database"
	"github.com/ghynmo/MVC_Learning/models"
	"github.com/labstack/echo/v4"
)

func GetBookController(echoContext echo.Context) error {

	books, err := database.GetBooks()
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   books,
	})
}

func GetBookbyIDController(echoContext echo.Context) error {

	id, _ := strconv.Atoi(echoContext.Param("id"))

	books, err := database.GetbyIDBook(id)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   books,
	})
}

func SaveBookController(echoContext echo.Context) error {
	var bookReq models.Book
	echoContext.Bind(&bookReq)

	result, err := database.StoreBook(bookReq)
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

func UpdateBookbyIDController(echoContext echo.Context) error {

	id, _ := strconv.Atoi(echoContext.Param("id"))

	var bookReq models.Book
	echoContext.Bind(&bookReq)

	result, err := database.UpdateBook(id, bookReq)
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

func DeleteBookbyIDController(echoContext echo.Context) error {

	id, _ := strconv.Atoi(echoContext.Param("id"))

	books, err := database.DeleteBook(id)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   books,
	})
}
