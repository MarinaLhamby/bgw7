package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Info struct {
	Name    string `json:"firstName"`
	Surname string `json:"lastName"`
}

func handleGreeting(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}
	var req Info
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hello %s %s", req.Name, req.Surname)
}

func main() {
	http.HandleFunc("/greetings", handleGreeting)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
