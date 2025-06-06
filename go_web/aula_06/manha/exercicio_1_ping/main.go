package main

import (
	"fmt"
	"net/http"
)

func handlePing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
func main() {
	http.HandleFunc("/ping", handlePing)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
