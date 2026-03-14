package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
	"time"
)

//go:embed frontend/dist
var frontendFiles embed.FS

var (
	urlStore = make(map[string]string)   // ← fixed: 'ake' → 'make'
	mu       sync.Mutex
)

func main() {
	rand.Seed(time.Now().UnixNano())

	distFS, err := fs.Sub(frontendFiles, "frontend/dist")
	if err != nil {
		log.Fatalf("Failed to load frontend: %v", err)
	}
	staticFS := http.FS(distFS)

	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		routeHandler(staticFS, w, r)
	})
	log.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if _, err := url.ParseRequestURI(req.URL); err != nil {
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return
	}

	short := generateShortURL()

	// Check for collisions (basic safeguard for demo)
	mu.Lock()
	for {
		if _, exists := urlStore[short]; !exists {
			break
		}
		short = generateShortURL()
	}
	urlStore[short] = req.URL
	mu.Unlock()

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	resp := map[string]string{
		"short_url": fmt.Sprintf("%s://%s/%s", scheme, r.Host, short),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func routeHandler(staticFS http.FileSystem, w http.ResponseWriter, r *http.Request) {
	// Root always serves the SPA index page
	if r.URL.Path == "/" {
		http.FileServer(staticFS).ServeHTTP(w, r)
		return
	}

	// Serve any file that exists in the embedded dist directory
	f, err := staticFS.Open(r.URL.Path)
	if err == nil {
		f.Close()
		http.FileServer(staticFS).ServeHTTP(w, r)
		return
	}

	// Otherwise treat the path segment as a short code and redirect
	redirectHandler(w, r)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	short := r.URL.Path[1:]

	mu.Lock()
	longURL, exists := urlStore[short]
	mu.Unlock()

	if !exists {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}

func generateShortURL() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
