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

func TestMethodNotAllowed(t *testing.T) {
	router := httprouter.New()

	router.MethodNotAllowed = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Gak boleh")
	})
	router.POST("/", func(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Fprint(rw, "POST")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Gak boleh", string(body))
}
