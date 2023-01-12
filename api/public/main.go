package public

import (
	"github.com/gin-gonic/gin"
)

// @Host localhost:8080
// @title Swagger Example API
// @version 2.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func Authentication(router *gin.Engine) {
	auth := router.Group("auth")

	auth.POST("/register", Register)
	// @BasePath /auth
	auth.POST("/login", Login)
	// //Corporation
	// auth.POST("/company", controller.CompanyAddNew)
	// auth.PATCH("/company/:ID/edit", controller.CompanyUpdate)

	// //Application
	// auth.POST("/application", controller.ApplicationAddNew)
	// auth.PATCH("/application/:ID/edit", controller.ApplicationUpdate)
}
