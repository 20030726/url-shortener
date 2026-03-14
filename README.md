# URL Shortener Service

A simple REST API built with Go that converts long URLs into short links and redirects users to the original URL.

## Features

* Create short URLs from long links
* Redirect short URLs to the original destination
* Simple REST API
* Lightweight implementation using Go standard library

---

## API Endpoints

### Create Short URL

**POST /shorten**

Request:

```json
{
  "url": "https://example.com"
}
```

Example using curl:

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

---

### Redirect to Original URL

**GET /{short_code}**

Example:

```
http://localhost:8080/abc123
```

The server will redirect the user to the original URL.

---

## Run Locally

Clone the repository:

```bash
git clone https://github.com/20030726/url-shortener.git
cd url-shortener
```

Run the server:

```bash
go run main.go
```

The server will start on:

```
http://localhost:8080
```

---

## Example Workflow

1. Send a POST request to `/shorten`
2. Receive a short URL
3. Open the short URL in a browser
4. The server redirects to the original link

---

## Project Structure

```
url-shortener/
├── LICENSE
├── README.md
├── go.mod
├── go.sum
└── main.go
```

---

## Contributing

Contributions are welcome.

1. Fork the repository
2. Create a new branch
3. Commit your changes
4. Open a pull request

---

## License

This project is licensed under the MIT License.

