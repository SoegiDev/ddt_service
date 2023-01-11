package controller

import (
	"ddtservice_agri/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GangAddNew(context *gin.Context) {
	var input model.Gang
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedEntry, err := input.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}

func GangUpdate(context *gin.Context) {
	// Get model if exist
	id := context.Param("ID")
	data_entries, err := model.EstateFindById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input model.UpdateEstate
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedEntry, err := data_entries.EstateChangeData(id, input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": updatedEntry})
}
