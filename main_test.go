package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSayHello(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello world")
	}

	req := httptest.NewRequest("GET", "/api/hello", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	// Status code test
	if w.Code != 200 {
		t.Error("Status code is not 200")
	}

	// Return value test
	if w.Body.String() == "Hello world " {
		t.Error("Return value is not Hello world ")
	}

}
