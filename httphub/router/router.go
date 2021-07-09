package router

import (
	"net/http"

	"github.com/ElMehdi19/httphub/httphub/helpers"
	"github.com/ElMehdi19/httphub/httphub/middlewares"
	"github.com/gorilla/mux"
)

var topLevelGetpaths []string

func New() http.Handler {
	handler := mux.NewRouter()
	handler.Use(middlewares.Recover, middlewares.Logger, middlewares.ContentType, middlewares.CORS)

	if helpers.IsDevMode() {
		handler.HandleFunc("/debug", SourceCodeHandler)
	}

	// HTTP Methods handlers
	handler.HandleFunc("/get", ViewGet).Methods("GET")
	handler.HandleFunc("/post", ViewPost).Methods("POST")
	handler.HandleFunc("/put", ViewPut).Methods("PUT", "OPTIONS")
	handler.HandleFunc("/patch", ViewPatch).Methods("PATCH", "OPTIONS")
	handler.HandleFunc("/delete", ViewDelete).Methods("DELETE", "OPTIONS")
	handler.HandleFunc("/any", ViewAny)

	// Request inspection handlers
	handler.HandleFunc("/request", ViewRequest).Methods("GET")
	handler.HandleFunc("/ip", ViewIP).Methods("GET")
	handler.HandleFunc("/user-agent", ViewUserAgent).Methods("GET")
	handler.HandleFunc("/headers", ViewHeaders).Methods("GET")

	// Status codes handlers
	handler.HandleFunc("/status/{code}", ViewStatusCodes)

	// Auth handlers
	handler.HandleFunc("/auth/basic/{user}/{passwd}", ViewBasicAuth).Methods("GET")
	handler.HandleFunc("/auth/basic-hidden/{user}/{passwd}", ViewBasicAuthHidden).Methods("GET")
	handler.HandleFunc("/auth/bearer", ViewBearerAuth).Methods("GET")

	// Response inspection handlers
	handler.HandleFunc("/response-headers", ViewResponseHeader).Methods("GET")
	handler.HandleFunc("/cache", ViewCache).Methods("GET")
	handler.HandleFunc("/cache/{value}", ViewCacheControl).Methods("GET")

	// Response formats handlers
	handler.HandleFunc("/json", ViewJSONResponse).Methods("GET")
	handler.HandleFunc("/xml", ViewXMLResponse).Methods("GET")
	handler.HandleFunc("/html", ViewHTMLResponse).Methods("GET")
	handler.HandleFunc("/txt", ViewTXTResponse).Methods("GET")

	// Cookies handlers
	handler.HandleFunc("/cookies", ViewCookies).Methods("GET")
	handler.HandleFunc("/cookies/set", ViewSetCookies).Methods("GET")
	handler.HandleFunc("/cookies/set/{name}/{value}", ViewSetCookie).Methods("GET")
	handler.HandleFunc("/cookies/delete", ViewDeleteCookies).Methods("GET")

	// Redirection handlers
	handler.HandleFunc("/redirect/{to}", ViewRedirect).Methods("GET")
	handler.HandleFunc("/redirect", ViewRedirectRandom).Methods("GET")

	handler.Walk(helpers.WalkRouterGET(&topLevelGetpaths, true, "redirect"))

	// Swagger UI
	handler.Handle("/swagger/{file}", http.StripPrefix("/swagger/", SwaggerUI()))
	handler.Handle("/", http.StripPrefix("/", SwaggerUI()))

	return handler
}
