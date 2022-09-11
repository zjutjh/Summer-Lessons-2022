package utility

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(code int, msg string, data gin.H, c *gin.Context) {
	if data != nil {
		c.JSON(code, gin.H{
			"msg":  msg,
			"data": data,
		})
	} else {
		c.JSON(code, gin.H{
			"msg": msg,
		})
	}
}

func ResponseBadRequest(c *gin.Context) {
	Response(http.StatusBadRequest, "Bad request", nil, c)
}

func ResponseInternalServerError(c *gin.Context) {
	Response(http.StatusInternalServerError, "Internal server error", nil, c)
}

func ResponseOK(c *gin.Context, data gin.H) {
	Response(http.StatusOK, "OK", data, c)
}
