package middleware

import (
	"net/http"

	"github.com/Anthya1104/glossika-be-oa-service/internal/app/util"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		// parse and validate the token
		claims, err := util.ParseToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errcode.WrapErr{
				HttpStatus: http.StatusUnauthorized,
				ErrCode:    errcode.UserInvalidAuth,
				RawErr:     err,
			})
			return
		}

		// put token into context
		c.Set("user_email", claims["email"])
		c.Set("user_id", claims["user_id"])
		c.Next()
	}
}
