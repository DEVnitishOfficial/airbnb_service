package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func init() {
	Validator = NewValidator()
}

func NewValidator() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}

func WriteJSONResponse(w http.ResponseWriter, status int, data any) error {
	// set the content type to the application json
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	// encode the data as josn and write it to the response
	return json.NewEncoder(w).Encode(data)
}

func WriteJSONSuccessResponse(w http.ResponseWriter, status int, message string, data any) error {

	// here map is a hash table in go, it's similir to object in js
	// map[string]any ----> key is string and value is of type any, it can be(string, int, bool, struct, etc)
	response := map[string]any{}

	response["status"] = "success"
	response["message"] = message
	response["data"] = data
	return WriteJSONResponse(w, status, response)

}

func WriteJSONErrorResponse(w http.ResponseWriter, status int, message string, err error) error {
	response := map[string]any{}

	response["status"] = "error"
	response["message"] = message
	response["error"] = err.Error()
	return WriteJSONResponse(w, status, response)
}

func ReadJSONBody(r *http.Request, result any) error {
	// Read the raw body into bytes
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("failed to read body: %w", err)
	}

	// Log the raw body as string (what Postman/client sent)
	fmt.Println("Incoming Request Body:", string(bodyBytes))

	// Reset r.Body so it can be read again by the decoder
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// Decode JSON
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(result)
}
