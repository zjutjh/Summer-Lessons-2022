package scoreController

import (
	"day5/qa_system/app/apiExpection"
	"day5/qa_system/app/controllers/submitController"
	"day5/qa_system/app/models"
	"day5/qa_system/app/services/questionServices"
	"day5/qa_system/app/utils"
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

func GetScore(c *gin.Context) {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
	var scoreForm models.ScoreForm
	sum := 0

	err := c.ShouldBindJSON(&scoreForm)
	if err != nil {
		log.Println("request parameter error")
		_ = c.AbortWithError(200, apiExpection.ParamError)
		return
	}

	id, _ := strconv.Atoi(scoreForm.ID)
	questions, err := questionServices.GetQuestions(id)
	if err != nil {
		log.Println("table questions error")
		_ = c.AbortWithError(200, apiExpection.ServerError)
		return
	}

	if len(questions) != len(scoreForm.Ans) {
		log.Println("answer list length error")
		_ = c.AbortWithError(200, apiExpection.ParamError)
		return
	}
	sort.SliceStable(scoreForm.Ans, func(i, j int) bool {
		return scoreForm.Ans[i].ID < scoreForm.Ans[j].ID
	})
	for i, value := range questions {
		flag := true
		answers := strings.Split(value.Answer, ";")

		if len(answers) != len(scoreForm.Ans[i].Key) {
			continue
		}
		if len(scoreForm.Ans[i].Key) > 1 {
			sort.SliceStable(scoreForm.Ans[i].Key, func(j, k int) bool {
				return scoreForm.Ans[i].Key[j] < scoreForm.Ans[i].Key[k]
			})
		}
		for j := 0; j < len(answers) && flag; j++ {
			answer, _ := strconv.Atoi(answers[j])
			if answer != scoreForm.Ans[i].Key[j] {
				flag = false
				break
			}
		}
		if !flag {
			continue
		}
		sum++
	}
	score := math.Round(((100.0/float64(len(questions)))*float64(sum))*1e2+0.5) * 1e-2

	e := submitController.SubmitData(scoreForm.ID, scoreForm.Name, scoreForm.UID, strconv.FormatFloat(score, 'f', 2, 64))
	if e == apiExpection.ReSubmit {
		utils.JsonSuccessResponse(c, "请勿重复提交", "", "SUCCESS")
	} else if e != nil {
		log.Println("table submit error")
		_ = c.AbortWithError(200, apiExpection.ServerError)
	} else {
		utils.JsonSuccessResponse(c, score, "", "SUCCESS")
	}
}
