package usecase

import (
	"fullstack/repository"
	"testing"
)

func TestBookUsecaseSearch(t *testing.T) {
	repo := repository.NewInMemoryRepo()
	repo.Create(&repository.Book{Title: "Go book"})
	repo.Create(&repository.Book{Title: "Python book"})
	uc := NewBookUsecase(repo)
	results, err := uc.SearchBooks("go")
	if err != nil {
		t.Fatalf("search error: %v", err)
	}
	if len(results) != 1 || results[0].Title != "Go book" {
		t.Fatalf("unexpected search result: %+v", results)
	}
}
