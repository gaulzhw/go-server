package user

import (
	"time"

	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"

	v1 "github.com/gaulzhw/go-server/pkg/apis/v1"
	"github.com/gaulzhw/go-server/pkg/core"
	metav1 "github.com/gaulzhw/go-server/pkg/meta/v1"
)

// Create add new user to the storage.
func (c *Controller) Create(ctx *gin.Context) {
	klog.Info("user create function called.")

	var r v1.User
	if err := ctx.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	//if errs := r.Validate(); len(errs) != 0 {
	//	core.WriteResponse(ctx, errors.WithCode(code.ErrValidation, errs.ToAggregate().Error()), nil)
	//	return
	//}

	r.Status = 1
	r.LoginedAt = time.Now()

	// Insert the user to the storage.
	if err := c.svc.Users().Create(ctx, &r, metav1.CreateOptions{}); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, r)
}
