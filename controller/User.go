package controller

import (
	"ddtservice_agri/helper"
	"ddtservice_agri/model"
	"ddtservice_agri/schema"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get Profile
// @Description Profile Meta User
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param ID  path      string  true  "Search User By ID"
// @Produce  json
// @Success 200 {object} schema.MetaUser
// @Router /user/:ID [get]
func UserGetProfiles(context *gin.Context) { // Get model if exist
	id := context.Param("ID")
	user, err := model.UserFindById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	helper.Metauser(context, user.Id)
}

// @Summary User Assign Role
// @Description User Assign Role User
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param ID  path      string  true  "Assign User By ID"
// @Param Register body schema.AssignUserRole true "Assign User Schema "
// @Produce  json
// @Success 200 {object} schema.MsgResponse
// @Router /user/:ID/assign_roles [get]
func UserAssignRole(context *gin.Context) {
	id := context.Param("ID")
	var input schema.AssignUserRole

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
	roleMap := []schema.Role{}
	for _, element := range data_roles {
		roleMap = append(roleMap, schema.Role{Id: element.Id})
	}
	users, err := model.UserFindById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"msg": "User Not Found"})
		return
	}
	updateRoles, err := users.UserAssignRoles(id, roleMap)
	fmt.Println(updateRoles)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, schema.MsgResponse{Msg: "Assign Completed"})
}

func UserAssignRoleApplication(context *gin.Context) { // Get model if exist
	id := context.Param("ID")
	var input schema.AssignRoleApplication

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input_role_application := input.RoleApplication
	data_roles, err := model.RoleAppMapByName(input_role_application)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	roleMap := []schema.RoleApplication{}
	for _, element := range data_roles {
		roleMap = append(roleMap, schema.RoleApplication{Id: element.Id})
	}
	accounts, err := model.AccountFindById(id, true)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"msg": "User Not Found"})
		return
	}
	updateRoleApplication, err := accounts.AccountSignRoleApp(accounts.Id, roleMap)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users, err := model.UserFindById(updateRoleApplication.UserId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"user": users})
}

func UserAssignDivisions(context *gin.Context) {
	id := context.Param("ID")
	var input schema.AddUserDivision

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input_division := input.Division
	data_divisions, err := model.DivisonFindMapById(input_division)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	divisionMap := []schema.Division{}
	for _, element := range data_divisions {
		divisionMap = append(divisionMap, schema.Division{Id: element.Id})
	}
	users, err := model.UserFindById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"msg": "User Not Found"})
		return
	}
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"msg": "User Not Found"})
		return
	}
	updateDivision, err := users.UserAssignDivision(id, divisionMap)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"user": updateDivision})
}

func UserAccountUpdate(context *gin.Context) {
	// Get model if exist
	id := context.Param("ID")
	data_entries, err := model.AccountFindById(id, true)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input model.UpdateAccount
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedEntry, err := data_entries.AccountChangeData(id, input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": updatedEntry})
}

func EmployeeAddNew(context *gin.Context) {
	var input schema.Register
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
