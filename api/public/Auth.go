package public

import (
	"ddtservice_agri/helper"
	"ddtservice_agri/model"
	"ddtservice_agri/schema"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Sign In
// @Description Sign In User
// @Accept  json
// @Param login body schema.SignInJsonSchema true "Login Schema "
// @Produce  json
// @Success 200 {object}  schema.LoginResponse
// @Router /sign_in [post]
func SignIn(context *gin.Context) {
	var input schema.SignInJsonSchema
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, schema.MsgResponse{Msg: err.Error()})
		return
	}
	if input.Username != "" {
		fmt.Println("Username tidak kosong")
	}
	user, err := model.UserFindByUsername(input.Username)

	fmt.Println(&user, err)
	if err != nil {
		context.JSON(http.StatusNotFound, schema.MsgResponse{Msg: "User Not Found"})
		return
	}

	err = user.ValidatePassword(input.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, schema.MsgResponse{Msg: "Wrong Password"})
		return
	}

	jwt, err := helper.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, schema.MsgResponse{Msg: err.Error()})
		return
	}
	context.JSON(http.StatusOK, schema.LoginResponse{Token: jwt})
}

//
// @Summary Sign Up
// @Description Sign Up User
// @Accept  json
// @Param login body schema.SignUpJsonSchema true "Login Schema "
// @Produce  json
// @Success 200 {object}  schema.MsgResponse
// @Router /sign_up [post]
func SignUp(context *gin.Context) {
	var input schema.SignUpJsonSchema
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input_roles := input.Role
	data_roles, err := model.FindRoleMapByName(input_roles)
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
		Id:        profileId,
		Username:  input.Username,
		Email:     input.Email,
		UserId:    savedUser.Id,
		CompanyId: input.CompanyId}

	savedProfile, err := emp.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(savedProfile)
	fmt.Println(savedUser)
	context.JSON(http.StatusCreated, schema.MsgResponse{Msg: "Register Completed"})
}
