package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/julienschmidt/httprouter"
)

type LogMiddleware struct {
	http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Receive Request")
	middleware.Handler.ServeHTTP(rw, r)
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Fprint(rw, "Middleware")
	})

	middleware := LogMiddleware{router}

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	middleware.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Middleware", string(body))
}
