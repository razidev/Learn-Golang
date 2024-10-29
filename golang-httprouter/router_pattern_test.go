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

func TestRouterPatternNamedParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id/items/:itemId", func(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		itemId := params.ByName("itemId")
		text := "Product " + id + " Item " + itemId
		fmt.Fprint(rw, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/products/2/items/3", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 2 Item 3", string(body))
}

func TestRouterPatternCatchAllParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
		image := params.ByName("image")
		text := "Image: " + image
		fmt.Fprint(rw, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/images/small/profile/razi.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Image: /small/profile/razi.png", string(body))
}
