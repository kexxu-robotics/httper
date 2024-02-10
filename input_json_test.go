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
