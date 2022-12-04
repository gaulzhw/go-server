package user

import (
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"

	"github.com/gaulzhw/go-server/pkg/core"
	metav1 "github.com/gaulzhw/go-server/pkg/meta/v1"
)

// Get get an user by the user identifier.
func (c *Controller) Get(ctx *gin.Context) {
	klog.Info("get user function called.")

	user, err := c.svc.Users().Get(ctx, ctx.Param("name"), metav1.GetOptions{})
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, user)
}
