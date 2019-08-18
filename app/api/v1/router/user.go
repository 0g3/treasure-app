package router

import (
	"github.com/0g3/treasure-app/app/api/v1/handler"
	"github.com/gin-gonic/gin"
)

func UserRouter(g *gin.RouterGroup) {
	users := g.Group("/users")
	{
		users.GET("/:id/batch_get", handler.BatchGetUser)
	}
}
