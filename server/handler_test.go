package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGolServerGet(t *testing.T) {
	testFile := tempTest("get")
	defer os.Remove(testFile)
	initDb(testFile)
	handler := http.HandlerFunc(NewGolServerHandler(testFile))

	req, err := http.NewRequest("GET", "/k1/test/test", nil)
	if err != nil {
		t.Errorf("create request failed %s", err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status code: got %v", status)
	}

	// Check the response body is what we expect.
	expected := `http://test/v1/test/test`
	if rr.Header().Get("Location") != expected {
		t.Errorf("handler returned unexpected location header: got %v", rr.Header().Get("Location"))
	}
}

func TestGolServerGetNotFound(t *testing.T) {
	testFile := tempTest("get_not_found")
	defer os.Remove(testFile)
	initDb(testFile)
	handler := http.HandlerFunc(NewGolServerHandler(testFile))

	req, err := http.NewRequest("GET", "/k_not_found/test/test", nil)
	if err != nil {
		t.Errorf("create request failed %s", err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v", status)
	}
}

func TestGolServerDump(t *testing.T) {
	testFile := tempTest("dump")
	defer os.Remove(testFile)
	initDb(testFile)
	handler := http.HandlerFunc(NewGolServerHandler(testFile))

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("create request failed %s", err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v", status)
	}

	expected := `{
	"k1": "http://test/v1",
	"k2": "http://v2"
}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body %v", rr.Body.String())
	}

}
