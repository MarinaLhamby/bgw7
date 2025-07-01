package main

import (
	"net/http"

	"github.com/MarinaLhamby/bgw7/desafio_de_fechamento/cmd/http/router"
)

// @title Desafio Go Web Docs
// @version 1.0.0
// @contact.name Marina Rocha Lhamby
// @host localhost:8080
// @BasePath /
func main() {
	r := router.NewRouter()

	if err := http.ListenAndServe(":8080", r.MapRoutes()); err != nil {
		panic(err)
	}
}
