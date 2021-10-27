package codatypes

import (
	"fmt"
	"github.com/artsafin/go-coda"
)

type UnexpectedFieldTypeErr struct {
	fieldID      string
	expectedType string
	row          *coda.Row
}

func (e UnexpectedFieldTypeErr) Error() string {
	v := e.row.Values[e.fieldID]
	return fmt.Sprintf("Unexpected type for field %s: expected %s. Got value: %#v of type %T", e.fieldID, e.expectedType, v, v)
}

type ErrContainer []error

func (c *ErrContainer) AddError(err error) {
	*c = append(*c, err)
}
func (c *ErrContainer) String() string {
	stringErr := ""
	for _, err := range *c {
		stringErr += err.Error() + "; "
	}

	return stringErr
}

// Deprecated: PanicIfError is deprecated - return error instead
func (c *ErrContainer) PanicIfError() {
	if len(*c) == 0 {
		return
	}

	panic(c.String())
}

func NewErrorContainer() ErrContainer {
	return make([]error, 0)
}
