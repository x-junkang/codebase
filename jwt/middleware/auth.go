package middleware

import (
	"codebase/jwt/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//get authorization header
		tokenString := ctx.GetHeader("Authorization")

		//validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		claims, err := model.ParseToken(tokenString)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		//通过验证后获取claims中的userID
		userId := claims.ID

		//如果用户存在 将user信息存入上下文
		ctx.Set("id", userId)
		ctx.Next()
	}
}
