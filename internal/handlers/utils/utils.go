package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ToJSON(w http.ResponseWriter, status int, res *Response) {
	res.Message = http.StatusText(status)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(res)
}

func RespondBadRequest(w http.ResponseWriter, err error) {
	status := http.StatusBadRequest
	ToJSON(w, status, &Response{Message: err.Error()})
}

func RespondInternalServerError(w http.ResponseWriter, err error) {
	//TODO: Log this
	status := http.StatusInternalServerError
	ToJSON(w, status, &Response{Message: err.Error()})
}

func RespondOK(w http.ResponseWriter, data interface{}) {
	status := http.StatusOK
	ToJSON(w, status, &Response{Data: data})
}
