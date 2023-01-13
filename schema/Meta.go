package schema

import (
	"time"
)

type MetaUser struct {
	Id           string
	Code         string
	Username     string
	Email        string
	Password     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	IsActive     bool
	Estates      []MetaEstate
	Employees    []MetaEmployee
	Roles        []MetaRoles
	Accounts     []MetaAccount
	ActivityLogs []MetaActivities
}

type MetaActivities struct {
	Id          uint
	Client      string
	Duration    string
	Method      string
	Status      string
	Path        string
	ReqBody     string
	PathOp      string
	UserId      string
	Source      string
	Application string
	Referrer    string
	RequestId   string
	CreatedAt   time.Time
}

type MetaAccount struct {
	Id              string
	UserId          string
	RoleApplication []MetaRoleApp
	Application     MetaApp
}

type MetaRoleApp struct {
	Id          uint
	Name        string
	Description string
}

type MetaApp struct {
	Id              string
	Name            string
	Description     string
	UpdatedNote     string
	Version         string
	AppPackage      string
	AppPackageClass string
	AssetApk        string
	AssetIcon       string
	IsActive        bool
}

type MetaRoles struct {
	Id          uint
	Name        string
	Description string
}

type MetaEmployee struct {
	Id           string
	Code         string
	Email        string
	Username     string
	Nik          string
	NickName     string
	FullName     string
	Picture      string
	PhoneNumber  string
	Address      string
	CompanyId    string
	Department   string
	OfficeNumber string
	Expired      int
	Company      MetaCompany
}

type MetaCompany struct {
	Id          string
	Code        string
	Name        string
	Description string
	PhoneNumber string
	Fax         string
	Sector      string
	Domain      string
	Address     string
}

type MetaEstate struct {
	Id          string
	Code        string
	Name        string
	Description string
	Company     MetaCompany
	Division    []MetaDivision
}

type MetaDivision struct {
	Id          string
	Code        string
	Name        string
	Description string
	Gang        []MetaGang
}

type MetaGang struct {
	Id   uint
	Code string
	Name string
}
