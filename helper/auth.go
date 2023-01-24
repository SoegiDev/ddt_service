package helper

import (
	"ddtservice_agri/model"
	"ddtservice_agri/schema"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Auth(context *gin.Context, input schema.SignInLDAPJsonSchema, data_ldap schema.LDAPLOGIN) (schema.LoginResponse, string) {
	user, err := model.UserFindByUsername(input.Username)
	var result schema.LoginResponse
	if err != nil {
		if data_ldap.AccountName != "" {
			input_roles := []string{"User"}
			data_roles, err := model.FindRoleMapById(input_roles)
			if err != nil {
				return result, err.Error()
			}
			userID, _ := GenerateUserId(3)
			profileId, _ := GenerateProfileId(3)

			roleMap := []schema.Role{}
			for _, element := range data_roles {
				roleMap = append(roleMap, schema.Role{Id: element.Id})
			}

			user := model.User{
				Id:       userID,
				Code:     data_ldap.EmployeeNumber,
				Roles:    roleMap,
				Email:    data_ldap.Email,
				Username: input.Username,
				Password: input.Password,
			}

			savedUser, err := user.Save()

			if err != nil {
				return result, err.Error()
			}

			emp := model.Employee{
				Id:          profileId,
				Code:        data_ldap.EmployeeNumber,
				Username:    input.Username,
				Email:       data_ldap.Email,
				Nik:         data_ldap.EmployeeNumber,
				NickName:    data_ldap.Surname,
				FullName:    data_ldap.Fullname,
				PhoneNumber: data_ldap.Phone,
				Department:  data_ldap.Department,
				UserCode:    &savedUser.Code}

			savedProfile, err := emp.Save()
			if err != nil {
				return result, err.Error()
			}
			fmt.Println(savedProfile)
			fmt.Println(savedUser)
			jwt, err := GenerateJWT(user)
			if err != nil {
				return result, err.Error()
			}
			result.Token = jwt
			return result, ""
		} else {
			return result, err.Error()
		}

	}

	err = user.ValidatePassword(input.Password)
	if err != nil {
		return result, err.Error()
	}

	jwt, err := GenerateJWT(user)
	if err != nil {
		return result, err.Error()
	}
	result.Token = jwt
	return result, ""
}
