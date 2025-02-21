package handler

import (
	"backend/pkg/handler/errorResponse"
	"errors"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	handlerTimeout      = 30 * time.Second
	authorizationHeader = "Authorization"
	userCTX             = "userId"
)

func getUserId(c *gin.Context) (int, error) {
	userId, ok := c.Get(userCTX)
	if !ok {
		errorResponse.NewErrorResponse(c, http.StatusInternalServerError, errors.New("user not found").Error())
		return 0, errors.New("user id not found")
	}
	id, ok := userId.(int)
	if !ok {
		errorResponse.NewErrorResponse(c, http.StatusInternalServerError, errors.New("user id is invalid type").Error())
		return 0, errors.New("user id is of invalid type")
	}
	return id, nil
}

func (h *Handler) timeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(handlerTimeout),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(
			func(c *gin.Context) {
				c.JSON(http.StatusRequestTimeout, map[string]interface{}{
					"error": "request timeout",
				})
			}),
	)
}
