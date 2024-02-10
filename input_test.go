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
		t.Fatalf("Error read int %v", err)
	}

	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != 200 {
		t.Fatalf("Expected status code 200")
	}
}

func TestString(t *testing.T) {

	r := httptest.NewRequest(http.MethodGet, "/testing?test=kexxu", nil)
	w := httptest.NewRecorder()

	test, err := CheckFormString(r, w, "test")
	if err != nil || test != "kexxu" || reflect.TypeOf(test).Kind() != reflect.String {
		t.Fatalf("Error read string %v", err)
	}

	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != 200 {
		t.Fatalf("Expected status code 200")
	}
}

func TestInt64Wrong(t *testing.T) {
	// invalid
	r := httptest.NewRequest(http.MethodGet, "/testing?test=wrong", nil)
	w := httptest.NewRecorder()

	_, err := CheckFormInt64(r, w, "test")
	if err == nil {
		t.Fatalf("Error read wrong int64, should give an error")
	}

	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != 400 {
		t.Fatalf("Was a wrong value, should return status 400 but returned %d", res.StatusCode)
	}

	// missing
	r = httptest.NewRequest(http.MethodGet, "/testing", nil)
	w = httptest.NewRecorder()

	_, err = CheckFormInt64(r, w, "test")
	if err == nil {
		t.Fatalf("Error read wrong int64, should give an error")
	}

	res2 := w.Result()
	defer res2.Body.Close()
	if res2.StatusCode != 400 {
		t.Fatalf("Was a wrong value, should return status 400 but returned %d", res2.StatusCode)
	}

}

func TestIntWrong(t *testing.T) {
	// invalid
	r := httptest.NewRequest(http.MethodGet, "/testing?test=wrong", nil)
	w := httptest.NewRecorder()

	_, err := CheckFormInt(r, w, "test")
	if err == nil {
		t.Fatalf("Error read wrong int, should give an error")
	}

	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != 400 {
		t.Fatalf("Was a wrong value, should return status 400 but returned %d", res.StatusCode)
	}

	// missing
	r = httptest.NewRequest(http.MethodGet, "/testing", nil)
	w = httptest.NewRecorder()

	_, err = CheckFormInt(r, w, "test")
	if err == nil {
		t.Fatalf("Error read wrong int, should give an error")
	}

	res2 := w.Result()
	defer res2.Body.Close()
	if res2.StatusCode != 400 {
		t.Fatalf("Was a wrong value, should return status 400 but returned %d", res2.StatusCode)
	}
}

func TestStringWrong(t *testing.T) {
	// missing
	r := httptest.NewRequest(http.MethodGet, "/testing", nil)
	w := httptest.NewRecorder()

	_, err := CheckFormString(r, w, "test")
	if err == nil {
		t.Fatalf("Error read wrong string, should give an error")
	}

	res2 := w.Result()
	defer res2.Body.Close()
	if res2.StatusCode != 400 {
		t.Fatalf("Was a wrong value, should return status 400 but returned %d", res2.StatusCode)
	}
}
