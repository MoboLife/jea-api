package common

import "github.com/gin-gonic/gin"

func SendError(ctx *gin.Context, err error, code int) {
	ctx.AbortWithStatusJSON(code, ctx.Error(err).JSON())
}

