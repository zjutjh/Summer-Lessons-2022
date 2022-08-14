package submitController

import (
	"day5/qa_system/app/apiExpection"
	"day5/qa_system/app/models"
	"day5/qa_system/app/services/submitService"
	"fmt"
	"log"
	"math"
	"strconv"
)

func SubmitData(ID, name, UID, score string) error {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
	submit, e := submitService.FetchSubmit(ID, UID)
	if e != nil {
		log.Println("fetch table submit error")
		return e
	}
	if submit.UID == UID {
		if submit.Num > 1 {
			return apiExpection.ReSubmit
		}
		scoreOld, _ := strconv.ParseFloat(submit.Score, 64)
		scoreNew, _ := strconv.ParseFloat(score, 64)
		maxScore := fmt.Sprintf("%.2f", math.Max(scoreOld, scoreNew))
		err := submitService.UpdateSubmit(models.Submit{
			PaperID: ID,
			Name:    name,
			UID:     UID,
			Score:   maxScore,
			Num:     2})
		if err != nil {
			log.Println("create table submit error")
			return err
		}
		return nil
	}

	err := submitService.CreateSubmit(models.Submit{
		PaperID: ID,
		Name:    name,
		UID:     UID,
		Score:   score,
		Num:     1})
	if err != nil {
		log.Println("create table submit error")
		return err
	}
	return nil
}
