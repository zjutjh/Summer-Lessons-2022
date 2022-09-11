package model

import (
	"class-schedule/db"
	"fmt"

	"gorm.io/gorm/clause"
)


type Class struct {
	UserID  uint64 `json:"-"`
	ClassID uint64 `gorm:"primaryKey" json:"class_id"`
	Name    string `gorm:"not null" json:"name" validate:"required"`
	Teacher string `json:"teacher"`
	Week    string `json:"week" validate:"required" gorm:"not null"`
	Time    string `json:"time" validate:"required" gorm:"not null"`
	WeekDay string `json:"week_day" validate:"required" gorm:"not null"`
	Address string `json:"address" validate:"required" gorm:"not null"`
}

func AddClass(class *Class) error {
	fmt.Println(class)
	return db.DB.Create(class).Error
}

func DeleteClassByID(ID uint64) error {
	c, err := GetClassByID(ID)
	if err != nil {
		return err
	}
	return db.DB.Select(clause.Associations).Delete(c).Error
}

func UpdateClass(class *Class) error {
	var c Class
	err := db.DB.Where("class_id = ?", class.ClassID).Take(&c).Error
	if err != nil {
		return err
	}
	class.ClassID = c.ClassID
	return db.DB.Omit("ClassID", "UserID").Updates(class).Error
}
func GetClassByID(ID uint64) (*Class, error) {
	var class Class
	err := db.DB.Where("class_id = ?", ID).First(&class).Error
	return &class, err
}

func GetClassesByUserID(userID uint64) ([]Class, error) {
	var classes []Class
	err := db.DB.Where("user_id = ?", userID).Find(&classes).Error
	return classes, err
}
