package common

import (
	"github.com/pkg/errors"
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
