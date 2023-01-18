package public

import (
	"ddtservice_agri/helper"
	"ddtservice_agri/model"
	"ddtservice_agri/schema"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Sign In with LDAP or without LDAP
// @Description Sign In with LDAP or without LDAP
// @Accept  json
// @Param login body schema.SignInLDAPJsonSchema true "LDAP AUTH Schema "
// @Produce  json
// @Success 200
// @Router /sign_in [post]
func SignIn_ldap(context *gin.Context) {
	var input schema.SignInLDAPJsonSchema
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, schema.MsgResponse{Msg: err.Error()})
		return
	}
	if input.Ldap {
		ldap, err := helper.AuthLDAP(context, input)
		if err != "" {
			context.JSON(http.StatusBadRequest, schema.MsgResponse{Msg: err})
			return
		}
		fmt.Println(ldap)
		jwt, err := helper.Auth(context, input, ldap)
		if err != "" {
			context.JSON(http.StatusBadRequest, schema.MsgResponse{Msg: err})
			return
		}
		fmt.Println(jwt)
		context.JSON(http.StatusOK, jwt)
	} else {
		var ldap schema.LDAPLOGIN
		jwt, err := helper.Auth(context, input, ldap)
		if err != "" {
			context.JSON(http.StatusBadRequest, schema.MsgResponse{Msg: err})
			return
		}
		context.JSON(http.StatusOK, jwt)
	}

}

//
// @Summary Sign Up
// @Description Sign Up User
// @Accept  json
// @Param login body schema.SignUpJsonSchema true "Login Schema "
// @Produce  json
// @Success 200 {object}  schema.MsgResponse
// @Router /sign_up [post]
func SignUp(context *gin.Context) {
	var input schema.SignUpJsonSchema
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input_roles := input.Role
	data_roles, err := model.FindRoleMapById(input_roles)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, _ := helper.GenerateUserId(3)
	userCode, _ := helper.GenerateCodeRandom()

	roleMap := []schema.Role{}
	for _, element := range data_roles {
		roleMap = append(roleMap, schema.Role{Id: element.Id})
	}

	user := model.User{
		Id:       userID,
		Code:     userCode,
		Roles:    roleMap,
		Email:    input.Email,
		Username: input.Username,
		Password: input.Password,
	}

	savedUser, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(savedUser)
	context.JSON(http.StatusCreated, schema.MsgResponse{Msg: "Register Completed"})
}
