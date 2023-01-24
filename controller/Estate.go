package controller

import (
	"ddtservice_agri/helper"
	"ddtservice_agri/model"
	"ddtservice_agri/schema"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create Estate
// @Description Create Estate
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param Estate body schema.RequestEstate true "Estate Schema "
// @Produce  json
// @Success 200 {object}  schema.MsgResponse
// @Router /estate [post]
// @Tags (B) Estate
func EstateAddNew(context *gin.Context) {
	var input schema.RequestEstate
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	unique, _ := helper.GenerateEstateId(3)
	newEstate := model.Estate{
		Id:          unique,
		Code:        input.Code,
		Name:        input.Name,
		Description: input.Description,
		CompanyCode: input.CompanyCode}
	savedEntry, err := newEstate.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(savedEntry)
	context.JSON(http.StatusCreated, schema.MsgResponse{Msg: "New Estate Completed"})
}

// @Summary Update Field Estate
// @Description Update Estate
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id   path string  true  "Estate ID OR Estate Code"
// @Param Estate body schema.RequestEstateUpdate true "Estate Schema "
// @Produce  json
// @Success 200 {object}  schema.MsgResponse
// @Router /estate/{id}/edit [patch]
// @Tags (B) Estate
func EstateUpdate(context *gin.Context) {
	id := context.Param("ID")
	data_entries, err := model.EstateFindById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	var input schema.RequestEstateUpdate
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedEntry, err := data_entries.EstateChangeData(id, input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(updatedEntry)
	context.JSON(http.StatusOK, schema.MsgResponse{Msg: "Estate Update Completed"})
}

// @Summary Get Estate By ID
// @Description Get Estate ID
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id   path string  true  "Estate ID OR Estate Code"
// @Produce  json
// @Success 200 {object}  schema.EstateResponse
// @Router /estate/{id} [get]
// @Tags (B) Estate
func EstateById(context *gin.Context) {
	id := context.Param("ID")
	get, err := model.EstateFindById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	estateResponse := schema.EstateResponse{
		Id:          get.Id,
		Code:        get.Code,
		Name:        get.Name,
		Description: get.Description,
		Company:     get.Company,
		Divisions:   get.Divisions,
		IsActive:    get.IsActive,
		IsDeleted:   get.IsDeleted}
	context.JSON(http.StatusOK, estateResponse)
}

// @Summary Get Estate
// @Description Get Estate
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Produce  json
// @Success 200 {array}  schema.EstateResponse
// @Router /estate [get]
// @Tags (B) Estate
func EstateByAll(context *gin.Context) {
	getData, err := model.EstateFindAll()
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	allEstate := []schema.EstateResponse{}
	for _, get := range getData {
		allEstate = append(allEstate, schema.EstateResponse{
			Id:          get.Id,
			Code:        get.Code,
			Name:        get.Name,
			Description: get.Description,
			Company:     get.Company,
			Divisions:   get.Divisions,
			IsActive:    get.IsActive,
			IsDeleted:   get.IsDeleted})
	}
	context.JSON(http.StatusOK, allEstate)
}
