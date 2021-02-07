package common

import (
	"fmt"
	"log"
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

type RequestLogger struct {
	RequestId string
}

func NewRequestLogger(r *http.Request) RequestLogger {
	return RequestLogger{RequestId: getRequestId(r)}
}

func (l *RequestLogger) prefix() string {
	return fmt.Sprintf("[%s] ", l.RequestId)
}

func (l *RequestLogger) Println(v ...interface{}) {
	log.Print(l.prefix(), fmt.Sprintln(v...))
}

func (l *RequestLogger) Printf(format string, v ...interface{}) {
	log.Print(l.prefix(), fmt.Sprintf(format, v...))
}
