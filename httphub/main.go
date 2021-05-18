package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ElMehdi19/httphub/httphub/router"
)

func main() {
	port := 5000
	mux := router.New()

	log.Printf("Live on http://127.0.0.1:%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
