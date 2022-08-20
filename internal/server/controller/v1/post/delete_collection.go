package post

import (
	"github.com/gin-gonic/gin"

	"github.com/gaulzhw/go-server/internal/pkg/constant"
	"github.com/gaulzhw/go-server/internal/pkg/log"
	"github.com/gaulzhw/go-server/pkg/core"
	metav1 "github.com/gaulzhw/go-server/pkg/meta/v1"
)

// DeleteCollection batch delete posts by multiple post ids.
func (p *PostController) DeleteCollection(c *gin.Context) {
	log.L(c).Info("batch delete post function called.")

	postIDs := c.QueryArray("postID")

	if err := p.srv.Posts().DeleteCollection(c, c.GetString(constant.XUsernameKey), postIDs, metav1.DeleteOptions{}); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
