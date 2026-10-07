# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Common Commands

These short command lines are used frequently during development.

| Purpose | Command | Notes |
|---------|---------|------|
| Build the server binary | `go build -o bookstore ./server` | Binary name is `bookstore`.
| Run the server | `go run ./server` | Serves on `:8080` and static SPA from `../client/dist`.
| Build the Vue app | `npm run build` | Works in `client/` folder. Produces `dist/`.
| Start the Vue dev server | `npm run dev` | Runs in `client/` and proxies `/api/*` to `localhost:8080`.
| Run all Go tests | `go test ./...` | Test the repository, usecase, and handler tests.
| Run a single Go test file | `go test ./server/... -run TestXxx` | Replace `TestXxx` with the test name.
| Lint Go code | `golangci-lint run` | Requires `golangci-lint` in PATH.

## High‑Level Architecture

The project is a small full‑stack demo split into a Go backend and a Vue 3 SPA.

### Backend (`/server`)

* **Echo** is the HTTP framework.
* **Repository layer** (`server/repository`) implements an in‑memory data store for books.
* **Usecase layer** (`server/usecase`) contains business logic and is independent from HTTP.
* **Handler layer** (`server/handler`) wires persistence and usecase to Echo routes.
* `main.go` glues everything together, starts the server, and serves static files from `../client/dist`.

This separation keeps routing, business rules, and data access loosely coupled, improving testability.

### Frontend (`/client`)

A standard Vue 3/ Vite setup. The SPA proxies API calls to the Go backend.

## Linting & Code Style

* Go is formatted with `go fmt`.
* No linting rules are currently enforced.

## Existing unit tests

None are shipped yet; the setup above includes test stubs for repository, usecase, and handler. Adding tests helps catch regression when refactoring.

## Ownership & Maintenance Notes

* The Go backend uses a trivial in‑memory store; replace with a real database in production.
* API routes are defined under `/api/books` and follow a standard REST pattern.
* The SPA expects a running Go server on `:8080` during development.

---

Feel free to ask for any additional command aliases or high‑level diagrams if needed.
