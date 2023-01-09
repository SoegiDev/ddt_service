package schema

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id           string `json:"id" gorm:"primaryKey;size:50;"`
	Email        string `gorm:"size:255;not null;unique" json:"email"`
	Username     string `gorm:"size:255;not null;unique" json:"username"`
	Password     string `gorm:"size:255;not null;" json:"-"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	IsDeleted    bool          `gorm:"type:bool;default:false" json:"delete"`
	IsActive     bool          `gorm:"type:bool;default:true" json:"status"`
	Employees    []Employee    `gorm:"foreignKey:UserId"`
	Divisions    []Division    `gorm:"many2many:user_divisions;"`
	Roles        []Role        `gorm:"many2many:user_roles;"`
	Articles     []Article     `gorm:"foreignKey:UserId"`
	Accounts     []Account     `gorm:"foreignKey:UserId"`
	ActivityLogs []ActivityLog `gorm:"foreignKey:UserId"`
}

func (User) TableName() string {
	return "users"
}

type AssignRole struct {
	Role []string `json:"role"`
}

type AddDivision struct {
	Division []string `json:"division"`
}

type UpdateDivision struct {
	Divisions []Division `gorm:"many2many:user_divisions;"`
}
