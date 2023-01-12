package model

import (
	"ddtservice_agri/database"
	"ddtservice_agri/schema"
	"fmt"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User schema.User

func (data *User) Save() (*User, error) {
	err := database.Database.Create(&data).Error
	return data, err
}

func (data *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	data.Password = string(passwordHash)
	data.Username = html.EscapeString(strings.TrimSpace(data.Username))
	return nil
}

func (data *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(password))
}

func UserFindByUsername(username string) (User, error) {
	var data User
	err := database.Database.Preload("Roles").Preload("Divisions").Preload("ActivityLogs").Preload("Accounts.Application").Preload("Accounts.RoleApplications").Preload("Employees.Company").Where("username = ?", username).Preload("Employees", "is_active NOT IN (?)", false).First(&data).Error
	return data, err
}

func UserFindById(id string) (User, error) {
	var data User
	err := database.Database.Preload("Divisions").Preload("ActivityLogs").Preload("Accounts.Application").Preload("Accounts.RoleApplications").Preload("Employees.Company").Where("id = ?", id).Preload("Employees", "is_active NOT IN (?)", false).First(&data).Error
	return data, err
}

func UserFindByCode(code string) (User, error) {
	var data User
	err := database.Database.Preload("Divisions").Preload("ActivityLogs").Preload("Accounts.Application").Preload("Accounts.RoleApplications").Preload("Employees.Company").Where("code = ?", code).Preload("Employees", "is_active NOT IN (?)", false).First(&data).Error
	return data, err
}

func UserFindByIdLogin(id string) (User, error) {
	var data User
	err := database.Database.Preload("ActivityLogs").Preload("Accounts.Application").Preload("Accounts.RoleApplications").Preload("Employees.Company").Where("id = ?", id).Preload("Employees", "is_active NOT IN (?)", false).First(&data).Error
	return data, err
}

func UserMetaFindById(id string) (User, error) {
	var data User
	err := database.Database.Where("id = ?", id).Preload("Accounts.RoleApplications").Preload("Accounts.Application").Preload("Employees.Company").Preload("Employees", "is_active NOT IN (?)", false).Preload(clause.Associations).First(&data).Error
	return data, err
}

func (data *User) UserAssignRoles(id string, roleUpdates []schema.Role) (User, error) {
	err := database.Database.Model(&data).Association("Roles").Replace(roleUpdates)
	if err != nil {
		return *data, err
	}
	res, _ := UserFindByIdLogin(id)
	return res, nil
}
func (data *User) UserAssignDivision(id string, divisionUpdate []schema.Division) (User, error) {
	err := database.Database.Model(&data).Association("Divisions").Replace(divisionUpdate)
	if err != nil {
		return *data, err
	}
	res, _ := UserFindByIdLogin(id)
	return res, nil
}
func (data *User) UserGetCount() int64 {
	var result int64
	database.Database.Model(&data).Count(&result)
	return result
}

func (data *User) IsAdmin() bool {
	fmt.Println(data.Roles)
	result := database.Database.Where("name IN (?)", "Subscriber").Find(&data.Roles)
	fmt.Println(result.RowsAffected)
	return false
}

func (data *User) CheckADMIN() bool {
	for _, n := range data.Roles {
		if "Subscriber" == n.Name {
			return true
		}
	}
	return false
}
func (data *User) IsSubscribe() bool {
	var count int64
	database.Database.Debug().Preload("Roles", "name IN (?)", "Administrator").Count(&count).First(&data)
	if count > 0 {
		fmt.Println("TOTAL", count)
		return true
	}
	return false
}
