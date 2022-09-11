package controller

import (
	"class-schedule/db/model"
	"class-schedule/utility"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var data LoginData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utility.ResponseBadRequest(c)
		return
	}
	var user *model.User
	user, err = model.GetUserByEmail(data.Email)
	if user == nil {
		utility.Response(http.StatusNotFound, "User not found", nil, c)
		return
	}
	if !utility.PasswordVerify(data.Password, user.Password) {
		utility.Response(http.StatusBadRequest, "Wrong Password", nil, c)
		return
	}
	fmt.Println("Login: ", user.UserID)
	token := utility.GenerateStandardJwt(&utility.JwtData{
		ID: strconv.Itoa(int(user.UserID)),
	})
	utility.Response(http.StatusOK, "OK", gin.H{"token": token}, c)
}

func Register(c *gin.Context) {
	var data model.User
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err)
		utility.ResponseBadRequest(c)
		return
	}
	data.Password, err = utility.PasswordHash(data.Password)
	if err != nil {
		log.Println(err)
		utility.ResponseInternalServerError(c)
		return
	}
	err = model.AddUser(&data)
	if err != nil {
		log.Println(err)
		utility.ResponseInternalServerError(c)
		return
	}
	utility.Response(http.StatusOK, "OK", nil, c)
}

func UpdateUser(c *gin.Context) {
	var data model.User
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err)
		utility.ResponseBadRequest(c)
		return
	}
	data.UserID = c.GetUint64("user_id")
	err = model.UpdateUser(&data)
}

func GetUserInfo(c *gin.Context) {
	userID, _ := c.Get("user_id")
	user, err := model.GetUserByID(userID.(uint64))
	if err != nil {
		log.Println(err)
		utility.ResponseInternalServerError(c)
		return
	}
	utility.Response(http.StatusOK, "OK", gin.H{
		"info": user,
	}, c)
}
