package errno

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	validCodes = []int{http.StatusOK, http.StatusBadRequest, http.StatusInternalServerError}

	unknownCoder = &errCode{C: 1, HTTP: http.StatusInternalServerError, Ext: "An internal server error occurred"}
)

// codes contains a map of error codes to metadata.
var codes = make(map[int]Coder)

type errorno int

func (e errorno) Error() Coder {
	return codes[int(e)]
}

// Coder defines an interface for an eeror code detail information.
type Coder interface {
	// Code returns the code of the coder.
	Code() int

	// HTTPStatus should be used for the associated error code.
	HTTPStatus() int

	// String external user facing error text.
	String() string

	error
}

// errCode implements errors Coder interface.
type errCode struct {
	// C refers to the code of the ErrCode.
	C int

	// HTTP status that should be used for the associated error code.
	HTTP int

	// Ext the external user facing error text.
	Ext string
}

// Code returns the integer code of ErrCode.
func (c *errCode) Code() int {
	return c.C
}

// HTTPStatus returns the associated HTTP status code, if any. Otherwise, returns 500.
func (c *errCode) HTTPStatus() int {
	if c.HTTP == 0 {
		return http.StatusInternalServerError
	}
	return c.HTTP
}

// String implements stringer, returns the external error message.
func (c *errCode) String() string {
	return c.Ext
}

func (c *errCode) Error() string {
	return fmt.Sprintf("code: %d, error: %s", c.C, c.Ext)
}

func register(errno errorno, httpStatus int, message string) {
	isValid := false
	for _, validCode := range validCodes {
		if validCode == httpStatus {
			isValid = true
			break
		}
	}
	if isValid {
		panic("code is not valid, only supports 200, 400, 500")
	}

	if errno <= 0 {
		panic("code should must large than 0")
	}
	if _, ok := codes[int(errno)]; ok {
		panic(fmt.Sprintf("code %d already exists", errno))
	}
	codes[int(errno)] = &errCode{
		C:    int(errno),
		HTTP: httpStatus,
		Ext:  message,
	}
}

// ParseCoder parse any error into Coder.
func ParseCoder(err error) Coder {
	unwrapper := err
	for ; unwrapper != nil; unwrapper = errors.Unwrap(unwrapper) {
		if v, ok := unwrapper.(*errCode); ok {
			if coder, ok := codes[v.C]; ok {
				return coder
			}
			return unknownCoder
		}
	}
	return unknownCoder
}
