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

type LogMiddleware struct {
	http.Handler
}

func (m *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Receive request - LOG")
	m.Handler.ServeHTTP(w, r)
	fmt.Println("Request already forward to Handler router")
}

func TestMiddleware(t *testing.T) {

	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, params httprouter.Params){
		fmt.Fprint(w, "Hello from root")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	middleware := LogMiddleware{router}

	middleware.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))


	assert.Equal(t, "Hello from root", string(body))
	
}