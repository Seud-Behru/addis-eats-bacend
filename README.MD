### Go backend 
# Addis Eats API

A food-delivery backend for Addis Ababa, written in Go using only the standard library (`net/http`, `encoding/json`). Customers in areas like Bole, Kazanchis, Piassa and Megenagna order dishes from restaurants such as Yod Abyssinia and Kategna. All prices are in **santim** (integers, never decimals).

This is the capstone project for **CodeOps Module 4 (Golang)** at IBT College Canada. It grows day by day and is assessed on Day 75.

> **Status:** Day 55. The first four endpoints are live. Data is stored **in memory** and is **lost when the server restarts**. This is expected until PostgreSQL arrives in Module 5. It is not a bug.

---

## Requirements

- Go **1.22 or newer** (method and wildcard routing is built into the standard library from 1.22)
- `curl` and optionally `jq` for trying the endpoints

## Run it

```bash
git clone https://github.com/<your-username>/addis-eats-api.git
cd addis-eats-api
go run -race ./cmd/api
```

The server listens on `:8080` by default. The `-race` flag turns on Go's race detector, which catches bugs that only appear when several requests arrive at once.

### Configuration

| Variable   | Default | Meaning                                            |
|------------|---------|----------------------------------------------------|
| `API_ADDR` | `:8080` | Address to listen on, for example `127.0.0.1:9000` |

```bash
API_ADDR=:9000 go run -race ./cmd/api
```

The server sets read (5s), write (10s) and idle (1 min) timeouts so a slow client cannot hold a connection open forever.

---

## Endpoints

| Method | Path           | Description                           | Success |
|--------|----------------|---------------------------------------|---------|
| GET    | `/health`      | Liveness check                        | `200`   |
| GET    | `/menu`        | List every dish                       | `200`   |
| GET    | `/menu/{id}`   | One dish, or a 404                    | `200`   |
| POST   | `/orders`      | Place an order; the server computes the total | `201`   |

Any other method on these paths returns `405 Method Not Allowed`, produced automatically by the router.

### Coming next

`GET /orders/{id}` (Day 56), `PATCH /orders/{id}` (Day 57), `POST /auth/login` (Day 61), `POST /menu` for restaurants (Day 62).

---

## Response format

Every response, including errors, uses one envelope.

**One resource**
```json
{ "data": { "id": "d2", "name": "Shiro", "price_santim": 12000 } }
```

**A collection** (an empty list is `200` with `[]`, never `null` and never `404`)
```json
{ "data": [ { "id": "d1", "name": "..." }, { "id": "d2", "name": "..." } ] }
```

**An error**
```json
{ "error": { "message": "dish not found" } }
```

Internal error details are written to the server log and never returned to the client.

---

## Endpoint details

### `GET /health`

```bash
curl -s localhost:8080/health
```

Returns `200` when the server is running.

### `GET /menu`

```bash
curl -s localhost:8080/menu | jq
```

Returns `200` with every dish in `data`.

### `GET /menu/{id}`

```bash
curl -s localhost:8080/menu/d2 | jq
```

| Status | When                               |
|--------|------------------------------------|
| `200`  | The dish exists                    |
| `404`  | No dish has that id                |
| `500`  | The store failed (logged in full)  |

### `POST /orders`

Request body:

```json
{
  "area": "Bole",
  "phone": "0911223344",
  "dish_ids": ["d1", "d2", "d4"]
}
```

```bash
curl -s -i -X POST localhost:8080/orders \
  -H "Content-Type: application/json" \
  -d '{"area":"Bole","phone":"0911223344","dish_ids":["d1","d2"]}'
```

Success response:

```
HTTP/1.1 201 Created
Location: /orders/o-7f3a
```
```json
{
  "data": {
    "id": "o-7f3a",
    "area": "Bole",
    "status": "placed",
    "total_santim": 71925
  }
}
```

The client sends only what it **wants** (area, phone, dish ids). The server decides the id, status, prices and total. Prices are always read from the menu, never trusted from the client.

| Status | When                                                                 |
|--------|----------------------------------------------------------------------|
| `201`  | Order created; `Location` header points at the new order             |
| `400`  | Broken JSON, unknown field, body over 1 MB, empty `dish_ids`, unknown dish id, or an area we do not deliver to |
| `500`  | The store failed (logged in full)                                    |

> Some of these `400` cases become `422 Unprocessable Entity` on Day 58, when proper validation arrives.

---

## Try everything

`requests.http` in the project root contains one call per endpoint, including the failure cases, with the status each one returns. It works with the VS Code REST Client extension or JetBrains HTTP Client.

Quick verification checklist:

| Request                      | Expect |
|------------------------------|--------|
| `GET /menu`                  | 200    |
| `GET /menu/d2`               | 200    |
| `GET /menu/zzz`              | 404    |
| `POST /menu`                 | 405    |
| `POST /orders` (valid)       | 201    |
| `POST /orders` (bad JSON)    | 400    |
| `POST /orders` (unknown dish)| 400    |
| `POST /orders` (area "Ayat") | 400    |

---

## Project layout

```
addis-eats-api/
  cmd/api/              everything that knows about HTTP
    main.go             builds the application and starts the server
    routes.go           every route, on one screen
    handlers.go         handlers (methods on *application) and error helpers
    orders.go           createOrder and its request struct
    helpers.go          writeJSON, writeError
  internal/             the domain: types, rules, errors (never imports net/http)
    menu/               Dish, DishReader, ErrNotFound, in-memory store
    order/              Order, Store, totals, in-memory store, id generation
    delivery/           delivery fee by Addis area
  requests.http         one call per endpoint
```

### Design decisions

- **Standard library only.** Go 1.22's `ServeMux` handles methods and path wildcards.
- **`application` struct for dependencies.** Handlers are methods on it. There are no package-level variables.
- **Small interfaces** (`menu.DishReader`, `order.Store`). When PostgreSQL arrives, only the constructor call in `main` changes.
- **The domain owns the money.** Totals and fees are computed in `internal/`, never in handlers.
- **Money is an integer** in santim, end to end.
- **Random order ids** from `crypto/rand`, so ids cannot be guessed or counted through.
- **Concurrency-safe stores.** Each in-memory store is protected by a `sync.RWMutex`.
- **One envelope** for every response, enforced by the helpers.
- **`/menu` is singular** because it names the restaurant's single menu document. It may be renamed to `/dishes` on Day 57.

---

## Code quality checks

All of these should print nothing:

```bash
gofmt -l .
go vet ./...
grep -rn "net/http" internal/
```

And this should start cleanly and survive two simultaneous `POST /orders`:

```bash
go run -race ./cmd/api
```

---

## Git workflow

- One branch per day, for example `feat/d55-first-endpoints`
- Commit messages in the style `feat: add POST /orders`
- Branches are merged to `main` through a pull request

---

## Roadmap

| Day   | Adds                                                  |
|-------|-------------------------------------------------------|
| 56    | `GET /orders/{id}`, routing and query strings         |
| 57    | Order status updates, pagination, REST design         |
| 58    | Validation and `422` responses                        |
| 59    | Panic recovery and request-logging middleware         |
| 60    | Structured logging                                    |
| 61    | Authentication (access and refresh tokens)            |
| 62    | Roles; restaurants can create dishes                  |
| 63    | Security review                                       |
| 68    | Concurrency and graceful shutdown                     |
| 75    | Capstone assessment                                   |
| Mod 5 | PostgreSQL replaces the in-memory stores              |

---

## Author

Seud, CodeOps Full Stack Software Development, IBT College Canada.
