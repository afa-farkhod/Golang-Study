package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go Http Server!")
}

func main() {
	http.HandleFunc("/", handler)
	port := 8080

	fmt.Printf("Server is starting at port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
