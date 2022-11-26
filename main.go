package main

import (
	"fmt"
	"golang/book"
	"golang/handler"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection error")
	}

	fmt.Println("Database connection succeed")
	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()
	router.GET("/", bookHandler.GetBooks)
	router.GET("/:id", bookHandler.GetBook)
	router.POST("/", bookHandler.PostBookHandler)
	router.PUT("/:id", bookHandler.UpdateBook)

	router.Run()
}

// main
// handler -> untuk menangkap request saja
// service -> tanggung jawabnya dengan business logic/fitur
// repository -> tanggung jawabnya dengan database
// db
// mysql
