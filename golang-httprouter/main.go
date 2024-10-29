package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Fprint(rw, "Hello httrouter")
	})

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	server.ListenAndServe()
}
