package core

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gaulzhw/go-server/pkg/errno"
)

// ErrResponse defines the return messages when an error occurred.
// Reference will be omitted if it does not exist.
type ErrResponse struct {
	// Code defines the business error code.
	Code int `json:"code"`

	// Message contains the detail of this message.
	// This message is suitable to be exposed to external
	Message string `json:"message"`

	Data any `json:"data,omitempty"`
}

// WriteResponse write an error or the response data into http response body.
// It use errors.ParseCoder to parse any error into errors.Coder
// errors.Coder contains error code, user-safe error message and http status code.
func WriteResponse(c *gin.Context, err error, data any) {
	if err != nil {
		coder := errno.ParseCoder(err)
		c.JSON(coder.HTTPStatus(), ErrResponse{
			Code:    coder.Code(),
			Message: coder.String(),
			Data:    data,
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
