package middleware

import (
	"class-schedule/utility"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Authorization(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		utility.Response(http.StatusUnauthorized, "No Token", nil, c)
		c.Abort()
		return
	}
	id, err := utility.ParseToken(token)
	if err != nil {
		utility.Response(http.StatusUnauthorized, "Bad Token", nil, c)
		c.Abort()
		return
	}
	id_int, ok := strconv.Atoi(id)
	if ok != nil {
		utility.Response(http.StatusInternalServerError, "Internal server error", nil, c)
		log.Println(ok)
		c.Abort()
	}
	c.Set("user_id", uint64(id_int))
	c.Next()
}
