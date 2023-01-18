package controller

import (
	"ddtservice_agri/helper"
	"ddtservice_agri/model"
	"ddtservice_agri/schema"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create Application
// @Description Create Application
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param Estate body schema.RequestApplication true "Estate Schema "
// @Produce  json
// @Success 200 {object}  schema.MsgResponse
// @Router /application [post]
// @Tags (E) Application
func ApplicationAddNew(context *gin.Context) {
	var input schema.RequestApplication
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	unique, _ := helper.GenerateAppId(3)
	newData := model.Applications{
		Id:              unique,
		Name:            input.Name,
		Description:     input.Description,
		UpdatedNote:     input.UpdatedNote,
		Version:         input.Version,
		AppPackage:      input.AppPackage,
		AppPackageClass: input.AppPackageClass,
		AssetApk:        input.AssetApk,
		AssetIcon:       input.AssetIcon}
	savedEntry, err := newData.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(savedEntry)
	context.JSON(http.StatusCreated, schema.MsgResponse{Msg: "Application Completed"})
}

// @Summary Update Application
// @Description Update Application
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param Estate body schema.RequestApplicationUpdate true "Application Schema "
// @Produce  json
// @Success 200 {object}  schema.MsgResponse
// @Router /application [patch]
// @Tags (E) Application
func ApplicationUpdate(context *gin.Context) {
	id := context.Param("ID")
	data_entries, err := model.AppsFindById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	var input schema.RequestApplicationUpdate
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedEntry, err := data_entries.AppsChangeData(id, input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(updatedEntry)
	context.JSON(http.StatusOK, schema.MsgResponse{Msg: "Application Update Completed"})
}

// @Summary Get Application By ID
// @Description Get Application ID
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id   path string  true  "Application ID OR Application Code"
// @Produce  json
// @Success 200 {object}  schema.ApplicationResponse
// @Router /application/{id} [get]
// @Tags (E) Application
func ApplicationById(context *gin.Context) {
	id := context.Param("ID")
	get, err := model.AppsFindById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	dtResponse := schema.ApplicationResponse{
		Id:              get.Id,
		Name:            get.Name,
		Description:     get.Description,
		UpdatedNote:     get.UpdatedNote,
		Version:         get.Version,
		AppPackage:      get.AppPackage,
		AppPackageClass: get.AppPackageClass,
		AssetApk:        get.AssetApk,
		AssetIcon:       get.AssetIcon}
	context.JSON(http.StatusOK, gin.H{"data": dtResponse})
}

// @Summary Get All Application
// @Description Get Application All
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Produce  json
// @Success 200 {array}  schema.ApplicationResponse
// @Router /application [get]
// @Tags (E) Application
func ApplicationByAll(context *gin.Context) {
	getData, err := model.AppsFindByAll()
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	allApplication := []schema.ApplicationResponse{}
	for _, get := range getData {
		allApplication = append(allApplication, schema.ApplicationResponse{
			Id:              get.Id,
			Name:            get.Name,
			Description:     get.Description,
			UpdatedNote:     get.UpdatedNote,
			Version:         get.Version,
			AppPackage:      get.AppPackage,
			AppPackageClass: get.AppPackageClass,
			AssetApk:        get.AssetApk,
			AssetIcon:       get.AssetIcon})
	}
	context.JSON(http.StatusOK, gin.H{"data": allApplication})
}
