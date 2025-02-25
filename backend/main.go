/*----------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See LICENSE in the project root for license information.
 *---------------------------------------------------------------------------------------*/

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

	// Start the server.
	portNumber := "9000"
	if err := http.ListenAndServe(":"+portNumber, nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
	fmt.Println("Server listening on port ", portNumber)
}
