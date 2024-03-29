package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type capsule map[string]any

func SendJSON(w http.ResponseWriter, statuscode int, headers http.Header, content capsule) error {
	payload, err := json.Marshal(content)
	if err != nil {
		return err
	}
	payload = append(payload, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statuscode)
	w.Write(payload)
	return nil
}

func ReadJSON(r *http.Request, dst any) error {
	// ? Is limiting payload size really necessary?
	// maxBytes := 1048576
	// r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		var maxBytesError *http.MaxBytesError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("payload contains incorrectly formatted JSON (at character %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("payload contains incorrectly formatted JSON")
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("payload contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("payload contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
		case errors.Is(err, io.EOF):
			return errors.New("payload must not be empty")
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("payload contains an unknown key %s", fieldName)
		case errors.As(err, &maxBytesError):
			return fmt.Errorf("payload must not be larger than %d bytes", maxBytesError.Limit)
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		default:
			return err
		}
	}

	err = decoder.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("payload must contain a single JSON object")
	}

	return nil
}
