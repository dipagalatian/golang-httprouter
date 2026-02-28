package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestMethodNotAllowedHandler(t *testing.T) {
	router := httprouter.New()
	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Method Not Allowed for this request")
	})

	router.POST("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello form POST method")
	})

	// this httptest using method GET which is not allowed
	// the available endpoint router only for POST method
	// will directly forward to router.MethodNotAllowed handler
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

	assert.Equal(t, "Method Not Allowed for this request", string(body))

}