package main

import (
	"fmt"
	"net/http"

	"lets-go/api/post"
)

func get_posts(w http.ResponseWriter, r *http.Request) {
	posts, err := post.GetAllPosts(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	format_response(posts)(w, r)
}

func get_post_by_id(w http.ResponseWriter, r *http.Request) {
	post, err := post.GetPostByID(r.Context(), r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	format_response(post)(w, r)
}

// Utility function to format the response as JSON for every API endpoint.
func format_response(data []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(data); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}
	}
}

func main() {
	// API endpoints go here.
	http.HandleFunc("/api/posts", get_posts)
	http.HandleFunc("/api/post", get_post_by_id)

	// Start the server.
	portNumber := "9000"
	if err := http.ListenAndServe(":"+portNumber, nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
	fmt.Println("Server listening on port ", portNumber)
}
