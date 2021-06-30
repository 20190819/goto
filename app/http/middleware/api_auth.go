package middleware

import (
	"github.com/gin-gonic/gin"
	"goto/pkg/jwt"
	"goto/pkg/utils"
)

func ApiAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		jwtserver := jwt.JwtServer{}

		if ok := jwtserver.AuthToken(token); !ok {
			utils.ErrorResponse(ctx, 401, -1, "登录授权失败")
			return
		} else {
			ctx.Next()
		}

	}
}
