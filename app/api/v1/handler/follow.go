package handler

import (
	"github.com/0g3/treasure-app/app/api/v1/model"
	"github.com/0g3/treasure-app/app/api/v1/repository"
	"github.com/0g3/treasure-app/app/internal/ctlutil"
	"github.com/gin-gonic/gin"
)

func CreateFollow(ctx *gin.Context) {
	req := new(model.Follow)
	if !ctlutil.Bind(ctx, req) {
		return
	}

	u, err := ctlutil.GetUserFromCtx(ctx)
	if err != nil {
		ctlutil.Err(ctx, 401, err)
		return
	}

	f := &model.Follow{
		FollowedUserID: req.FollowedUserID,
	}
	f.UserID = u.ID
	if err := repository.Follow.Create(f); err != nil {
		ctlutil.Err(ctx, 500, err)
		return
	}

	ctx.JSON(200, f)
}
