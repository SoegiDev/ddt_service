package schema

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	Id            string `json:"id" gorm:"primaryKey;type:varchar(20)"`
	UserId        string `json:"user_id"`
	ApplicationId string `json:"application_id"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
	IsActive      bool              `gorm:"type:bool;default:true" json:"status"`
	Roles         []RoleApplication `gorm:"many2many:account_role_applications;"`
}

func (Account) TableName() string {
	return "accounts"
}
