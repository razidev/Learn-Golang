package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello World")
	}

	server := http.Server{
		Addr:    "localhost:4000",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, r.Method)
		fmt.Fprintln(rw, r.RequestURI)
	}

	server := http.Server{
		Addr:    "localhost:4000",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello World")
	})
	mux.HandleFunc("/hi", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hi")
	})
	mux.HandleFunc("/images/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "images")
	})
	mux.HandleFunc("/images/thumbnail/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "thumbnail")
	})

	server := http.Server{
		Addr:    "localhost:4000",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
