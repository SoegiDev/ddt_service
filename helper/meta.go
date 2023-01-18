package helper

import (
	"ddtservice_agri/model"
	"ddtservice_agri/schema"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Metauser(context *gin.Context, id string) {
	dt, err := model.UserMetaFindById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	getDivisiEstate := []string{}
	for _, element := range dt.Divisions {
		getDivisiEstate = append(getDivisiEstate, element.EstateCode)
	}
	data_estate, err := model.EstateFindDivisionById(getDivisiEstate)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	meta_roles := []schema.MetaRoles{}
	for _, element := range dt.Roles {
		meta_roles = append(meta_roles, schema.MetaRoles{
			Id:          element.Id,
			Name:        element.Name,
			Description: element.Description})
	}
	meta_accounts := []schema.MetaAccount{}
	for _, element := range dt.Accounts {
		meta_role_app := []schema.MetaRoleApp{}
		for _, el := range element.RoleApplications {
			meta_role_app = append(meta_role_app, schema.MetaRoleApp{
				Id:          el.Id,
				Name:        el.Name,
				Description: el.Description})
		}
		map_app := schema.MetaApp{
			Id:              element.Application.Id,
			Name:            element.Application.Name,
			Description:     element.Application.Description,
			UpdatedNote:     element.Application.UpdatedNote,
			Version:         element.Application.Version,
			AppPackage:      element.Application.AppPackage,
			AppPackageClass: element.Application.AppPackageClass,
			AssetApk:        element.Application.AssetApk,
			AssetIcon:       element.Application.AssetIcon,
			IsActive:        element.Application.IsActive}
		meta_accounts = append(meta_accounts, schema.MetaAccount{
			Id:              element.Id,
			UserCode:        element.UserCode,
			RoleApplication: meta_role_app,
			Application:     map_app})
	}
	meta_employees := []schema.MetaEmployee{}
	for _, element := range dt.Employees {
		map_comp := schema.MetaCompany{
			Id:          element.Company.Id,
			Code:        element.Company.Code,
			Name:        element.Company.Name,
			Description: element.Company.Description,
			PhoneNumber: element.Company.PhoneNumber,
			Fax:         element.Company.Fax,
			Sector:      element.Company.Sector,
			Domain:      element.Company.Domain,
			Address:     element.Company.Address}
		meta_employees = append(meta_employees, schema.MetaEmployee{
			Id:           element.Id,
			Code:         element.Code,
			Email:        element.Email,
			Username:     element.Username,
			Nik:          element.Nik,
			NickName:     element.NickName,
			FullName:     element.FullName,
			Picture:      element.Picture,
			PhoneNumber:  element.PhoneNumber,
			Address:      element.Address,
			CompanyCode:  element.CompanyCode,
			Department:   element.Department,
			OfficeNumber: element.OfficeNumber,
			Expired:      element.Expired,
			Company:      map_comp})
	}
	map_estate := []schema.MetaEstate{}
	for _, element := range data_estate {
		map_division := []schema.MetaDivision{}
		fmt.Println(data_estate)
		map_company := schema.MetaCompany{
			Id:          element.Company.Id,
			Code:        element.Company.Code,
			Name:        element.Company.Name,
			Description: element.Company.Description,
			PhoneNumber: element.Company.PhoneNumber,
			Fax:         element.Company.Fax,
			Sector:      element.Company.Sector,
			Domain:      element.Company.Domain,
			Address:     element.Company.Address}
		for _, division := range dt.Divisions {
			map_gang := []schema.MetaGang{}
			for _, gg := range division.Gangs {
				map_gang = append(map_gang, schema.MetaGang{
					Id:   strconv.Itoa(int(gg.Id)),
					Code: gg.Code,
					Name: gg.Name})
			}
			map_division = append(map_division, schema.MetaDivision{
				Id:          division.Id,
				Code:        division.Code,
				Name:        division.Name,
				Description: division.Description,
				Gang:        map_gang})
		}

		map_estate = append(map_estate, schema.MetaEstate{
			Id:          element.Id,
			Code:        element.Code,
			Name:        element.Name,
			Company:     map_company,
			Description: element.Description,
			Division:    map_division})
	}

	meta_activities := []schema.MetaActivities{}
	for _, element := range dt.ActivityLogs {
		meta_activities = append(meta_activities, schema.MetaActivities{
			Id:          element.Id,
			Client:      element.Client,
			Duration:    element.Duration,
			Method:      element.Method,
			Status:      element.Status,
			Path:        element.Path,
			ReqBody:     element.ReqBody,
			PathOp:      element.PathOp,
			UserId:      element.UserId,
			Source:      element.Source,
			Application: element.Application,
			Referrer:    element.Referrer,
			RequestId:   element.RequestId,
			CreatedAt:   element.CreatedAt})

	}

	user_meta := schema.MetaUser{
		Id:           dt.Id,
		Code:         dt.Code,
		Username:     dt.Username,
		Email:        dt.Email,
		Password:     dt.Password,
		CreatedAt:    dt.CreatedAt,
		UpdatedAt:    dt.UpdatedAt,
		IsActive:     dt.IsActive,
		Estates:      map_estate,
		Employees:    meta_employees,
		Roles:        meta_roles,
		Accounts:     meta_accounts,
		ActivityLogs: meta_activities,
	}
	context.JSON(http.StatusOK, gin.H{"data": user_meta})
}
