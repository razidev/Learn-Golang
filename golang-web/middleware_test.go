package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Before execute handler")
	middleware.Handler.ServeHTTP(rw, r)
	fmt.Println("After execute handler")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi error")
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(rw, "Error: %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(rw, r)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("handler executed")
		fmt.Fprint(rw, "Hello Middleware")
	})
	mux.HandleFunc("/foo", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("foo executed")
		fmt.Fprint(rw, "Hello Foo")
	})
	mux.HandleFunc("/panic", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("panic executed")
		panic("upps")
	})

	LogMiddleware := &LogMiddleware{
		Handler: mux,
	}

	errorHandler := &ErrorHandler{
		Handler: LogMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
