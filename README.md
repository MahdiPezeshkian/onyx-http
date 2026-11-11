# onyx-http

[![Go Version](https://img.shields.io/badge/go-1.27-00ADD8.svg)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

HTTP-specific helpers for writing JSON responses using `onyx-response`
envelopes. Prefer the Go standard library `net/http`.

## Installation

```bash
go get github.com/MahdiPezeshkian/onyx-http@v0.1.0
```

## Import path

```go
import "github.com/MahdiPezeshkian/onyx-http"
```

Package name: `httpresponse`

## Basic usage

```go
func handler(w http.ResponseWriter, r *http.Request) {
	user := &User{ID: 1, Name: "Ada"}
	_ = httpresponse.WriteSuccess(w, http.StatusOK, user, "ok")
}
```

Handlers choose HTTP status codes explicitly:

```go
_ = httpresponse.WriteFailure(w, http.StatusUnauthorized, "invalid refresh token")
```

## Public API overview

| API | Purpose |
|---|---|
| `WriteJSON` | Encode any value as JSON with an explicit status |
| `WriteSuccess` / `WriteFailure` | Write `onyx-response` envelopes |
| `WritePagedSuccess` / `WritePagedFailure` | Write paginated envelopes |

## Design decisions

- Framework-agnostic: works with `net/http` and any framework that exposes
  `http.ResponseWriter`.
- **No centralized business-error → status-code policy.** Callers pass status
  codes.
- Depends on `onyx-response` for envelope shapes only.

## Dependencies

- `github.com/MahdiPezeshkian/onyx-response`
- Go standard library (`net/http`, `encoding/json`)

## Limitations

- Does not integrate with specific routers or middleware stacks.
- Does not translate domain/`onyx-errors` values into HTTP responses
  automatically.

## License

MIT — see [LICENSE](LICENSE).
