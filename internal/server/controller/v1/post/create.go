package post

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/gaulzhw/go-server/internal/pkg/model/server/v1"

	"github.com/gaulzhw/go-server/internal/pkg/constant"
	"github.com/gaulzhw/go-server/internal/pkg/log"
	"github.com/gaulzhw/go-server/pkg/core"
	"github.com/gaulzhw/go-server/pkg/errno"
	metav1 "github.com/gaulzhw/go-server/pkg/meta/v1"
)

// Create add new post to the storage.
func (p *PostController) Create(c *gin.Context) {
	log.L(c).Info("post create function called.")

	var r v1.Post

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	// Add username
	r.Username = c.GetString(constant.XUsernameKey)

	if err := r.Validate(); err != nil {
		core.WriteResponse(c, errno.ErrValidation, nil)

		return
	}

	// Insert the post to the storage.
	if err := p.srv.Posts().Create(c, &r, metav1.CreateOptions{}); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, map[string]string{"postID": r.PostID})
}
