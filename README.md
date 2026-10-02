# Go + Gin starter

A minimal [Gin](https://github.com/gin-gonic/gin) HTTP service. It is the starting point for the
Go + Gin briefs in the DayZer0 catalog — clone it, confirm it runs, then build your project's
endpoints on top. It ships with one health route and one passing test; nothing project-specific.

## Stack

- Go 1.22+
- Gin (`github.com/gin-gonic/gin`)

## Run it

```bash
go mod download   # fetch Gin the first time
go run .
```

The server listens on `:8080` (override with `PORT`). Confirm it is up:

```bash
curl http://localhost:8080/health
# {"status":"ok"}
```

## Test it

```bash
go test ./...
```

## Layout

- `main.go` — the entry point. `newRouter()` registers routes; add yours beside `/health`.
- `main_test.go` — the health-route test, a pattern to copy for your own handlers.

## Working the tickets

Each ticket is a GitHub issue on this repo. For a work ticket: branch `issue-<N>-<slug>` off
`main`, commit your work, open a pull request with `Closes #<N>` in its body, and Raj reviews it.
An approved pull request merges and marks the ticket submitted.
