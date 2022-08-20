package post

import (
	"github.com/gin-gonic/gin"

	"github.com/gaulzhw/go-server/internal/pkg/constant"
	"github.com/gaulzhw/go-server/internal/pkg/log"
	"github.com/gaulzhw/go-server/pkg/core"
	metav1 "github.com/gaulzhw/go-server/pkg/meta/v1"
)

// Get get a post by the post identifier.
func (p *PostController) Get(c *gin.Context) {
	log.L(c).Info("get post function called.")

	post, err := p.srv.Posts().Get(c, c.GetString(constant.XUsernameKey), c.Param("postID"), metav1.GetOptions{})
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, post)
}
