package router

import (
	"github.com/gin-gonic/gin"
)

func Route(e *gin.Engine) {
	r := e.Group("/v1")
	TweetRouter(r)
}
