package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	port := 5000
	handler := mux.NewRouter()

	log.Printf("Live on http://127.0.0.1:%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
