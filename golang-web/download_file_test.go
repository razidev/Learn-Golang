package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func DownloadFile(rw http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")

	if file == "" {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw, "Bad Request")
		return
	}

	rw.Header().Add("Content-Disposition", "attachment; filename=\""+file+"\"")
	http.ServeFile(rw, r, "./resources/"+file)
}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
