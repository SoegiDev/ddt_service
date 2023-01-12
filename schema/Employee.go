package schema

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	Id           string `json:"id" gorm:"primaryKey;size:50;"`
	Code         string `gorm:"size:20;" json:"code"`
	Email        string `gorm:"size:255;not null;unique" json:"email"`
	Username     string `gorm:"size:255;not null;unique" json:"username"`
	Nik          string `gorm:"size:20;" json:"nik"`
	NickName     string `gorm:"size:255;" json:"nickname"`
	FullName     string `gorm:"size:255;" json:"fullname"`
	Picture      string `gorm:"size:255;" json:"picture"`
	PhoneNumber  string `gorm:"size:15;" json:"phone_number"`
	Address      string `gorm:"type:text;" json:"address"`
	CompanyId    string `json:"company_id" gorm:"size:50;"`
	Company      Company
	Department   string `gorm:"type:text;" json:"department"`
	OfficeNumber string `gorm:"type:text;" json:"office_number"`
	Expired      int    `gorm:"type:int;" json:"expired_time"`
	CreatedBy    string `gorm:"size:13;" json:"created_by"`
	ExpiredAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	Deleted      bool `gorm:"type:bool;default:false" json:"deleted"`
	UserId       string
	IsActive     bool `gorm:"type:bool;default:true" json:"status"`
}

type JsonEmployeeCreate struct {
	Code            string   `json:"code" binding:"required"`
	Username        string   `json:"username" binding:"required"`
	Email           string   `json:"email" binding:"required"`
	Password        string   `json:"password" binding:"required"`
	Roles           []string `json:"role" binding:"required"`
	RoleApplication []string `json:"role_application"`
	CompanyId       string   `json:"company_id" binding:"required"`
	Divisions       []string `json:"divisi" binding:"required"`
	Nik             string   `json:"nik"`
	NickName        string   `json:"nickname"`
	FullName        string   `json:"fullname"`
	Picture         string   `json:"picture"`
	PhoneNumber     string   `json:"phone_number"`
	Address         string   `json:"address"`
	Department      string   `gorm:"type:text;" json:"department"`
	OfficeNumber    string   `gorm:"type:text;" json:"office_number"`
	UpdatedAt       time.Time
	CreatedAt       time.Time
}

func (Employee) TableName() string {
	return "employees"
}
