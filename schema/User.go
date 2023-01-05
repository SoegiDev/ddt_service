package schema

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        string `json:"id" gorm:"primaryKey;type:varchar(100)"`
	Email     string `gorm:"size:255;not null;unique" json:"email"`
	Username  string `gorm:"size:255;not null;unique" json:"username"`
	Password  string `gorm:"size:255;not null;" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Deleted   bool       `gorm:"type:bool" json:"deleted"`
	IsActive  bool       `gorm:"type:bool;default:true" json:"status"`
	Employees []Employee `gorm:"foreignKey:UserId"`
	Divisions []Division `gorm:"many2many:user_divisions;"`
	Roles     []Role     `gorm:"many2many:user_roles;"`
	Articles  []Articles `gorm:"foreignKey:UserId"`
}

func (User) TableName() string {
	return "users"
}

type AddDivision struct {
	Division []string `json:"division"`
}

type UpdateDivision struct {
	Divisions []Division `gorm:"many2many:user_divisions;"`
}
