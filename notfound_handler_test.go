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

func TestNotFoundHandler(t *testing.T) {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Not Found, try another path")
	})
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello root")
	})

	// request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/something", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

	// assert.Equal(t, "Hello root", string(body))
	assert.Equal(t, "Not Found, try another path", string(body))
	
}