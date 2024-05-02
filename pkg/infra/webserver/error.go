package webserver

import (
	"encoding/json"
	"errors"
	"io"
)

type ResponseError struct {
	Message string `json:"error"`
}

func raiseWebserverError(err error) *ResponseError {
	return &ResponseError{err.Error()}
}

func ReponseErrorFromJSON(input []byte) (*ResponseError, error) {
	var resErr ResponseError
	if err := json.Unmarshal(input, &resErr); err != nil {
		return nil, errors.New("failed to build response from json input")
	}
	return &resErr, nil
}

func ResponseErrorFromJSONStream(r io.Reader) (*ResponseError, error) {
	var resErr ResponseError
	if err := json.NewDecoder(r).Decode(&resErr); err != nil {
		return nil, errors.New("failed to build response from json input stream")
	}
	return &resErr, nil
}
