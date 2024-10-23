package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateAutoEscape(rw http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(rw, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  "<p>Ini adalah dead bodies<script>alert('Anda di hack')</script></p>",
	})
}

func TemplateAutoEscapeDisable(rw http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(rw, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  template.HTML("<p>Ini adalah dead body<p>"),
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	// TemplateAutoEscape(recorder, request)
	TemplateAutoEscapeDisable(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		// Handler: http.HandlerFunc(TemplateAutoEscape),
		Handler: http.HandlerFunc(TemplateAutoEscapeDisable),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TemplateXSS(rw http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(rw, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  template.HTML(r.URL.Query().Get("body")),
	})
}

func TestTemplateXSS(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?body=<p>alert</p>", nil)
	recorder := httptest.NewRecorder()

	TemplateXSS(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateXSSServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateXSS),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
