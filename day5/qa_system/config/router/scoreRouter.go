package router

import (
	"day5/qa_system/app/controllers/scoreController"
	"github.com/gin-gonic/gin"
)

func scoreRouterInit(r *gin.RouterGroup) {
	r.Any("/submit", scoreController.GetScore)
}
