/*----------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See LICENSE in the project root for license information.
 *---------------------------------------------------------------------------------------*/

package main

import (
	"fmt"
	"io"
	"net/http"

	"lets-go/hello"
)

func handle(w http.ResponseWriter, r *http.Request) {
	// I asked GitHub Copilot about the lint errors "return value of io.WriteString is not checked"
	// and it showed me the following code. I'm not sure if this is the best way to handle the error,
	// but I'm going to use it for now.
	if _, err := io.WriteString(w, hello.Hello()); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func main() {
	portNumber := "9000"
	http.HandleFunc("/", handle)
	fmt.Println("Server listening on port ", portNumber)

	if err := http.ListenAndServe(":"+portNumber, nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
