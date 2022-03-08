package apiutils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func writeError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	log.WithField("code", code).WithError(errors.New(msg)).Error("http error")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(&ErrorResponse{
		Code:    code,
		Message: msg,
	})
	if err != nil {
		log.WithError(err).Error("could not write error to response")
	}
}

func WriteSuccessMsg(w http.ResponseWriter, msg string) {
	WriteOKBody(w, &SuccessResponse{
		Message: msg,
	})
}

func WriteOKBody(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		log.WithError(err).Error("could not write error to response")
	}
}

func BadRequest(w http.ResponseWriter, msg string) {
	writeError(w, http.StatusBadRequest, msg)
}

func InternalError(w http.ResponseWriter, msg string) {
	writeError(w, http.StatusInternalServerError, msg)
}

func NotFound(w http.ResponseWriter, msg string) {
	writeError(w, http.StatusNotFound, msg)
}

func Forbidden(w http.ResponseWriter, msg string) {
	writeError(w, http.StatusForbidden, msg)
}

func Unauthorized(w http.ResponseWriter, msg string) {
	writeError(w, http.StatusUnauthorized, msg)
}

// ReadBody reads the body and tells if we should continue handling the request.
func ReadBody(dst interface{}, w http.ResponseWriter, r *http.Request) bool {
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		BadRequest(w, fmt.Sprintf("could not parse body: %v", err))
	}
	return err == nil
}

// ListenAndServe lets the server be closed whenever the context is closed.
func ListenAndServe(ctx context.Context, server *http.Server, startMsg string) error {
	go func() {
		<-ctx.Done()
		server.Close()
	}()
	log.Info(startMsg)
	return server.ListenAndServe()
}
