package middleware

import (
	"github.com/gin-gonic/gin"
	"learn_go/learnGo/common"
	"learn_go/learnGo/model"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		//获取authorization header
		tokenString := context.GetHeader("Authorization")

		// validate token format
		if !strings.HasPrefix(tokenString, "Bearer") || tokenString == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParserToken(tokenString)
		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort()
			return
		} else {
			//验证通过后获取userID
			userID := claims.UserID
			db := common.GetDB()
			var user model.User
			db.First(&user, userID)
			// 用户不存在
			if userID == 0 {
				context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
				context.Abort()
				return
			} else {
				//将user信息写入上下文
				context.Set("user", user)
				context.Next()
			}

		}
	}
}
