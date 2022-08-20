package post

import (
	"github.com/gin-gonic/gin"

	"github.com/gaulzhw/go-server/internal/pkg/constant"
	"github.com/gaulzhw/go-server/internal/pkg/log"
	"github.com/gaulzhw/go-server/pkg/core"
	metav1 "github.com/gaulzhw/go-server/pkg/meta/v1"
)

// Delete delete a post by the post identifier.
func (p *PostController) Delete(c *gin.Context) {
	log.L(c).Info("delete post function called.")

	if err := p.srv.Posts().Delete(c, c.GetString(constant.XUsernameKey), c.Param("postID"), metav1.DeleteOptions{Unscoped: true}); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
