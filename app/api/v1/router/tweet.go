package router

import (
	"github.com/0g3/treasure-app/app/api/v1/handler"
	"github.com/0g3/treasure-app/app/api/v1/middleware"
	"github.com/gin-gonic/gin"
)

func TweetRouter(g *gin.RouterGroup) {
	tweets := g.Group("/tweets")
	{
		authRequiredTweets := tweets.Group("", middleware.AuthMiddleware())
		{
			authRequiredTweets.POST("", handler.CreateTweet)
		}
	}
}
