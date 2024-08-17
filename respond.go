package httper

import (
	"encoding/json"
	"net/http"
)

type Status struct {
	Status string `json:"status"`
}

func RespondSuccess(w http.ResponseWriter) {
	RespondJson(w, Status{Status: "Success"})
}

func RespondStatus(w http.ResponseWriter, status string) {
	RespondJson(w, Status{Status: status})
}

func RespondJson(w http.ResponseWriter, v any) {
	js, _ := json.MarshalIndent(v, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func RespondJsonGzipped(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Encoding", "gzip")
	gw := gzip.NewWriter(w)
	defer gw.Close()
	json.NewEncoder(gw).Encode(v)
}
