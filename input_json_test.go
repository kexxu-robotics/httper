package httper

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBasic(t *testing.T) {
	reqObj := struct {
		A string
		B int
		C int64
	}{
		A: "Test",
		B: 1,
		C: 2,
	}
	body, _ := json.Marshal(reqObj)

	r := httptest.NewRequest(http.MethodPost, "/testing", bytes.NewReader(body))
	w := httptest.NewRecorder()

	obj := struct {
		A string
		B int
		C int64
	}{}

	err := GetJsonBodyDefault(r, w, &obj)
	if err != nil {
		t.Fatalf("Error read json %v", err)
	}

	if obj.A != "Test" || obj.B != 1 || obj.C != 2 {
		t.Fatalf("Values not correctly read from json body %v", obj)
	}

	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != 200 {
		t.Fatalf("Expected status code 200")
	}
}

func TestMalformed(t *testing.T) {
	body := []byte("this is not json as you can see")

	r := httptest.NewRequest(http.MethodPost, "/testing", bytes.NewReader(body))
	w := httptest.NewRecorder()

	obj := struct {
		A string
		B int
		C int64
	}{}

	err := GetJsonBodyDefault(r, w, &obj)
	if err == nil {
		t.Fatalf("Error read json %v", err)
	}

	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != 400 {
		t.Fatalf("Expected status code 400 got %d", res.StatusCode)
	}
}

func TestTooBig(t *testing.T) {
	reqObj := struct {
		A []byte
	}{
		A: make([]byte, 2000),
	}
	body, _ := json.Marshal(reqObj)

	r := httptest.NewRequest(http.MethodPost, "/testing", bytes.NewReader(body))
	w := httptest.NewRecorder()

	obj := struct {
		A []int
	}{}

	err := GetJsonBody(r, w, &obj, 1024, true)
	if err == nil {
		t.Fatalf("Should give an error, the input size is too big")
	}

	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != 413 {
		t.Fatalf("Expected status code 413 got %d", res.StatusCode)
	}
}
