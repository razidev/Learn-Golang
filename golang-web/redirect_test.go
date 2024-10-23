package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirectTo(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello Redirect")
}

func RedirectFrom(rw http.ResponseWriter, r *http.Request) {
	http.Redirect(rw, r, "/redirect-to", http.StatusTemporaryRedirect)
}

func RedirectOut(rw http.ResponseWriter, r *http.Request) {
	http.Redirect(rw, r, "https://www.facebook.com", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-out", RedirectOut)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
