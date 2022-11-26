package book

type Service interface {
	FindAll() ([]Book, error)
	FindById(id int) (Book, error)
	Create(bookInput BookInput) (Book, error)
	Update(bookI BookInput, id int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	return s.repository.FindAll()
}

func (s *service) FindById(id int) (Book, error) {
	return s.repository.FindById(id)
}

func (s *service) Create(bookInput BookInput) (Book, error) {
	price, _ := bookInput.Price.Int64()

	book := Book{
		Title: bookInput.Title,
		Price: int(price),
	}

	return s.repository.Create(book)
}

func (s *service) Update(bookI BookInput, id int) (Book, error) {
	book, _ := s.repository.FindById(id)

	price, _ := bookI.Price.Int64()

	book.Title = bookI.Title
	book.Price = int(price)

	return s.repository.Update(book, id)
}
