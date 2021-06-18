package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ElMehdi19/httphub/httphub/helpers"
	"github.com/ElMehdi19/httphub/httphub/router"
)

func main() {
	mux := router.New()

	log.Printf("Live on http://127.0.0.1:%d", helpers.PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", helpers.PORT), mux))
}
