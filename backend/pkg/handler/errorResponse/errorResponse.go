package errorResponse

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorWeb struct {
	Message string `json:"message"`
}

func NewErrorResponse(ctx *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	ctx.AbortWithStatusJSON(statusCode, errorWeb{message})
}
