package controller

import (
	"class-schedule/db/model"
	"class-schedule/utility"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddClass(c *gin.Context) {
	user_id, _ := c.Get("user_id")
	var data model.Class
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err)
		utility.ResponseBadRequest(c)
		return
	}
	data.UserID = user_id.(uint64)
	err = model.AddClass(&data)
	if err != nil {
		log.Println(err)
		utility.ResponseInternalServerError(c)
		return
	}
	utility.ResponseOK(c, nil)
}

func DeleteClass(c *gin.Context) {
	classID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		utility.ResponseBadRequest(c)
	}
	err = model.DeleteClassByID(uint64(classID))
	if err == gorm.ErrRecordNotFound {
		utility.Response(http.StatusNotFound, "Class not found", nil, c)
		return
	} else if err != nil {
		utility.ResponseInternalServerError(c)
		return
	}
	utility.ResponseOK(c, nil)
}

func UpdateClass(c *gin.Context) {
	var class model.Class
	err := c.ShouldBindJSON(&class)
	fmt.Println(class)
	if err != nil {
		utility.ResponseBadRequest(c)
		return
	}
	err = model.UpdateClass(&class)
	if err == gorm.ErrRecordNotFound {
		utility.Response(http.StatusNotFound, "Class not found", nil, c)
		return
	} else if err != nil {
		utility.ResponseInternalServerError(c)
		return
	}
	utility.ResponseOK(c, nil)
}

func GetClasses(c *gin.Context) {
	user_id, _ := c.Get("user_id")
	classes, err := model.GetClassesByUserID(user_id.(uint64))
	if err != nil {
		utility.ResponseInternalServerError(c)
		return
	}
	utility.ResponseOK(c, gin.H{
		"class": classes,
	})
}
