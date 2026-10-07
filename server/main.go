package main

import (
	"fullstack/handler"
	"fullstack/repository"
	"fullstack/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialise the in‑memory repository.
	repo := repository.NewInMemoryRepo()
	// Wire up the usecase layer.
	uc := usecase.NewBookUsecase(repo)
	// Register the HTTP routes.
	handler.RegisterRoutes(e, uc)

	// Serve the built Vue SPA from the dist folder.
	e.Static("/", "../client/dist")

	e.Logger.Fatal(e.Start(":8080"))
}
