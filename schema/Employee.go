package schema

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	Id            string   `json:"id" gorm:"primaryKey;size:50;"`
	Code          string   `gorm:"size:20;unique" json:"code"`
	Email         string   `gorm:"size:255;not null;unique" json:"email"`
	Username      string   `gorm:"size:255;not null;unique" json:"username"`
	Nik           string   `gorm:"size:20;" json:"nik"`
	NickName      string   `gorm:"size:255;" json:"nickname"`
	FullName      string   `gorm:"size:255;" json:"fullname"`
	Picture       string   `gorm:"size:255;" json:"picture"`
	PhoneNumber   string   `gorm:"size:15;" json:"phone_number"`
	Address       string   `gorm:"type:text;" json:"address"`
	CompanyCode   *string  `json:"company_code" gorm:"size:20; unique"`
	Company       *Company `gorm:"foreignKey:CompanyCode;references:Code;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Department    string   `gorm:"type:text;" json:"department"`
	OfficeNumber  string   `gorm:"type:text;" json:"office_number"`
	Expired       int      `gorm:"type:int;" json:"expired_time"`
	FiscalYear    string   `gorm:"size:10;" json:"fiscal_year"`
	FiscalPeriod  string   `gorm:"size:10;" json:"fiscal_period"`
	JobPos        string   `gorm:"size:10;" json:"job_pos"`
	EmployeeType  string   `gorm:"size:10;" json:"type"`
	CostCenter    string   `gorm:"size:10;" json:"cost_center"`
	Mandt         string   `gorm:"size:10;" json:"mandt"`
	ValidFrom     string   `gorm:"size:20;" json:"valid_from"`
	HarvesterCode string   `gorm:"size:10;" json:"harvester_code"`
	GradeRate     string   `gorm:"size:20;" json:"grade_rate"`
	CreatedBy     string   `gorm:"size:13;" json:"created_by"`
	ExpiredAt     time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
	Deleted       bool    `gorm:"type:bool;default:false" json:"deleted"`
	UserCode      *string `json:"user_code" gorm:"size:20; unique"`
	User          *User   `gorm:"foreignKey:UserCode;references:Code;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	IsActive      bool    `gorm:"type:bool;default:true" json:"status"`
}

func (Employee) TableName() string {
	return "employees"
}

type JsonEmployeeCreate struct {
	Code            string   `json:"code" binding:"required"`
	Username        string   `json:"username" binding:"required"`
	Email           string   `json:"email" binding:"required"`
	Password        string   `json:"password" binding:"required"`
	Roles           []string `json:"role"`
	RoleApplication []string `json:"role_application"`
	CompanyId       string   `json:"company_id"`
	Divisions       []string `json:"division"`
	Application     []string `json:"application"`
	Nik             string   `json:"nik"`
	NickName        string   `json:"nickname"`
	FullName        string   `json:"fullname"`
	Picture         string   `json:"picture"`
	PhoneNumber     string   `json:"phone_number"`
	Address         string   `json:"address"`
	Department      string   `json:"department"`
	OfficeNumber    string   `json:"office_number"`
}

type JsonEmployeeUpdate struct {
	Code            string   `json:"code" binding:"required"`
	Username        string   `json:"username" binding:"required"`
	Email           string   `json:"email" binding:"required"`
	Password        string   `json:"password"`
	Roles           []string `json:"role"`
	RoleApplication []string `json:"role_application"`
	CompanyId       string   `json:"company_id"`
	Divisions       []string `json:"division"`
	Application     []string `json:"application"`
	Nik             string   `json:"nik"`
	NickName        string   `json:"nickname"`
	FullName        string   `json:"fullname"`
	Picture         string   `json:"picture"`
	PhoneNumber     string   `json:"phone_number"`
	Address         string   `json:"address"`
	Department      string   `json:"department"`
	OfficeNumber    string   `json:"office_number"`
	UserId          string   `json:"user_id"`
}
