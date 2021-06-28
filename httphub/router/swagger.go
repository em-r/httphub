package router

import (
	"log"
	"net/http"
	"os"
	"path"
)

func SwaggerUI() http.Handler {
	dir, err := os.Getwd()
	if os.IsNotExist(err) {
		log.Println("error occured while trying to server SwaggerUI:", err.Error())
		return http.NotFoundHandler()
	}
	swagger := path.Join(dir, "/swaggerui")
	return http.FileServer(http.Dir(swagger))
}
