package router

import (
	"github.com/0g3/treasure-app/app/api/v1/handler"
	"github.com/0g3/treasure-app/app/api/v1/middleware"
	"github.com/gin-gonic/gin"
)

func FollowRouter(g *gin.RouterGroup) {
	follows := g.Group("/follows")
	{
		authFollows := follows.Group("", middleware.AuthMiddleware())
		{
			authFollows.POST("", handler.CreateFollow)
		}
	}
}
