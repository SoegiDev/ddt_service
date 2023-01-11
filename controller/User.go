package controller

import (
	"ddtservice_agri/helper"
	"ddtservice_agri/model"
	"ddtservice_agri/schema"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"user": updateRoles})
}

func UserGetProfiles(context *gin.Context) { // Get model if exist
	id := context.Param("ID")
	user, err := model.UserFindById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	helper.Metauser(context, user.Id)
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
