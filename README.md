# Bookstore App

> A simple full‑stack application that demonstrates how to expose a **Go** REST API and consume it with a **Vue 3** SPA.

## TL;DR

```bash
# 1️⃣ Run the backend
cd server
go run .

# 2️⃣ Run the frontend
cd client
npm run dev   # http://localhost:5173
```

The two parts are completely decoupled: the Vue app talks to the Go server via `/api/books` endpoints.

## Project structure

```
├─ client                      # Vue 3 SPA
│  ├─ src                     # Vue components & entry point
│  ├─ vite.config.js          # Vite dev server config (runs on :5173)
│  └─ package.json           # NPM dependencies (Vue, Vite, etc.)
├─ server                      # Go backend
│  ├─ main.go                 # HTTP server entry point
│  ├─ repository/             # Data access layer
│  │  └─ book_repository.go
│  ├─ usecase/                # Business logic layer
│  │  └─ book_usecase.go
│  └─ handler/                 # HTTP routes layer
│     └─ book_handler.go
└─ README.md
```

## Tech stack

| Layer | Technology | Reason |
|---|---|---|
| Backend | **Go (1.22+)** | Fast, statically-compiled, easy to ship Docker images |
| Server framework | **Echo** | Light‑weight routing & middleware |
| Database | In‑memory map (replaceable with PostgreSQL/MySQL) | Simplicity for demo |
| Frontend | **Vue 3** | Modern reactive UI, built with **Vite** |

## API

The Go server exposes the following endpoints under `/api/books`:

- `GET /api/books` – List all books (supports `?search=` query for title filtering)
- `GET /api/books/:id` – Get a book by ID
- `POST /api/books` – Create a new book (JSON body required)
- `PUT /api/books/:id` – Update an existing book
- `DELETE /api/books/:id` – Delete a book

The `/search=` query is case‑insensitive and matches a substring of the title.

## Running locally

```bash
# start backend
cd server
go run .
# start frontend (in a separate terminal)
cd client
npm run dev
```

The Vue dev server proxies API requests to `localhost:8080`, which is where the Go server runs.

## Build for production

```bash
# 1️⃣ Build Go binary
cd server
go build -o bookstore

# 2️⃣ Build Vue assets
cd client
npm run build   # output to dist/

# 3️⃣ (Optional) Bundle assets into Go binary
# This step copies the built dist into server/static and serves it.
```

Deploy the `bookstore` binary together with the `client/dist` folder or serve the SPA separately.

## Development hints

- **Hot‑reload**: Run `go run .` and `npm run dev` concurrently for instant feedback.
- **Database**: Replace the in‑memory store with a real database by changing `repository.NewInMemoryRepo()`.
- **Testing**: Run `go test ./...` to run unit tests for repository, usecase, and handler.
- **Linting**: Run `golangci-lint run` (requires the tool to be installed).

## Contributing

1. Fork the repo.
2. Create a topic branch.
3. Commit & push.
4. Open a Pull Request.
5. Follow the linting conventions (go fmt, eslint). Add tests when possible.

## License

MIT © 2026 Eka Fahma
