# Bookstore App

> A simple full‑stack application that demonstrates how to expose a **Go** REST API and consume it with a **Vue 3** SPA.

## TL;DR

```bash
# 1️⃣ Run the backend
cd server
# go modules are auto‑downloaded, no go mod download needed
go run . 

# 2️⃣ Run the frontend
cd client
npm install    # or pnpm, yarn
npm run dev   # opens http://localhost:5173
```

The two parts are completely decoupled: the Vue app talks to the Go server via `GET /api/books`, `POST /api/books`, etc. Feel free to modify either side and redeploy.

---

## Project structure

```
├─ client                      # Vue 3 SPA
│  ├─ src                     # Vue components & entry point
│  ├─ vite.config.js          # Vite dev server config (runs on :5173)
│  └─ package.json           # NPM dependencies (Vue, Vite, etc.)
├─ server                      # Go  backend
│  ├─ main.go                 # HTTP server entry point, routes, handlers
│  ├─ go.mod                  # Go module definition
│  └─ go.sum
└─ README.md
```

## Tech stack

| Layer | Technology | Reason |
|---|---|---|
| Backend | **Go (1.22+)** | Fast, statically‑compiled, easy to ship Docker images |
| Server framework | **net/http** + **gorilla/mux** (optional) | Classic minimal routing |
| Database | In‑memory (map) | Simplicity for demo; replaceable with PostgreSQL or MySQL |
| Frontend | **Vue 3** | Modern reactive UI, built with **Vite** |
| Packaging | **Vite** | Lightning‑fast dev server & build pipeline |
| IDE | Any – the repo is tiny and follows standard layout |

> ⚠️  The included database is in‑memory; in production replace with a real persistence layer or a cloud backend.

## Getting started

### Prerequisites

- **Go 1.22+** installed: <https://go.dev/dl>
- Node 20+ (or any compatible LTS) and NPM/Yarn/Pnpm: <https://nodejs.org>

### 1️⃣ Run the backend

```bash
cd server
# Download modules and run
GO111MODULE=on go run .
```

The server starts on `http://localhost:8080`. It exposes the following endpoints:

- `GET /api/books` – list all books
- `GET /api/books/:id` – fetch by id
- `POST /api/books` – create a new book (JSON body)
- `PUT /api/books/:id` – update existing
- `DELETE /api/books/:id` – delete

### 2️⃣ Run the frontend

```bash
cd client
npm install   # or pnpm/yarn
npm run dev   # http://localhost:5173
```

The Vue app proxies API calls to the Go server via Vite's dev‑proxy. In production the built SPA static files can be served by the Go server or any static host.

## Building for production

```bash
# 1️⃣ Build the Go binary
cd server
go build -o bookstore

# 2️⃣ Build the Vue assets
cd client
npm run build   # output to dist/

# 3️⃣ (Optional) Bundling assets into the Go binary
# This step copies the built dist into server/static and serves it.
```

After the above, deploy the `bookstore` binary along with the `client/dist` folder (or serve the SPA separately).

## Development hints

- **Hot‑reload**: Running the two servers concurrently (`go run .` and `npm run dev`) gives instant feedback.
- **Cross‑origin**: In dev mode Vite proxies `/api/*` to `http://localhost:8080`. In prod ensure the SPA’s `base` URL matches Go's host.
- **Database**: Replace the in‑memory `map` with a real DB. Update the `main.go` package accordingly.
- **Testing**: The repo currently has no tests. Add Go unit tests in `server/` and Vue tests with `vitest`.

## Contributing

1. Fork the repo.
2. Create a new branch for your feature/bugfix.
3. Commit and push.
4. Open a Pull Request.
5. Follow the linting/formatting rules (go fmt, eslint).
6. Mention any relevant issue.

## License

MIT © 2026 Eka Fahma

---

> 💡 If you run into any bugs or have feature ideas, open an issue or drop a PR. Happy hacking!
