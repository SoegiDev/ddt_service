package controller

import (
	"ddtservice_agri/helper"
	"ddtservice_agri/model"
	"ddtservice_agri/schema"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create Division
// @Description Create Division
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param Estate body schema.RequestDivision true "Estate Schema "
// @Produce  json
// @Success 200 {object}  schema.MsgResponse
// @Router /division [post]
// @Tags (C) Division
func DivisionAddNew(context *gin.Context) {
	var input schema.RequestDivision
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	unique, _ := helper.GenerateDivisionId(3)
	newData := model.Division{
		Id:          unique,
		Code:        input.Code,
		Name:        input.Name,
		Description: input.Description,
		EstateCode:  input.EstateCode}
	savedEntry, err := newData.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(savedEntry)
	context.JSON(http.StatusCreated, schema.MsgResponse{Msg: "Division Completed"})
}

// @Summary Update Division
// @Description Update Division
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param Estate body schema.RequestDivisionUpdate true "Company Schema "
// @Produce  json
// @Success 200 {object}  schema.MsgResponse
// @Router /division [patch]
// @Tags (C) Division
func DivisionUpdate(context *gin.Context) {
	id := context.Param("ID")
	data_entries, err := model.DivisionFindById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	var input schema.RequestDivisionUpdate
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedEntry, err := data_entries.DivisionChangeData(id, input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(updatedEntry)
	context.JSON(http.StatusOK, schema.MsgResponse{Msg: "DIvision Update Completed"})
}

// @Summary Get Division By ID
// @Description Get Division ID
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id   path string  true  "Division ID OR Division Code"
// @Produce  json
// @Success 200 {object}  schema.DivisionResponse
// @Router /division/{id} [get]
// @Tags (C) Division
func DivisionById(context *gin.Context) {
	id := context.Param("ID")
	get, err := model.DivisionFindById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	dtResponse := schema.DivisionResponse{
		Id:          get.Id,
		Code:        get.Code,
		Name:        get.Name,
		Description: get.Description,
		EstateCode:  get.EstateCode,
		Gangs:       get.Gangs}
	context.JSON(http.StatusOK, gin.H{"data": dtResponse})
}

// @Summary Get All Division
// @Description Get Division All
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Produce  json
// @Success 200 {array}  schema.DivisionResponse
// @Router /division [get]
// @Tags (C) Division
func DivisionByAll(context *gin.Context) {
	getData, err := model.DivisionFindAll()
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	allDivision := []schema.DivisionResponse{}
	for _, get := range getData {
		allDivision = append(allDivision, schema.DivisionResponse{
			Id:          get.Id,
			Code:        get.Code,
			Name:        get.Name,
			Description: get.Description,
			EstateCode:  get.EstateCode,
			Gangs:       get.Gangs})
	}
	context.JSON(http.StatusOK, gin.H{"data": allDivision})
}
