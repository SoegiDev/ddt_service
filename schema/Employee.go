package schema

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	Id           string `json:"id" gorm:"primaryKey;size:50;"`
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

type EmployeeAdd struct {
	Id          string `json:"id"`
	NickName    string `json:"nickname"`
	FullName    string `json:"fullname"`
	PhoneNumber string `json:"phone_number"`
	Picture     string `json:"picture"`
	Address     string `json:"address"`
	UserId      string `json:"user_id"`
}

func (Employee) TableName() string {
	return "employees"
}
