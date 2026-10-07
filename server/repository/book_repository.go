// updated to include SearchByTitle method
package repository

import (
    "errors"
    "strings"
    "sync"
)

// Book represents a simple book entity.
// It mirrors the struct used in the original inline implementation.
type Book struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
}

// BookStore is the interface that must be satisfied by any
// persistence layer for books.
type BookStore interface {
    GetAll() ([]Book, error)
    GetByID(id int) (*Book, error)
    Create(book *Book) (*Book, error)
    Update(id int, book *Book) (*Book, error)
    Delete(id int) error
    SearchByTitle(substring string) ([]Book, error)
}

// InMemoryRepo is a very small, thread‑safe in‑memory implementation.
// It is sufficient for the demo and mirrors the state handling that was
// previously embedded in main.go.
// The mutex protects the map and nextID counter.
// All methods return errors that reflect the HTTP status code used by
// the handlers.
//
// The error returned by GetByID when no record exists is implicitly nil.
// The handler will convert nil to 404.

type InMemoryRepo struct {
    sync.Mutex
    books  map[int]Book
    nextID int
}

// NewInMemoryRepo constructs an empty repository with a starting ID of 1.
func NewInMemoryRepo() *InMemoryRepo {
    return &InMemoryRepo{books: map[int]Book{}, nextID: 1}
}

func (r *InMemoryRepo) GetAll() ([]Book, error) {
    r.Lock()
    defer r.Unlock()
    out := make([]Book, 0, len(r.books))
    for _, b := range r.books {
        out = append(out, b)
    }
    return out, nil
}

func (r *InMemoryRepo) GetByID(id int) (*Book, error) {
    r.Lock()
    defer r.Unlock()
    b, ok := r.books[id]
    if !ok {
        return nil, nil
    }
    return &b, nil
}

func (r *InMemoryRepo) Create(book *Book) (*Book, error) {
    r.Lock()
    defer r.Unlock()
    book.ID = r.nextID
    r.nextID++
    r.books[book.ID] = *book
    return book, nil
}

func (r *InMemoryRepo) Update(id int, book *Book) (*Book, error) {
    r.Lock()
    defer r.Unlock()
    if _, ok := r.books[id]; !ok {
        return nil, errors.New("not found")
    }
    book.ID = id
    r.books[id] = *book
    return book, nil
}

func (r *InMemoryRepo) Delete(id int) error {
    r.Lock()
    defer r.Unlock()
    if _, ok := r.books[id]; !ok {
        return errors.New("not found")
    }
    delete(r.books, id)
    return nil
}

func (r *InMemoryRepo) SearchByTitle(substring string) ([]Book, error) {
    r.Lock()
    defer r.Unlock()
    var out []Book
    lower := strings.ToLower(substring)
    for _, b := range r.books {
        if strings.Contains(strings.ToLower(b.Title), lower) {
            out = append(out, b)
        }
    }
    return out, nil
}
