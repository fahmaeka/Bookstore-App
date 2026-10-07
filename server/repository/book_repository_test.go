package repository

import (
    "testing"
)

func TestInMemoryRepoCRUD(t *testing.T) {
    repo := NewInMemoryRepo()
    // Create
    b1 := &Book{Title: "First"}
    created, err := repo.Create(b1)
    if err != nil || created.ID == 0 {
        t.Fatalf("create failed: %v", err)
    }

    // GetByID
    got, err := repo.GetByID(created.ID)
    if err != nil || got == nil || got.Title != "First" {
        t.Fatalf("getbyid failed: %v", err)
    }

    // Update
    updated := &Book{Title: "Updated"}
    _, err = repo.Update(created.ID, updated)
    if err != nil { t.Fatalf("update failed: %v", err) }
    got, _ = repo.GetByID(created.ID)
    if got.Title != "Updated" { t.Fatalf("update not applied: %v", got.Title) }

    // Search
    repo.Create(&Book{Title: "Search me"})
    found, _ := repo.SearchByTitle("search")
    if len(found) == 0 { t.Fatalf("search returned empty") }

    // Delete
    err = repo.Delete(created.ID)
    if err != nil { t.Fatalf("delete failed: %v", err) }
    got, _ = repo.GetByID(created.ID)
    if got != nil { t.Fatalf("expected nil after delete") }
}
