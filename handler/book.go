package handler

import (
	"fmt"
	"golang/book"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func (h *bookHandler) GetBook(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)

	if err != nil {
		log.Fatal("Cannot convert id from string to int")
	}

	book, err := h.bookService.FindById(intID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (h *bookHandler) PostBookHandler(c *gin.Context) {
	//title, price

	var bookInput book.BookInput
	var errorValidate []string

	if err := c.ShouldBindJSON(&bookInput); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errorValidate = append(errorValidate, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorValidate,
		})
		return
	}

	book, err := h.bookService.Create(bookInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (h *bookHandler) UpdateBook(c *gin.Context) {

	id := c.Param("id")
	intID, err := strconv.Atoi(id)

	if err != nil {
		log.Fatal("Cannot convert id from string to int")
	}

	var bookInput book.BookInput
	var errorValidate []string

	if err := c.ShouldBindJSON(&bookInput); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errorValidate = append(errorValidate, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorValidate,
		})
		return
	}

	book, err := h.bookService.Update(bookInput, intID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}
