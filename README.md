# URL Shortener Service

A simple URL shortener with a **TypeScript / Vite web UI** and a Go REST API backend. Open `http://localhost:8080` in your browser to shorten URLs without using the terminal.

## Features

* 🔗 Shorten long URLs via the web UI or REST API
* ↩️ Redirect short URLs to the original destination
* 📋 One-click copy of the shortened link
* 🕓 In-page history of recently shortened links
* Lightweight Go backend with embedded frontend (single binary)

---

## Quick Start

### Prerequisites

| Tool | Minimum version |
|------|----------------|
| Go   | 1.21           |
| Node.js | 18          |
| npm  | 9              |

### Build & Run

```bash
# 1. Clone the repo
git clone https://github.com/20030726/url-shortener.git
cd url-shortener

# 2. Install frontend dependencies, compile TypeScript, then build the Go binary
make build

# 3. Start the server
./url-shortener
```

The server (and web UI) will be available at **http://localhost:8080**.

> **One-step alternative:**
> ```bash
> make run
> ```

---

## Web UI

Open **http://localhost:8080** in your browser:

1. Paste any long URL into the input box.
2. Click **Shorten**.
3. Copy the generated short link with the **Copy** button.
4. The last 10 shortened links appear in the *Recent links* list below.

---

## REST API

You can still use the API directly if you prefer.

### Create Short URL

**POST /shorten**

```bash
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url":"https://golang.org"}'
```

Response:

```json
{
  "short_url": "http://localhost:8080/abc123"
}
```

### Redirect to Original URL

**GET /{short_code}**

```
http://localhost:8080/abc123
```

The server redirects the browser to the original URL.

---

## Development

### Frontend (TypeScript + Vite)

```bash
cd frontend
npm install
npm run dev        # dev server on http://localhost:5173 (proxies /shorten to :8080)
npm run build      # compile to frontend/dist/
```

### Backend (Go)

```bash
go run main.go     # hot-reload via `go run`
go test ./...      # run tests
```

### All-in-one

```bash
make build   # build frontend then Go binary
make run     # build + start server
make test    # run Go tests
make clean   # remove build artefacts
```

---

## Project Structure

```
url-shortener/
├── frontend/            # TypeScript + Vite web UI
│   ├── src/
│   │   ├── api.ts       # API client (calls /shorten)
│   │   ├── main.ts      # UI logic
│   │   └── style.css    # Styles
│   ├── dist/            # Compiled assets (embedded in Go binary)
│   ├── index.html
│   └── package.json
├── main.go              # Go HTTP server (embeds frontend/dist)
├── main_test.go
├── Makefile
├── go.mod
└── README.md
```

---

## Contributing

1. Fork the repository
2. Create a new branch
3. Commit your changes
4. Open a pull request

---

## License

This project is licensed under the MIT License.
