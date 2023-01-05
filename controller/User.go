package controller

import (
	"ddtservice_agri/model"
	"ddtservice_agri/schema"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserAddDivisions(context *gin.Context) {
	id := context.Param("ID")
	var input schema.AddDivision

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input_division := input.Division
	data_divisions, err := model.FindDivisionMapByName(input_division)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	divisionMap := []schema.Division{}
	for _, element := range data_divisions {
		divisionMap = append(divisionMap, schema.Division{Id: element.Id})
	}
	users, err := model.FindUserById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"msg": "User Not Found"})
		return
	}
	updateDivision, err := users.UpdateDivision(id, divisionMap)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"user": updateDivision})
}
