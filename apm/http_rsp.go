package apm

import (
	"encoding/json"
	"net/http"
)

const (
	jsonContentType = "application/json; charset=utf-8"
)

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Body    any    `json:"body"`
}

type httpStatus struct {
}

var HttpStatus = &httpStatus{}

func (h *httpStatus) OK(w http.ResponseWriter) {
	status := &Status{Code: http.StatusOK, Message: "success"}

	data, _ := json.Marshal(status)
	w.Header().Add("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// OKBody 携带具体业务信息的方法，例如message和body
func (h *httpStatus) OKBody(w http.ResponseWriter, msg string, body any) {
	status := &Status{Code: http.StatusOK, Message: msg, Body: body}

	data, _ := json.Marshal(status)
	w.Header().Add("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// 业务级别错误
func (h *httpStatus) Fail(w http.ResponseWriter, msg string, body any) {
	status := &Status{Code: http.StatusBadRequest, Message: msg, Body: body}

	data, _ := json.Marshal(status)
	w.Header().Add("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusBadRequest)
	w.Write(data)
}

// 服务器级别错误
func (h *httpStatus) Error(w http.ResponseWriter, msg string, body any) {
	status := &Status{Code: http.StatusInternalServerError, Message: msg, Body: body}

	data, _ := json.Marshal(status)
	w.Header().Add("Content-Type", jsonContentType)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(data)
}
