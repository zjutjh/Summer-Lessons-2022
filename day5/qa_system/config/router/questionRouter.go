package router

import (
	"day5/qa_system/app/controllers/questionController"
	"github.com/gin-gonic/gin"
)

func questionRouterInit(r *gin.RouterGroup) {
	r.Any("/getQuestion", questionController.GetQuestions)
}
