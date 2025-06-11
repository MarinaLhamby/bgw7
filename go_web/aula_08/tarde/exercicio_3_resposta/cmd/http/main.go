package main

import (
	"fmt"
	"net/http"

	"github.com/bgw7/exercicio_3_resposta/cmd/http/router"
)

func main() {
	r := router.NewRouter()

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r.MapRoutes())
}
