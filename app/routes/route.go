package routes

import (
	"be-empower-hr/app/middlewares"
	_datacompanies "be-empower-hr/features/Companies/data_companies"
	_companyHandler "be-empower-hr/features/Companies/handler"
	_companyService "be-empower-hr/features/Companies/service"
	_userData "be-empower-hr/features/Users/data_users"
	_userHandler "be-empower-hr/features/Users/handler"
	_userService "be-empower-hr/features/Users/service"

	// _scheduleData "be-empower-hr/features/Schedule/data_schedule"
	// _schduleService "be-empower-hr/features/Schedule/service"
	// _scheduleHandler "be-empower-hr/features/Schedule/handler"

	"be-empower-hr/utils"
	"be-empower-hr/utils/cloudinary"
	"be-empower-hr/utils/encrypts"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {
	middlewares := middlewares.NewMiddlewares()
	hashService := encrypts.NewHashService()
	accountUtility := utils.NewAccountUtility()

	userData := _userData.New(db)
	userService := _userService.New(userData, hashService, middlewares, accountUtility)
	userHandlerAPI := _userHandler.New(userService)

	// scheduleData := _scheduleData.New(db)
	// scheduleService := _schduleService.New(scheduleData, accountUtility)
	// scheduleHandlerAPI := _scheduleHandler.New(scheduleService)

	// api company
	cm := _datacompanies.NewCompanyModels(db)
	cs := _companyService.NewCompanyServices(cm)
	cl := cloudinary.NewCloudinaryUtility()
	ch := _companyHandler.NewCompanyHandler(cs, cl)

	//handler admin
	e.POST("/admin", userHandlerAPI.RegisterAdmin)
	e.POST("/login", userHandlerAPI.Login)
	e.GET("/admin", userHandlerAPI.GetProfile, middlewares.JWTMiddleware())
	e.DELETE("/admin", userHandlerAPI.DeleteAccountAdmin, middlewares.JWTMiddleware())
	e.PUT("/admin", userHandlerAPI.UpdateProfileAdmins, middlewares.JWTMiddleware())

	// handler company
	e.POST("/companies", ch.UpdateCompany())
	e.GET("/companies", ch.GetCompany())
}
