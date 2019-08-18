package middleware

import (
	"firebase.google.com/go/auth"
	"github.com/0g3/treasure-app/app/api/v1/repository"
	"github.com/0g3/treasure-app/app/config"
	"github.com/0g3/treasure-app/app/internal/ctlutil"
	"github.com/0g3/treasure-app/app/internal/firebaseutil"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

var firebaseClient *auth.Client

func init() {
	var err error
	firebaseClient, err = firebaseutil.NewAuthClient(config.Env().KeyFilePath)
	if err != nil {
		log.Fatal("firebase clientの生成に失敗:", firebaseClient)
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idToken := strings.Replace(ctx.Request.Header.Get("Authorization"), "Bearer ", "", 1)
		token, err := firebaseClient.VerifyIDToken(ctx, idToken)
		if err != nil {
			ctlutil.Err(ctx, 401, err)
			ctx.Abort()
			return
		}

		fbu, err := firebaseClient.GetUser(ctx, token.UID)
		if err != nil {
			ctlutil.Err(ctx, 500, err)
			ctx.Abort()
			return
		}

		u, err := repository.User.Get(fbu.UID)
		if err != nil {
			ctlutil.Err(ctx, 500, err)
			ctx.Abort()
			return
		}

		ctx.Set("user", u)
	}
}
