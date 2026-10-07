package handler

import (
	"fullstack/repository"
	"fullstack/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// RegisterRoutes registers the book CRUD endpoints on the provided
// Echo instance using the supplied usecase.
func RegisterRoutes(e *echo.Echo, uc *usecase.BookUsecase) {
	// GET /api/books (with optional search query)
	e.GET("/api/books", func(c echo.Context) error {
		query := c.QueryParam("search")
		books, err := uc.SearchBooks(query)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, books)
	})

	// GET /api/books/:id
	e.GET("/api/books/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
		}
		book, err := uc.GetBook(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if book == nil {
			return echo.NewHTTPError(http.StatusNotFound, "book not found")
		}
		return c.JSON(http.StatusOK, book)
	})

	// POST /api/books
	e.POST("/api/books", func(c echo.Context) error {
		var b repository.Book
		if err := c.Bind(&b); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		created, err := uc.CreateBook(&b)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, created)
	})

	// PUT /api/books/:id
	e.PUT("/api/books/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
		}
		var b repository.Book
		if err := c.Bind(&b); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		updated, err := uc.UpdateBook(id, &b)
		if err != nil {
			if err.Error() == "not found" {
				return echo.NewHTTPError(http.StatusNotFound, "book not found")
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, updated)
	})

	// DELETE /api/books/:id
	e.DELETE("/api/books/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
		}
		err = uc.DeleteBook(id)
		if err != nil {
			if err.Error() == "not found" {
				return echo.NewHTTPError(http.StatusNotFound, "book not found")
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.NoContent(http.StatusNoContent)
	})
}
