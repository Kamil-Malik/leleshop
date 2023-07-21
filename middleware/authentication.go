package middleware

import (
	"leleshop/util"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := util.VerifyToken(ctx)

		if err != nil {
			return
		}

		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}
