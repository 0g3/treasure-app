package ctlutil

import "github.com/gin-gonic/gin"

func Bind(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBind(obj); err != nil {
		Err(c, 400, err)
		return false
	}
	return true
}
