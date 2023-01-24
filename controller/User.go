package controller

import (
	"ddtservice_agri/helper"
	"ddtservice_agri/model"
	"ddtservice_agri/schema"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//
// @Summary User New
// @Description User New
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param login body schema.SignUpJsonSchema true "Login Schema "
// @Produce  json
// @Success 200 {object}  schema.MsgResponse
// @Router /user_new [post]
// @Tags User Credential
func UserAddNew(context *gin.Context) {
	var input schema.SignUpJsonSchema
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input_roles := input.Role
	data_roles, err := model.FindRoleMapById(input_roles)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, _ := helper.GenerateUserId(3)
	profileId, _ := helper.GenerateProfileId(3)

	roleMap := []schema.Role{}
	for _, element := range data_roles {
		roleMap = append(roleMap, schema.Role{Id: element.Id})
	}

	user := model.User{
		Id:       userID,
		Roles:    roleMap,
		Email:    input.Email,
		Username: input.Username,
		Password: input.Password,
	}

	savedUser, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	emp := model.Employee{
		Id:       profileId,
		Username: input.Username,
		Email:    input.Email,
		UserCode: &savedUser.Code}

	savedProfile, err := emp.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(savedProfile)
	fmt.Println(savedUser)
	context.JSON(http.StatusCreated, schema.MsgResponse{Msg: "Register Completed"})
}

// @Summary User Profile
// @Description User Profile
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id   path string  true  "USER ID OR USER Code"
// @Produce  json
// @Success 200 {object}  schema.MetaUser
// @Router /user_profile/{id} [get]
// @Tags User Credential
func User_Profile(context *gin.Context) { // Get model if exist
	id := context.Param("ID")
	user, err := model.UserFindById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	helper.Metauser(context, user.Id)
}

// @Summary User Profile All
// @Description User Profile All
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Produce  json
// @Success 200 {array}  schema.User
// @Router /user_profile [get]
// @Tags User Credential
func User_All(context *gin.Context) { // Get model if exist
	user, err := model.UserFindAll()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	context.JSON(http.StatusOK, user)
}
