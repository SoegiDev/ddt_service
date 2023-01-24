package controller

import (
	"ddtservice_agri/model"
	"ddtservice_agri/schema"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create Gang
// @Description Create Gang
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param Estate body schema.RequestGang true "Estate Schema "
// @Produce  json
// @Success 200 {object}  schema.MsgResponse
// @Router /gang [post]
// @Tags (D) Gang
func GangAddNew(context *gin.Context) {
	var input schema.RequestGang
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newData := model.Gang{
		Code:         input.Code,
		Name:         input.Name,
		DivisionCode: input.DivisionCode}
	savedEntry, err := newData.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(savedEntry)
	context.JSON(http.StatusCreated, schema.MsgResponse{Msg: "Gang Completed"})
}

// @Summary Update Gang
// @Description Update Gang
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param Estate body schema.RequestGangUpdate true "Company Schema "
// @Produce  json
// @Success 200 {object}  schema.MsgResponse
// @Router /gang [patch]
// @Tags (D) Gang
func GangUpdate(context *gin.Context) {
	id := context.Param("ID")
	data_entries, err := model.GangFindById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	var input schema.RequestGangUpdate
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedEntry, err := data_entries.GangChangeData(id, input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(updatedEntry)
	context.JSON(http.StatusOK, schema.MsgResponse{Msg: "DIvision Update Completed"})
}

// @Summary Get Gang By ID
// @Description Get Gang ID
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id   path string  true  "Division ID OR Division Code"
// @Produce  json
// @Success 200 {object}  schema.GangResponse
// @Router /gang/{id} [get]
// @Tags (D) Gang
func GangById(context *gin.Context) {
	id := context.Param("ID")
	get, err := model.GangFindById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	dtResponse := schema.GangResponse{
		Id:           strconv.Itoa(int(get.Id)),
		Code:         get.Code,
		Name:         get.Name,
		DivisionCode: get.DivisionCode}
	context.JSON(http.StatusOK, dtResponse)
}

// @Summary Get All Gang
// @Description Get Gang All
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Produce  json
// @Success 200 {array}  schema.GangResponse
// @Router /gang [get]
// @Tags (D) Gang
func GangByAll(context *gin.Context) {
	getData, err := model.GangFindAll()
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	allGang := []schema.GangResponse{}
	for _, get := range getData {
		allGang = append(allGang, schema.GangResponse{
			Id:           strconv.Itoa(int(get.Id)),
			Code:         get.Code,
			Name:         get.Name,
			DivisionCode: get.DivisionCode})
	}
	context.JSON(http.StatusOK, allGang)
}
