package protect

import (
	"ddtservice_agri/controller"
	"ddtservice_agri/middleware"

	"github.com/gin-gonic/gin"
)

// @Host localhost:8080

// @Title Swagger User Service
// @Version 1.0
// @Description This is User Service Test.
// @termsOfService http://swagger.io/terms/

// @contact.name Fajar soegi
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @BasePath /api
func Protected(router *gin.Engine) {
	protectedRoutes := router.Group("/api")

	protectedRoutes.Use(middleware.JWTAuthMiddleware())

	//User
	protectedRoutes.POST("/user_new", controller.UserAddNew)
	protectedRoutes.GET("/user_profile/:ID", controller.User_Profile)

	//Employee
	protectedRoutes.POST("/employee", controller.EmployeeAddNew)
	protectedRoutes.PATCH("/employee/:ID/edit", controller.EmployeeUpdate)
	protectedRoutes.GET("/employee/:ID", controller.EmployeeMeta)

	// Company
	protectedRoutes.POST("/company", controller.CompanyAddNew)
	protectedRoutes.PATCH("/company/:ID/edit", controller.CompanyUpdate)
	protectedRoutes.GET("/company/:ID", controller.CompanyById)
	protectedRoutes.GET("/company", controller.CompanyByAll)

	// Estate
	protectedRoutes.POST("/estate", controller.EstateAddNew)
	protectedRoutes.PATCH("/estate/:ID/edit", controller.EstateUpdate)
	protectedRoutes.GET("/estate/:ID", controller.EstateById)
	protectedRoutes.GET("/estate", controller.EstateByAll)

	// Division
	protectedRoutes.POST("/division", controller.DivisionAddNew)
	protectedRoutes.PATCH("/division/:ID/edit", controller.DivisionUpdate)
	protectedRoutes.GET("/division/:ID", controller.DivisionById)
	protectedRoutes.GET("/division", controller.DivisionByAll)

	// Gang
	protectedRoutes.POST("/gang", controller.GangAddNew)
	protectedRoutes.PATCH("/gang/:ID/edit", controller.GangUpdate)
	protectedRoutes.GET("/gang/:ID", controller.GangById)
	protectedRoutes.GET("/gang", controller.GangByAll)

	// Article
	protectedRoutes.POST("/article", controller.ArticleAddNew)
	protectedRoutes.GET("/article/:ID", controller.ArticleFindById)
	protectedRoutes.GET("/article", controller.AllArticles)
	protectedRoutes.PATCH("/article/:ID/edit", controller.ArticleUpdate)

	// Application
	protectedRoutes.POST("/application", controller.ApplicationAddNew)
	protectedRoutes.PATCH("/application/:ID/edit", controller.ApplicationUpdate)
	protectedRoutes.GET("/application/:ID", controller.ApplicationById)
	protectedRoutes.GET("/application", controller.ApplicationByAll)
}
