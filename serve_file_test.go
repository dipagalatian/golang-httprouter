package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var resourceEmbed embed.FS

func TestServeFile(t *testing.T) {
	router := httprouter.New()

	dir, _ := fs.Sub(resourceEmbed, "resources")

	// the end of path must be *filepath - read the ServeFiles docs
	router.ServeFiles("/files/*filepath", http.FS(dir))

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/files/hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

	assert.Equal(t, "Hello httprouter", string(body))
}