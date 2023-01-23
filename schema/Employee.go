package schema

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	Id           string   `json:"id" gorm:"primaryKey;size:50;"`
	Code         string   `gorm:"size:20;unique" json:"code"`
	Email        string   `gorm:"size:255;not null;unique" json:"email"`
	Username     string   `gorm:"size:255;not null;unique" json:"username"`
	Nik          string   `gorm:"size:20;" json:"nik"`
	NickName     string   `gorm:"size:255;" json:"nickname"`
	FullName     string   `gorm:"size:255;" json:"fullname"`
	Picture      string   `gorm:"size:255;" json:"picture"`
	PhoneNumber  string   `gorm:"size:15;" json:"phone_number"`
	Address      string   `gorm:"type:text;" json:"address"`
	CompanyCode  *string  `json:"company_code" gorm:"size:20; unique"`
	Company      *Company `gorm:"foreignKey:CompanyCode;references:Code;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Department   string   `gorm:"type:text;" json:"department"`
	OfficeNumber string   `gorm:"type:text;" json:"office_number"`
	Expired      int      `gorm:"type:int;" json:"expired_time"`
	CreatedBy    string   `gorm:"size:13;" json:"created_by"`
	ExpiredAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	Deleted      bool   `gorm:"type:bool;default:false" json:"deleted"`
	UserCode     string `json:"user_code" gorm:"size:20; unique"`
	User         User   `gorm:"foreignKey:UserCode;references:Code;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	IsActive     bool   `gorm:"type:bool;default:true" json:"status"`
}

func (Employee) TableName() string {
	return "employees"
}

type JsonEmployeeCreate struct {
	Code            string   `json:"code" binding:"required"`
	Username        string   `json:"username" binding:"required"`
	Email           string   `json:"email" binding:"required"`
	Password        string   `json:"password" binding:"required"`
	Roles           []string `json:"role" binding:"required"`
	RoleApplication []string `json:"role_application" binding:"required"`
	CompanyId       string   `json:"company_id" binding:"required"`
	Divisions       []string `json:"division" binding:"required"`
	Application     []string `json:"application" binding:"required"`
	Nik             string   `json:"nik" binding:"required"`
	NickName        string   `json:"nickname" binding:"required"`
	FullName        string   `json:"fullname" binding:"required"`
	Picture         string   `json:"picture" binding:"required"`
	PhoneNumber     string   `json:"phone_number" binding:"required"`
	Address         string   `json:"address" binding:"required"`
	Department      string   `json:"department" binding:"required"`
	OfficeNumber    string   `json:"office_number" binding:"required"`
}
