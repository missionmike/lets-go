package main

import (
	"fmt"
	"net/http"

	"lets-go/api/post"
)

// HandlerFunc is a custom type that returns an error
type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

// Middleware to handle common error responses and HTTP method validation
func withErrorHandling(h HandlerFunc, allowedMethod string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if allowedMethod != "" && r.Method != allowedMethod {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		err := h(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func writeJSON(w http.ResponseWriter, data []byte) error {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write response: %w", err)
	}
	return nil
}

func getAllPosts(w http.ResponseWriter, r *http.Request) error {
	posts, err := post.GetAllPosts(r.Context())
	if err != nil {
		return err
	}
	return writeJSON(w, posts)
}

func getPostByID(w http.ResponseWriter, r *http.Request) error {
	post, err := post.GetPostByID(r.Context(), r.URL.Query().Get("id"))
	if err != nil {
		return err
	}
	return writeJSON(w, post)
}

func createPost(w http.ResponseWriter, r *http.Request) error {
	title := r.FormValue("title")
	desc := r.FormValue("desc")
	id, err := post.CreatePost(r.Context(), title, desc)
	if err != nil {
		return err
	}
	return writeJSON(w, []byte(id))
}

func updatePost(w http.ResponseWriter, r *http.Request) error {
	id := r.FormValue("id")
	title := r.FormValue("title")
	desc := r.FormValue("desc")
	err := post.UpdatePost(r.Context(), id, title, desc)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// API endpoints
	http.HandleFunc("/api/posts", withErrorHandling(getAllPosts, http.MethodGet))
	http.HandleFunc("/api/post", withErrorHandling(getPostByID, http.MethodGet))
	http.HandleFunc("/api/post/create", withErrorHandling(createPost, http.MethodPost))
	http.HandleFunc("/api/post/update", withErrorHandling(updatePost, http.MethodPut))

	// Start the server
	portNumber := "9000"
	fmt.Println("Server starting on port", portNumber)
	if err := http.ListenAndServe(":"+portNumber, nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
