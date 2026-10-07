package handler

import (
	"bytes"
	"encoding/json"
	"fullstack/repository"
	"fullstack/usecase"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestBookCRUDHandler(t *testing.T) {
	repo := repository.NewInMemoryRepo()
	uc := usecase.NewBookUsecase(repo)
	e := echo.New()
	RegisterRoutes(e, uc)

	// Create book
	b := map[string]string{"Title": "Test Book"}
	body, _ := json.Marshal(b)
	req := httptest.NewRequest("POST", "/api/books", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	if rec.Code != http.StatusCreated {
		t.Fatalf("unexpected status: %d", rec.Code)
	}
	var created struct {
		ID    int
		Title string
	}
	json.Unmarshal(rec.Body.Bytes(), &created)

	// Get by ID
	req = httptest.NewRequest("GET", "/api/books/"+strconv.Itoa(created.ID), nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("get status: %d", rec.Code)
	}

	// Update book
	updated := map[string]string{"Title": "Updated"}
	body, _ = json.Marshal(updated)
	req = httptest.NewRequest("PUT", "/api/books/"+strconv.Itoa(created.ID), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("update status: %d", rec.Code)
	}

	// Delete book
	req = httptest.NewRequest("DELETE", "/api/books/"+strconv.Itoa(created.ID), nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	if rec.Code != http.StatusNoContent {
		t.Fatalf("delete status: %d", rec.Code)
	}

	// Verify deletion
	req = httptest.NewRequest("GET", "/api/books/"+strconv.Itoa(created.ID), nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected 404 after delete, got %d", rec.Code)
	}
}
