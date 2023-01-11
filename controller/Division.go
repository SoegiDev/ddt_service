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

func DivisionUpdate(context *gin.Context) {
	// Get model if exist
	id := context.Param("ID")
	data_entries, err := model.DivisionFindById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input model.UpdateDivision
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ddt := model.Entry{Content: input.Content}
	// database.Database.Model(&entryContent).Updates(ddt)

	updatedEntry, err := data_entries.DivisionChangeData(id, input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": updatedEntry})
}
