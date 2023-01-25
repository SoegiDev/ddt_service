package public

import (
	"github.com/gin-gonic/gin"
)

// @Host user-test2.azurewebsites.net
// @title Authentication User Service
// @version 1.0.0
// @description User Service OPEN API.
// @termsOfService http://swagger.io/terms/

// @contact.name ICT INDOAGRI
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func Authentication(router *gin.Engine) {
	auth := router.Group("auth")
	// @BasePath /auth
	auth.POST("/sign_up", SignUp)
	auth.POST("/sign_in", SignIn_ldap)
}
