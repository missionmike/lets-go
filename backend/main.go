/*----------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See LICENSE in the project root for license information.
 *---------------------------------------------------------------------------------------*/

package main

import (
	"fmt"
	"io"
	"net/http"

	"lets-go/api"
	"lets-go/api/post"
)

func handle(w http.ResponseWriter, r *http.Request) {
	// I asked GitHub Copilot about the lint errors "return value of io.WriteString is not checked"
	// and it showed me the following code. I'm not sure if this is the best way to handle the error,
	// but I'm going to use it for now.
	if _, err := io.WriteString(w, api.Hello()); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func main() {
	portNumber := "9000"

	http.HandleFunc("/api/", handle)
	http.HandleFunc("/api/posts", func(w http.ResponseWriter, r *http.Request) {
		posts, err := post.GetAllPosts(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := io.WriteString(w, fmt.Sprintf("%v", posts)); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}
	})

	if err := http.ListenAndServe(":"+portNumber, nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}

	fmt.Println("Server listening on port ", portNumber)
}
