package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonSuccessResponse(c *gin.Context, data interface{}, id, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  msg,
		"name": id,
		"data": data,
	})
}
