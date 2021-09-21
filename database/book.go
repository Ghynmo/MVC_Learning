package database

import (
	"github.com/ghynmo/MVC_Learning/config"
	"github.com/ghynmo/MVC_Learning/models"
)

func GetBooks() (*[]models.Book, error) {

	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return &[]models.Book{}, err
	}

	return &books, nil
}

func GetbyIDBook(id int) (*models.Book, error) {
	var book models.Book

	if err := config.DB.Where("id = ?", id).First(&book).Error; err != nil {
		return &models.Book{}, err
	}

	return &book, nil
}

func StoreBook(book models.Book) (*models.Book, error) {
	if err := config.DB.Save(&book).Error; err != nil {
		return &models.Book{}, err
	}

	return &book, nil
}

func UpdateBook(book models.Book) (*models.Book, error) {
	if err := config.DB.Where("id = ?", book.ID).Updates(&book).Error; err != nil {
		return &models.Book{}, err
	}

	return &book, nil
}

func DeleteBook(id int) (*models.Book, error) {
	var book models.Book

	if err := config.DB.Where("id = ?", id).Delete(&book).Error; err != nil {
		return &models.Book{}, err
	}

	return &book, nil
}
