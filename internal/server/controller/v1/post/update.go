package post

import (
	"github.com/gin-gonic/gin"

	"github.com/gaulzhw/go-server/internal/pkg/constant"
	"github.com/gaulzhw/go-server/internal/pkg/log"
	"github.com/gaulzhw/go-server/pkg/core"
	"github.com/gaulzhw/go-server/pkg/errno"
	metav1 "github.com/gaulzhw/go-server/pkg/meta/v1"
)

// UpdateRequest specify fields can be updated for post resource.
type UpdateRequest struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
}

// Update update a post info by the post identifier.
func (p *PostController) Update(c *gin.Context) {
	log.L(c).Info("update post function called.")

	var r UpdateRequest

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	post, err := p.srv.Posts().Get(c, c.GetString(constant.XUsernameKey), c.Param("postID"), metav1.GetOptions{})
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	if r.Title != nil {
		post.Title = *r.Title
	}
	if r.Content != nil {
		post.Content = *r.Content
	}

	if err := post.Validate(); err != nil {
		core.WriteResponse(c, errno.ErrValidation, nil)

		return
	}

	// Save changed fields.
	if err := p.srv.Posts().Update(c, post, metav1.UpdateOptions{}); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
