package questionServices

import (
	"day5/qa_system/app/models"
	"day5/qa_system/config/database"
)

func GetQuestions(id int) ([]models.Question, error) {
	var questions []models.Question

	result := database.DB.Where(
		models.Question{
			QuestionnaireID: id,
		}).Find(&questions)
	if result.Error != nil {
		return nil, result.Error
	}

	return questions, nil
}
