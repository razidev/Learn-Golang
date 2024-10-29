package main

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/julienschmidt/httprouter"
)

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {
	router := httprouter.New()
	dir, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(dir))

	request := httptest.NewRequest("GET", "http://localhost:3000/files/hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello HTTPRouter", string(body))
}
