package common

import (
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func JoinErrors(errs []error) error {
	if len(errs) == 0 {
		return nil
	}

	errsStrs := make([]string, 0, len(errs))
	for _, e := range errs {
		errsStrs = append(errsStrs, e.Error())
	}
	return errors.New(strings.Join(errsStrs, ";"))
}


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