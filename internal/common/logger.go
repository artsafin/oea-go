package common

import (
	"go.uber.org/zap"
	"net/http"
)

const (
	RequestIdContextKey = "requestId"
)

func getRequestId(r *http.Request) string {
	id, ok := r.Context().Value(RequestIdContextKey).(string)
	if !ok {
		return "no-request-id"
	}
	return id
}

func NewRequestLogger(r *http.Request, logger *zap.SugaredLogger) *zap.SugaredLogger {
	return logger.With("requestId", getRequestId(r))
}
