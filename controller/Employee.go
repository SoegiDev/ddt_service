package controller

import (
	"ddtservice_agri/helper"
	"ddtservice_agri/model"
	"ddtservice_agri/schema"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get Meta User By ID or Code
// @Description Get Meta User By ID or Code
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id   path string  true  "USER ID OR USER Code"
// @Produce  json
// @Success 200 {object} schema.MetaUser
// @Router /employee/{id} [get]
// @Tags Employee
func EmployeeMeta(context *gin.Context) { // Get model if exist
	id := context.Param("ID")
	user, err := model.UserFindById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	helper.Metauser(context, user.Id)
}

// @Summary Employee Add New
// @Description Employee Create New
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param Register body schema.JsonEmployeeCreate true "Employee Add New Schema "
// @Produce  json
// @Success 200 {object} schema.MsgResponse
// @Router /employee [post]
// @Tags Employee
func EmployeeAddNew(context *gin.Context) {
	var input schema.JsonEmployeeCreate
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, _ := helper.GenerateUserId(3)
	employeeId, _ := helper.GenerateProfileId(3)
	data_roles, err := model.FindRoleMapById(input.Roles)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roleMap := []schema.Role{}
	for _, element := range data_roles {
		roleMap = append(roleMap, schema.Role{Id: element.Id})
	}

	data_divisions, err := model.DivisonFindMapById(input.Divisions)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	divisionMap := []schema.Division{}
	for _, element := range data_divisions {
		divisionMap = append(divisionMap, schema.Division{Id: element.Id})
	}
	user := model.User{
		Id:        userID,
		Code:      input.Code,
		Roles:     roleMap,
		Email:     input.Email,
		Divisions: divisionMap,
		Username:  input.Username,
		Password:  input.Password,
	}
	savedUser, err := user.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data_apps, err := model.AppsFindMapById(input.Application)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, element := range data_apps {
		accountId, _ := helper.GenerateAccountId(3)
		divisionMap = append(divisionMap, schema.Division{Id: element.Id})
		data_role_app, err := model.RoleAppMapByName(input.RoleApplication)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		roleAppMap := []schema.RoleApplication{}
		for _, element := range data_role_app {
			roleAppMap = append(roleAppMap, schema.RoleApplication{Id: element.Id})
		}
		account := model.Account{
			Id:               accountId,
			UserId:           savedUser.Id,
			ApplicationId:    element.Id,
			RoleApplications: roleAppMap,
		}
		savedAccount, err := account.Save()
		fmt.Println(savedAccount)
		if err != nil {
			context.JSON(http.StatusBadRequest, schema.MsgResponse{Msg: err.Error()})
			return
		}
	}

	emp := model.Employee{
		Id:           employeeId,
		Code:         input.Code,
		Email:        input.Email,
		Username:     input.Username,
		Nik:          input.Nik,
		NickName:     input.NickName,
		FullName:     input.FullName,
		Picture:      input.Picture,
		PhoneNumber:  input.PhoneNumber,
		Address:      input.Address,
		CompanyId:    input.CompanyId,
		Department:   input.Department,
		OfficeNumber: input.OfficeNumber,
		UserId:       savedUser.Id}

	saveEmployee, err := emp.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(saveEmployee)
	context.JSON(http.StatusCreated, schema.MsgResponse{Msg: "Employee Created"})
}

// @Summary Employee Update
// @Description Employee Update Field
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param Register body schema.JsonEmployeeCreate true "Employee Update "
// @Produce  json
// @Success 200 {object} schema.MsgResponse
// @Router /employee/{id}/edit [patch]
// @Tags Employee
func EmployeeUpdate(context *gin.Context) {
	id := context.Param("ID")
	getData, err := model.EmployeeFindById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var input schema.JsonEmployeeCreate
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := getData.UserId
	employeeId := getData.Id
	data_roles, err := model.FindRoleMapById(input.Roles)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roleMap := []schema.Role{}
	for _, element := range data_roles {
		roleMap = append(roleMap, schema.Role{Id: element.Id})
	}

	data_divisions, err := model.DivisonFindMapById(input.Divisions)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	divisionMap := []schema.Division{}
	for _, element := range data_divisions {
		divisionMap = append(divisionMap, schema.Division{Id: element.Id})
	}
	user := model.User{
		Id:        userID,
		Code:      input.Code,
		Roles:     roleMap,
		Email:     input.Email,
		Divisions: divisionMap,
		Username:  input.Username,
		Password:  input.Password,
	}
	savedUser, err := user.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data_apps, err := model.AppsFindMapById(input.Application)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, element := range data_apps {
		accountId, _ := helper.GenerateAccountId(3)
		divisionMap = append(divisionMap, schema.Division{Id: element.Id})
		data_role_app, err := model.RoleAppMapByName(input.RoleApplication)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		roleAppMap := []schema.RoleApplication{}
		for _, element := range data_role_app {
			roleAppMap = append(roleAppMap, schema.RoleApplication{Id: element.Id})
		}
		account := model.Account{
			Id:               accountId,
			UserId:           savedUser.Id,
			ApplicationId:    element.Id,
			RoleApplications: roleAppMap,
		}
		savedAccount, err := account.Save()
		fmt.Println(savedAccount)
		if err != nil {
			context.JSON(http.StatusBadRequest, schema.MsgResponse{Msg: err.Error()})
			return
		}
	}

	emp := model.Employee{
		Id:           employeeId,
		Code:         input.Code,
		Email:        input.Email,
		Username:     input.Username,
		Nik:          input.Nik,
		NickName:     input.NickName,
		FullName:     input.FullName,
		Picture:      input.Picture,
		PhoneNumber:  input.PhoneNumber,
		Address:      input.Address,
		CompanyId:    input.CompanyId,
		Department:   input.Department,
		OfficeNumber: input.OfficeNumber,
		UserId:       savedUser.Id}

	saveEmployee, err := emp.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(saveEmployee)
	context.JSON(http.StatusCreated, schema.MsgResponse{Msg: "Employee Created"})
}
