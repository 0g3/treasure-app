package handler

import (
	"errors"
	"github.com/0g3/treasure-app/app/api/v1/service"
	"github.com/0g3/treasure-app/app/internal/ctlutil"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func BatchGetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctlutil.Err(ctx, 400, errors.New("パスが不正"))
		return
	}
	res, err := service.BatchGetUser(id)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			ctlutil.Err(ctx, 404, err)
			return
		}
		ctlutil.Err(ctx, 500, err)
		return
	}
	ctx.JSON(200, res)
}
