package ctlutil

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"log"
)

func Err(ctx *gin.Context, code int, err error) {
	ctx.Status(code)
	txt := fmt.Sprintf("⚠️️ status code: %d\n%+v", code, errors.WithStack(err))
	log.Println(txt)
}
