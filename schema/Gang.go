package schema

import (
	"gorm.io/gorm"
)

type Gang struct {
	Id         uint   `json:"id" gorm:"primaryKey;AUTO_INCREMENT;not_null"`
	Name       string `gorm:"size:50;" json:"name"`
	DeletedAt  gorm.DeletedAt
	IsDeleted  bool   `gorm:"type:bool;default:true" json:"delete"`
	IsActive   bool   `gorm:"type:bool;default:true" json:"status"`
	DivisionId string `json:"division_id"`
	Division   Division
}

func (Gang) TableName() string {
	return "gangs"
}
