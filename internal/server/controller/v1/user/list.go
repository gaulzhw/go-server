package user

import (
	"github.com/gin-gonic/gin"

	"github.com/gaulzhw/go-server/internal/pkg/log"
	"github.com/gaulzhw/go-server/pkg/core"
	"github.com/gaulzhw/go-server/pkg/errno"
	metav1 "github.com/gaulzhw/go-server/pkg/meta/v1"
)

// List list the users in the storage.
// Only administrator can call this function.
func (u *UserController) List(c *gin.Context) {
	log.L(c).Info("list user function called.")

	var r metav1.ListOptions
	if err := c.ShouldBindQuery(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	users, err := u.srv.Users().List(c, r)
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, users)
}
