package nameServices

import (
	"day5/qa_system/app/apiExpection"
	"day5/qa_system/app/models"
	"day5/qa_system/config/database"
	"log"
)

func GetName(id string, time int64) (*string, error) {
	var name []models.NameMap

	result := database.DB.Where(
		models.NameMap{
			ID: id,
		}).Find(&name)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		log.Println("name_map parameter error")
		return nil, apiExpection.ParamError
	}
	time_ := &name[0].Time.Time
	name_ := &name[0].Name
	if time > time_.Unix() {
		return name_, apiExpection.TimeOut
	}

	return name_, nil
}
