package schema

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	Id               string `json:"id" gorm:"primaryKey;size:50;"`
	UserCode         string `json:"user_code"`
	ApplicationId    string `json:"application_id" gorm:"size:50;"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
	IsDeleted        bool              `gorm:"type:bool;default:false" json:"delete"`
	IsActive         bool              `gorm:"type:bool;default:true" json:"status"`
	RoleApplications []RoleApplication `gorm:"many2many:account_role_applications;"`
	Application      Application
}

func (Account) TableName() string {
	return "accounts"
}

type AssignRoleApplication struct {
	RoleApplication []string `json:"role_application"`
}

type UpdateAccount struct {
	ApplicationId string `json:"application_id"`
	UpdatedAt     time.Time
}
