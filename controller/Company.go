package controller

import (
	"ddtservice_agri/helper"
	"ddtservice_agri/model"
	"ddtservice_agri/schema"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create Company
// @Description Create Company
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param Estate body schema.RequestCompany true "Estate Schema "
// @Produce  json
// @Success 200 {object}  schema.MsgResponse
// @Router /company [post]
// @Tags (A) Company
func CompanyAddNew(context *gin.Context) {
	var input schema.RequestCompany
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	unique, _ := helper.GenerateCorporationId(3)
	newData := model.Company{
		Id:          unique,
		Code:        input.Code,
		Name:        input.Name,
		Description: input.Description,
		PhoneNumber: input.PhoneNumber,
		Fax:         input.Fax,
		Sector:      input.Sector,
		Domain:      input.Domain,
		Address:     input.Address}
	savedEntry, err := newData.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(savedEntry)
	context.JSON(http.StatusCreated, schema.MsgResponse{Msg: "Company Completed"})
}

// @Summary Update Company
// @Description Update Company
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param Estate body schema.RequestCompanyUpdate true "Company Schema "
// @Produce  json
// @Success 200 {object}  schema.MsgResponse
// @Router /estate [patch]
// @Tags (A) Company
func CompanyUpdate(context *gin.Context) {
	id := context.Param("ID")
	data_entries, err := model.CompanyFindById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	var input schema.RequestCompanyUpdate
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedEntry, err := data_entries.CompanyChangeData(id, input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(updatedEntry)
	context.JSON(http.StatusOK, schema.MsgResponse{Msg: "Company Update Completed"})
}

// @Summary Get Company By ID
// @Description Get Company ID
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id   path string  true  "Company ID OR Company Code"
// @Produce  json
// @Success 200 {object}  schema.CompanyResponse
// @Router /company/{id} [get]
// @Tags (A) Company
func CompanyById(context *gin.Context) {
	id := context.Param("ID")
	get, err := model.CompanyFindById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	dtResponse := schema.CompanyResponse{
		Id:          get.Id,
		Code:        get.Code,
		Name:        get.Name,
		Description: get.Description,
		PhoneNumber: get.PhoneNumber,
		Fax:         get.Fax,
		Sector:      get.Sector,
		Domain:      get.Domain,
		Address:     get.Address,
		Employees:   get.Employees,
		Estates:     get.Estates}
	context.JSON(http.StatusOK, gin.H{"data": dtResponse})
}

// @Summary Get All Company
// @Description Get Company All
// @Accept  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Produce  json
// @Success 200 {array}  schema.CompanyResponse
// @Router /company [get]
// @Tags (A) Company
func CompanyByAll(context *gin.Context) {
	getData, err := model.CompanyFindAll()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	allCompany := []schema.CompanyResponse{}
	for _, get := range getData {
		allCompany = append(allCompany, schema.CompanyResponse{
			Id:          get.Id,
			Code:        get.Code,
			Name:        get.Name,
			Description: get.Description,
			PhoneNumber: get.PhoneNumber,
			Fax:         get.Fax,
			Sector:      get.Sector,
			Domain:      get.Domain,
			Address:     get.Address,
			Employees:   get.Employees,
			Estates:     get.Estates})
	}
	context.JSON(http.StatusOK, gin.H{"data": allCompany})
}
