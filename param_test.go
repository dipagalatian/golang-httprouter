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

// Params are used to capture dynamic segments of the URL path. They are defined in the route pattern using a colon (:) followed by the parameter name.
// For example, in the route pattern "/products/:id", ":id" is a parameter that can capture any value in that segment of the URL path.
// When a request matches this route, the value of the parameter can be accessed using the Params object passed to the handler function.
// In this example, if a request is made to "/products/1", the parameter "id" will capture the value "1", which can be retrieved using p.ByName("id") in the handler function.
func TestParam(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Println("Params:", p)
		text := "Product ID: " + p.ByName("id")
		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/products/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product ID: 1", string(body))
}