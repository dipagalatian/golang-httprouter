package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	server := http.Server{
		Addr: ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}