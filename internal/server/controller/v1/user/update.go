package user

import (
	"github.com/gin-gonic/gin"

	"github.com/gaulzhw/go-server/internal/pkg/log"
	"github.com/gaulzhw/go-server/pkg/core"
	"github.com/gaulzhw/go-server/pkg/errno"
	metav1 "github.com/gaulzhw/go-server/pkg/meta/v1"
)

// UpdateRequest specify fields can be updated for user resource.
type UpdateRequest struct {
	Nickname *string `json:"nickname"`
	Email    *string `json:"email"`
}

// Update update a user info by the user identifier.
func (u *UserController) Update(c *gin.Context) {
	log.L(c).Info("update user function called.")

	var r UpdateRequest

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	user, err := u.srv.Users().Get(c, c.Param("name"), metav1.GetOptions{})
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	if r.Nickname != nil {
		user.Nickname = *r.Nickname
	}
	if r.Email != nil {
		user.Email = *r.Email
	}

	if err := user.Validate(); err != nil {
		core.WriteResponse(c, errno.ErrValidation, nil)

		return
	}

	// Save changed fields.
	if err := u.srv.Users().Update(c, user, metav1.UpdateOptions{}); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, user)
}
