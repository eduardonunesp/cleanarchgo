package webserver

import (
	"encoding/json"
	"errors"
	"io"
)

type responseError struct {
	Message string `json:"error"`
}

func raiseWebserverError(err error) *responseError {
	return &responseError{err.Error()}
}

func ReponseErrorFromJSON(input []byte) (*responseError, error) {
	var resErr responseError
	if err := json.Unmarshal(input, &resErr); err != nil {
		return nil, errors.New("failed to build response from json input")
	}
	return &resErr, nil
}

func ResponseErrorFromJSONStream(r io.Reader) (*responseError, error) {
	var resErr responseError
	if err := json.NewDecoder(r).Decode(&resErr); err != nil {
		return nil, errors.New("failed to build response from json input stream")
	}
	return &resErr, nil
}
