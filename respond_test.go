package httper

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestSuccess(t *testing.T) {

	w := httptest.NewRecorder()
	RespondSuccess(w)

	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != 200 {
		t.Fatalf("Expected status code 200")
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Expected error to be nil got %v", err)
	}

	correct := Status{Status: "Success"}
	correctJson, _ := json.MarshalIndent(correct, "", "  ")

	if string(data) != string(correctJson) {
		t.Errorf("Expected %s  got %s", correct, data)
	}

}

func TestStatus(t *testing.T) {

	w := httptest.NewRecorder()
	status := "Processing"
	RespondStatus(w, status)

	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != 200 {
		t.Fatalf("Expected status code 200")
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Expected error to be nil got %v", err)
	}

	correct := Status{Status: status}
	correctJson, _ := json.MarshalIndent(correct, "", "  ")

	if string(data) != string(correctJson) {
		t.Errorf("Expected %s  got %s", correctJson, data)
	}

}

func TestJson(t *testing.T) {

	w := httptest.NewRecorder()

	obj := struct {
		A string
		B string
	}{
		A: "Testing A",
		B: "Testing B",
	}

	RespondJson(w, obj)

	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != 200 {
		t.Fatalf("Expected status code 200")
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Expected error to be nil got %v", err)
	}

	correctJson, _ := json.MarshalIndent(obj, "", "  ")

	if string(data) != string(correctJson) {
		t.Errorf("Expected %s  got %s", correctJson, data)
	}

}
