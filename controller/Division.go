package controller

import (
	"ddtservice_agri/helper"
	"ddtservice_agri/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DivisionAddNew(context *gin.Context) {
	var input model.Division
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	unique, _ := helper.GenerateDivisionId(3)
	input.Id = unique

	savedEntry, err := input.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}
