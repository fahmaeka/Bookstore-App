// updated to add SearchBooks method
package usecase

import (
	"fullstack/repository"
)

// BookUsecase contains methods for CRUD operations.
// It can be unit‑tested without initializing an HTTP server.

type BookUsecase struct {
	repo repository.BookStore
}

// NewBookUsecase returns a new usecase with the provided repository.
func NewBookUsecase(r repository.BookStore) *BookUsecase {
	return &BookUsecase{repo: r}
}

func (u *BookUsecase) ListBooks() ([]repository.Book, error) {
	return u.repo.GetAll()
}

func (u *BookUsecase) GetBook(id int) (*repository.Book, error) {
	return u.repo.GetByID(id)
}

func (u *BookUsecase) CreateBook(b *repository.Book) (*repository.Book, error) {
	return u.repo.Create(b)
}

func (u *BookUsecase) UpdateBook(id int, b *repository.Book) (*repository.Book, error) {
	return u.repo.Update(id, b)
}

func (u *BookUsecase) DeleteBook(id int) error {
	return u.repo.Delete(id)
}

// SearchBooks returns books whose title contains the given query (case‑insensitive).
// If the query is empty, it falls back to ListBooks.
func (u *BookUsecase) SearchBooks(query string) ([]repository.Book, error) {
	if query == "" {
		return u.ListBooks()
	}
	return u.repo.SearchByTitle(query)
}
