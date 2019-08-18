package handler

import (
	"github.com/0g3/treasure-app/app/api/v1/model"
	"github.com/0g3/treasure-app/app/api/v1/repository"
	"github.com/0g3/treasure-app/app/internal/ctlutil"
	"github.com/0g3/treasure-app/app/pkg/bzwdclient"
	"github.com/gin-gonic/gin"
)

func CreateTweet(ctx *gin.Context) {
	req := new(model.Tweet)
	if !ctlutil.Bind(ctx, req) {
		return
	}

	t := new(model.Tweet)

	u, err := ctlutil.GetUserFromCtx(ctx)
	if err != nil {
		ctlutil.Err(ctx, 401, err)
		return
	}
	t.UserID = u.ID

	translatedBody, err := bzwdclient.Translate(req.Body, bzwdclient.LV1)
	if err != nil {
		ctlutil.Err(ctx, 500, err)
		return
	}
	t.Body = translatedBody

	if err := repository.Tweet.Create(t); err != nil {
		ctlutil.Err(ctx, 500, err)
		return
	}

	ctx.JSON(200, t)
}
