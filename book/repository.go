package book

import "gorm.io/gorm"

//contains methods
type Repository interface {
	FindAll() ([]Book, error)
	FindById(id int) (Book, error)
	Create(book Book) (Book, error)
	Update(bookDB Book, id int) (Book, error)
}

//struct from interface
type repository struct {
	db *gorm.DB
}

//inialize struct
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Book, error) {
	var books []Book

	if err := r.db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (r *repository) FindById(id int) (Book, error) {
	var book Book
	err := r.db.Find(&book, id).Error
	return book, err
}

func (r *repository) Create(book Book) (Book, error) {
	err := r.db.Create(&book).Error
	return book, err
}

func (r *repository) Update(book Book, id int) (Book, error) {
	err := r.db.Save(&book)
	return book, err.Error
}
