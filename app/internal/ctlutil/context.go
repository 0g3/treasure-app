package ctlutil

import (
	"errors"
	"github.com/0g3/treasure-app/app/api/v1/model"
	"github.com/gin-gonic/gin"
)

func GetUserFromCtx(ctx *gin.Context) (*model.User, error) {
	u, exists := ctx.Get("user")
	if !exists {
		return nil, errors.New("contextにuserが存在しません")
	}
	return u.(*model.User), nil
}
