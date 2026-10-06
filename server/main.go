package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // simple in‑memory store
    type Book struct { ID int `json:"id"`; Title string `json:"title"` }
    books := []Book{{ID:1,Title:"First"}}
    var nextID = 2

    // API routes
    e.GET("/api/books", func(c echo.Context) error { return c.JSON(http.StatusOK, books) })
    e.POST("/api/books", func(c echo.Context) error {
        var b Book
        if err := c.Bind(&b); err != nil { return err }
        b.ID = nextID; nextID++
        books = append(books, b)
        return c.JSON(http.StatusCreated, b)
    })

    // static: serve from ./client/dist
    e.Static("/", "../client/dist")

    e.Logger.Fatal(e.Start(":8080"))
}
