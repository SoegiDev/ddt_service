package controller

import (
	"ddtservice_agri/helper"
	"ddtservice_agri/model"
	"ddtservice_agri/schema"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddEmployee(context *gin.Context) {
	var input schema.EmployeeAdd
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := model.FindUserById(input.UserId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User Not Found"})
		return
	}

	input.UserId = user.Id
	input.Id = helper.GenerateSecureToken(10)
	addEmp := model.Employee{
		Id:          input.Id,
		UserId:      input.UserId,
		Address:     input.Address,
		PhoneNumber: input.PhoneNumber,
		FullName:    input.FullName,
		NickName:    input.NickName,
	}
	savedEntry, err := addEmp.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}

func GetEmployeeAll(context *gin.Context) {
	emp, err := model.FindEmployeeAll()
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"msg": "Employee Not Found"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": emp})
}

func GetEmployeeById(context *gin.Context) {
	id := context.Param("ID")
	emp, err := model.FindEmployeeById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"msg": "Employee Not Found"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": emp})
}
