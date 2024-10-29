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

func TestParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id", func(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		text := "Product " + id
		fmt.Fprint(rw, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/products/2", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 2", string(body))
}
