package httper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func CheckFormInt64(r *http.Request, w http.ResponseWriter, key string) (int64, error) {
	value := r.FormValue(key)
	if value == "" {
		http.Error(w, fmt.Sprintf("please provide a GET/POST value for key '%s'", key), 400)
		return 0, errors.New(fmt.Sprintf("form value '%s' not found", key))
	}
	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("please provide a valid integer for key '%s'", key), 400)
		return 0, err
	}
	return intValue, nil
}

func CheckFormInt(r *http.Request, w http.ResponseWriter, key string) (int, error) {
	value := r.FormValue(key)
	if value == "" {
		http.Error(w, fmt.Sprintf("please provide a GET/POST value for key '%s'", key), 400)
		return 0, errors.New(fmt.Sprintf("form value '%s' not found", key))
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		http.Error(w, fmt.Sprintf("please provide a valid integer for key '%s'", key), 400)
		return 0, err
	}
	return intValue, nil
}

func CheckFormString(r *http.Request, w http.ResponseWriter, key string) (string, error) {
	value := r.FormValue(key)
	if value == "" {
		http.Error(w, fmt.Sprintf("please provide a GET/POST value for key '%s'", key), 400)
		return "", errors.New(fmt.Sprintf("form value '%s' not found", key))
	}
	return value, nil
}

func GetJsonBodyDefault(r *http.Request, w http.ResponseWriter, v any) error {
	maxBytes := int64(1048576) // 1mb
	return GetJsonBody(r, w, v, maxBytes, true)
}

func GetJsonBody(r *http.Request, w http.ResponseWriter, v any, maxBytes int64, allowUnknownFields bool) error {

	// If the Content-Type header is present, check that it has the value
	// application/json. Note that we parse and normalize the header to remove
	// any additional parameters (like charset or boundary information) and normalize
	// it by stripping whitespace and converting to lowercase before we check the
	// value.
	ct := r.Header.Get("Content-Type")
	if ct != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
		if mediaType != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return errors.New(msg)
		}
	}

	// Use http.MaxBytesReader to enforce a maximum read of 1MB from the
	// response body. A request body larger than that will now result in
	// Decode() returning a "http: request body too large" error.
	r.Body = http.MaxBytesReader(w, r.Body, maxBytes)

	// Setup the decoder and call the DisallowUnknownFields() method on it.
	// This will cause Decode() to return a "json: unknown field ..." error
	// if it encounters any extra unexpected fields in the JSON. Strictly
	// speaking, it returns an error for "keys which do not match any
	// non-ignored, exported fields in the destination".
	dec := json.NewDecoder(r.Body)
	if !allowUnknownFields {
		dec.DisallowUnknownFields()
	}

	err := dec.Decode(&v)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		msg := ""

		switch {
		// Catch any syntax errors in the JSON and send an error message
		// which interpolates the location of the problem to make it
		// easier for the client to fix.
		case errors.As(err, &syntaxError):
			msg = fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			http.Error(w, msg, http.StatusBadRequest)

		// In some circumstances Decode() may also return an
		// io.ErrUnexpectedEOF error for syntax errors in the JSON. There
		// is an open issue regarding this at
		// https://github.com/golang/go/issues/25956.
		case errors.Is(err, io.ErrUnexpectedEOF):
			msg = fmt.Sprintf("Request body contains badly-formed JSON")
			http.Error(w, msg, http.StatusBadRequest)

		// Catch any type errors, like trying to assign a string in the
		// JSON request body to a int field in our Person struct. We can
		// interpolate the relevant field name and position into the error
		// message to make it easier for the client to fix.
		case errors.As(err, &unmarshalTypeError):
			msg = fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			http.Error(w, msg, http.StatusBadRequest)

		// Catch the error caused by extra unexpected fields in the request
		// body. We extract the field name from the error message and
		// interpolate it in our custom error message. There is an open
		// issue at https://github.com/golang/go/issues/29035 regarding
		// turning this into a sentinel error.
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg = fmt.Sprintf("Request body contains unknown field %s", fieldName)
			http.Error(w, msg, http.StatusBadRequest)

		// An io.EOF error is returned by Decode() if the request body is
		// empty.
		case errors.Is(err, io.EOF):
			msg = "Request body must not be empty"
			http.Error(w, msg, http.StatusBadRequest)

		// Catch the error caused by the request body being too large. Again
		// there is an open issue regarding turning this into a sentinel
		// error at https://github.com/golang/go/issues/30715.
		case err.Error() == "http: request body too large":
			msg = "Request body must not be larger than 1MB"
			http.Error(w, msg, http.StatusRequestEntityTooLarge)

		// Otherwise default to logging the error and sending a 500 Internal
		// Server Error response.
		default:
			msg = fmt.Sprint(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return errors.New(msg)
	}

	// Call decode again, using a pointer to an empty anonymous struct as
	// the destination. If the request body only contained a single JSON
	// object this will return an io.EOF error. So if we get anything else,
	// we know that there is additional data in the request body.
	err = dec.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		msg := "Request body must only contain a single JSON object"
		http.Error(w, msg, http.StatusBadRequest)
		return errors.New(msg)
	}

	return nil
}
