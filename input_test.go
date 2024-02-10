package httper

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestInt64(t *testing.T) {

	r := httptest.NewRequest(http.MethodGet, "/testing?test=42", nil)
	w := httptest.NewRecorder()

	test, err := CheckFormInt64(r, w, "test")
	if err != nil || test != 42 || reflect.TypeOf(test).Kind() != reflect.Int64 {
		t.Fatalf("Error read int64 %v", err)
	}

	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != 200 {
		t.Fatalf("Expected status code 200")
	}
}

func TestInt(t *testing.T) {

	r := httptest.NewRequest(http.MethodGet, "/testing?test=42", nil)
	w := httptest.NewRecorder()

	test, err := CheckFormInt(r, w, "test")
	if err != nil || test != 42 || reflect.TypeOf(test).Kind() != reflect.Int {
		t.Fatalf("Error read int64 %v", err)
	}

	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != 200 {
		t.Fatalf("Expected status code 200")
	}
}
