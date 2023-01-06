package schema

import (
	"time"

	"gorm.io/gorm"
)

type Application struct {
	Id              string `json:"id" gorm:"primaryKey;type:varchar(100)"`
	Name            string `gorm:"size:100;" json:"name"`
	Description     string `gorm:"size:255;" json:"description"`
	UpdatedNote     string `gorm:"type:text;" json:"updated_note"`
	Version         string `gorm:"size:50;" json:"version"`
	AppPackage      string `gorm:"size:100;" json:"app_package"`
	AppPackageClass string `gorm:"size:100;" json:"app_package_class"`
	AssetApk        string `gorm:"size:100;" json:"asset_apk"`
	AssetIcon       string `gorm:"size:100;" json:"asset_icon"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
	Deleted         bool      `gorm:"type:bool;default:false" json:"deleted"`
	Accounts        []Account `gorm:"foreignKey:ApplicationId"`
}

func (Application) TableName() string {
	return "applications"
}