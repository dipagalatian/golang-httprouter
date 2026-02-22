package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// Create a new router
	router := httprouter.New()

	// Register routes and httprouter.handle
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello HTTPRouter GET")
	})

	server := http.Server{
		Addr: ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}