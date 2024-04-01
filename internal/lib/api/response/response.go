package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func OK() Response {
	return Response{
		Status: StatusOK,
	}
}

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}

func WriteJSON(w http.ResponseWriter, status int, value any) (int, error) {
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(true)
	encoder.Encode(value)

	w.Header().Set("Content-Type", "applications/json")
	w.WriteHeader(status)

	return w.Write(buf.Bytes())
}

func WriteOKJSON(w http.ResponseWriter, value any)(int, error) {
	return WriteJSON(w, http.StatusOK, value)
}

func WriteErrorJSON(w http.ResponseWriter, status int, errorMessage string) (int, error) {
	return WriteJSON(w, status, Error(errorMessage))
}

func WriteValidationErrorsJSON(w http.ResponseWriter, status int, errs validator.ValidationErrors) (int, error) {
	return WriteJSON(w, status, ValidationErrsToResponse(errs))
}

func ValidationErrsToResponse(errs validator.ValidationErrors) Response {
	var errMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("Поле %s обязательно для заполнения", err.Field()))
		case "email":
			errMsgs = append(errMsgs, fmt.Sprintf("Поле %s должно быть валидным email'ом", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("Неверно заполнено поле %s", err.Field()))
		}
	}

	return Response{
		Status: StatusError,
		Error:  strings.Join(errMsgs, ", "),
	}
}
