package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Internet-in-a-Box!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on http://127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
