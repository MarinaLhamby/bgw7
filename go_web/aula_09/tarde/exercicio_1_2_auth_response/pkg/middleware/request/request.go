package request

import (
	"fmt"
	"net/http"
	"time"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Verbo: %s\nData e hora: %s\nUrl consulta: %s\nTamanho consulta: %d", r.Method, time.Now().String(), r.URL, r.ContentLength)
		next.ServeHTTP(w, r)
	})
}
