package common

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

func WriteHTTPErrAndLog(logger *zap.SugaredLogger, resp http.ResponseWriter, err interface{}, status ...int) {
	st := http.StatusBadRequest
	if len(status) > 0 {
		st = status[0]
	}

	logger.Errorf("http %v, error: %v", st, err)

	resp.WriteHeader(st)
	_, fmterr := fmt.Fprintf(resp, "error: %v", err)
	if fmterr != nil {
		logger.Errorf("fmt error: %v", fmterr)
	}
}