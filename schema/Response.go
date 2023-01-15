package schema

type LoginResponse struct {
	Token string
}

type MsgResponse struct {
	Msg string
}

type EstateResponse struct {
	Id          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Company     Company
	Description string     `json:"description"`
	IsDeleted   bool       `json:"delete"`
	IsActive    bool       `json:"status"`
	Divisions   []Division `json:"Division"`
}

type CompanyResponse struct {
	Id          string `json:"id"`
	Code        string `json:"code" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Fax         string `json:"fax" binding:"required"`
	Sector      string `json:"sector" binding:"required"`
	Domain      string `json:"domain" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Employees   []Employee
	Estates     []Estate
}

type DivisionResponse struct {
	Id          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsDeleted   bool   `json:"delete"`
	IsActive    bool   `json:"status"`
	EstateId    string `json:"estate_id"`
	Gangs       []Gang `json:"gang"`
}

type GangResponse struct {
	Id         string `json:"id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	DivisionId string `json:"division_id"`
}

type ApplicationResponse struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	UpdatedNote     string `json:"updated_note"`
	Version         string `json:"version"`
	AppPackage      string `json:"app_package"`
	AppPackageClass string `json:"app_package_class"`
	AssetApk        string `json:"asset_apk"`
	AssetIcon       string `json:"asset_icon"`
}
