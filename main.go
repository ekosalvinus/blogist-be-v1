package main

import (
	"encoding/json"
	"encoding/base64"
	"log"
	"net/http"
	"strings"
	"time"
)

type Article struct {
	UUID      string    `json:"uuid"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var articles = []Article{
	{UUID: "a1f3c9e2-7b4e-4c91-9d01-2f8d5e7a1001", Title: "Vue 3 Folder Structure for Production Projects", Author: "Daniel Carter", CreatedAt: parseTime("2026-03-02T10:00:00Z"), UpdatedAt: parseTime("2026-03-02T10:00:00Z")},
	{UUID: "b2d4e8f1-6c3a-4f11-8b2e-3c9a7b2f1002", Title: "Composition API vs Options API (Real Use Case)", Author: "Michael Thompson", CreatedAt: parseTime("2026-03-02T10:01:00Z"), UpdatedAt: parseTime("2026-03-02T10:01:00Z")},
	{UUID: "c3a9f7b2-5d2f-4a21-9c3b-4e1d8f3a1003", Title: "How to Connect Vue 3 to a REST API Cleanly", Author: "Andrew Walker", CreatedAt: parseTime("2026-03-02T10:02:00Z"), UpdatedAt: parseTime("2026-03-02T10:02:00Z")},
	{UUID: "d4b1c6e3-4e8a-4b91-a1d4-5f2c9e6b1004", Title: "Setting Up an Axios Service Layer in Vue", Author: "Christopher Hall", CreatedAt: parseTime("2026-03-02T10:03:00Z"), UpdatedAt: parseTime("2026-03-02T10:03:00Z")},
	{UUID: "e5c2d7a4-3f9b-4d11-b2e5-6a3f1c7d1005", Title: "JWT Authentication in Vue 3 (Best Practices)", Author: "Jonathan Reed", CreatedAt: parseTime("2026-03-02T10:04:00Z"), UpdatedAt: parseTime("2026-03-02T10:04:00Z")},
	{UUID: "f6d3e8b5-2a1c-4e21-c3f6-7b4e2d8c1006", Title: "Optimizing Vue App Bundle Size", Author: "Matthew Brooks", CreatedAt: parseTime("2026-03-02T10:05:00Z"), UpdatedAt: parseTime("2026-03-02T10:05:00Z")},
	{UUID: "07e4f9c6-1b2d-4f31-d4a7-8c5f3e9d1007", Title: "Deploying Nuxt 3 to an Ubuntu VPS", Author: "Ryan Mitchell", CreatedAt: parseTime("2026-03-02T10:06:00Z"), UpdatedAt: parseTime("2026-03-02T10:06:00Z")},
	{UUID: "18f5a1d7-9c3e-4a41-e5b8-9d6a4f1e1008", Title: "How to Build a Reusable Component System in Vue", Author: "Kevin Foster", CreatedAt: parseTime("2026-03-02T10:07:00Z"), UpdatedAt: parseTime("2026-03-02T10:07:00Z")},
	{UUID: "29a6b2e8-8d4f-4b51-f6c9-ae7b5c2f1009", Title: "Modern State Management Without Vuex (Pinia)", Author: "Ethan Cooper", CreatedAt: parseTime("2026-03-02T10:08:00Z"), UpdatedAt: parseTime("2026-03-02T10:08:00Z")},
	{UUID: "3ab7c3f9-7e5a-4c61-a7da-bf8c6d3a1010", Title: "Professional Frontend Error Handling Strategy", Author: "Nathan Hughes", CreatedAt: parseTime("2026-03-02T10:09:00Z"), UpdatedAt: parseTime("2026-03-02T10:09:00Z")},
	{UUID: "4bc8d4a1-6f7b-4d71-b8eb-c09d7e4b1011", Title: "Lazy Loading Routes in Vue for Better Performance", Author: "Lucas Bennett", CreatedAt: parseTime("2026-03-02T10:10:00Z"), UpdatedAt: parseTime("2026-03-02T10:10:00Z")},
	{UUID: "5cd9e5b2-5a8c-4e81-c9fc-d1ae8f5c1012", Title: "Migrating Legacy Vue Projects to Vue 3: A Real Experience", Author: "Oliver Grant", CreatedAt: parseTime("2026-03-02T10:11:00Z"), UpdatedAt: parseTime("2026-03-02T10:11:00Z")},
}

func parseTime(s string) time.Time {
	t, _ := time.Parse(time.RFC3339, s)
	return t
}

func jsonResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// Middleware: Basic Auth
func basicAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Basic ") {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			jsonResponse(w, http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
			return
		}

		payload, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(authHeader, "Basic "))
		if err != nil {
			jsonResponse(w, http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
			return
		}

		parts := strings.SplitN(string(payload), ":", 2)
		if len(parts) != 2 || parts[0] != "admin" || parts[1] != "admin123" {
			jsonResponse(w, http.StatusUnauthorized, map[string]string{"message": "Invalid credentials"})
			return
		}

		next(w, r)
	}
}

// Handler: GET /blog (no middleware)
func getBlogHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		jsonResponse(w, http.StatusMethodNotAllowed, map[string]string{"message": "Method not allowed"})
		return
	}
	jsonResponse(w, http.StatusOK, map[string]any{
		"message": "success",
		"data":    articles,
	})
}

// Handler: GET /articles (with basic auth middleware)
func getArticlesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		jsonResponse(w, http.StatusMethodNotAllowed, map[string]string{"message": "Method not allowed"})
		return
	}
	jsonResponse(w, http.StatusOK, map[string]any{
		"message": "success",
		"data":    articles,
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/blog", getBlogHandler)
	mux.HandleFunc("/articles", basicAuthMiddleware(getArticlesHandler))

	println("Server running on http://localhost:8080")
	println("  GET /blog      -> no auth required")
	println("  GET /articles  -> Basic Auth (admin:admin123)")
	log.Println("Server API in Action")
	http.ListenAndServe(":8080", mux)
}