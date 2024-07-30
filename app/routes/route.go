package routes

import (
	"be-empower-hr/app/middlewares"
	_datacompanies "be-empower-hr/features/Companies/data_companies"
	_companyHandler "be-empower-hr/features/Companies/handler"
	_companyService "be-empower-hr/features/Companies/service"
	_userData "be-empower-hr/features/Users/data_users"
	_userHandler "be-empower-hr/features/Users/handler"
	_userService "be-empower-hr/features/Users/service"

	_attData "be-empower-hr/features/Attendance/data_attendance"
	_attHandler "be-empower-hr/features/Attendance/handler"
	_attService "be-empower-hr/features/Attendance/service"

	// _scheduleData "be-empower-hr/features/Schedule/data_schedule"
	// _schduleService "be-empower-hr/features/Schedule/service"
	// _scheduleHandler "be-empower-hr/features/Schedule/handler"
	_scheduleData "be-empower-hr/features/Schedule/data_schedule"
	_scheduleHandler "be-empower-hr/features/Schedule/handler"
	_schduleService "be-empower-hr/features/Schedule/service"

	_leavesData "be-empower-hr/features/Leaves/data_leaves"
	_leavesHandler "be-empower-hr/features/Leaves/handler"
	_leavesDataService "be-empower-hr/features/Leaves/service"
  
	_annoData "be-empower-hr/features/Announcement/data_announcement"
	_annoService "be-empower-hr/features/Announcement/service"
	_annoHandler "be-empower-hr/features/Announcement/handler"

	"be-empower-hr/utils"
	"be-empower-hr/utils/cloudinary"
	"be-empower-hr/utils/encrypts"
	"be-empower-hr/utils/pdf"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {
	middlewares := middlewares.NewMiddlewares()
	hashService := encrypts.NewHashService()
	accountUtility := utils.NewAccountUtility()
	pdfUtility := pdf.NewPdfUtility()
	cloudinaryUtility := cloudinary.NewCloudinaryUtility()
	company := _datacompanies.NewCompanyModels(db)
	userData := _userData.New(db)
	userService := _userService.New(userData, hashService, middlewares, accountUtility, company)
	attData := _attData.NewAttandancesModel(db)
	attService := _attService.New(attData, hashService, middlewares, accountUtility, pdfUtility)
	attHandler := _attHandler.New(attService)
	
	annoData := _annoData.NewModelAnnouncement(db)
	annoService := _annoService.New(annoData, hashService, middlewares, accountUtility, cloudinaryUtility)
	annoHandler := _annoHandler.New(annoService)

	scheduleData := _scheduleData.New(db)
	scheduleService := _schduleService.New(scheduleData, accountUtility)
	scheduleHandlerAPI := _scheduleHandler.New(scheduleService)
	userHandlerAPI := _userHandler.New(userService, cloudinaryUtility)

	leavesData := _leavesData.New(db)
	leavesDataService := _leavesDataService.New(leavesData)
	leavesHandlerAPI := _leavesHandler.New(leavesDataService)

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
	e.PUT("/employment", userHandlerAPI.UpdateProfileEmployment, middlewares.JWTMiddleware())
	e.PUT("/employment/:id", userHandlerAPI.UpdateEmploymentEmployee, middlewares.JWTMiddleware())
	e.GET("/dashboard", userHandlerAPI.DasboardAdmin, middlewares.JWTMiddleware())

	//handler memployees
	e.POST("/employee", userHandlerAPI.CreateNewEmployee, middlewares.JWTMiddleware())
	e.GET("/employee", userHandlerAPI.GetAllAccount, middlewares.JWTMiddleware())
	e.GET("/employee/:id", userHandlerAPI.GetProfileById, middlewares.JWTMiddleware())
	e.PUT("/employee", userHandlerAPI.UpdateProfileEmployees, middlewares.JWTMiddleware())
	e.DELETE("/employee/:id", userHandlerAPI.DeleteAccountEmployees, middlewares.JWTMiddleware())
	e.GET("/dashboard/employee", userHandlerAPI.DashboardEmployees, middlewares.JWTMiddleware())

	e.POST("/attendance", attHandler.AddAttendance, middlewares.JWTMiddleware())
	e.PUT("/attendance/:attendance_id", attHandler.UpdateAttendance, middlewares.JWTMiddleware())
	e.DELETE("/attendance/:attendance_id", attHandler.DeleteAttendance, middlewares.JWTMiddleware())
	e.GET("/attendance", attHandler.GetAllAttendancesHandler, middlewares.JWTMiddleware())
  
	e.GET("/attendance/:attendance_id", attHandler.GetAttendancesHandler, middlewares.JWTMiddleware())
	e.GET("/attendance/download", attHandler.DownloadPdf)

	e.GET("/attendance/download", attHandler.DownloadPdf, middlewares.JWTMiddleware())
	e.GET("/attendance/user/:employee_id", attHandler.GetAttendancesHandler, middlewares.JWTMiddleware())
	e.GET("/attendance/:attendance_id", attHandler.GetAttendancesbyID, middlewares.JWTMiddleware())

	// handler company
	e.PUT("/companies", ch.UpdateCompany(), middlewares.JWTMiddleware())
	e.GET("/companies", ch.GetCompany(), middlewares.JWTMiddleware())

	// handler scheduled
	e.POST("/schedule", scheduleHandlerAPI.CreateSchedule(), middlewares.JWTMiddleware())
	e.GET("/schedule", scheduleHandlerAPI.GetAllSchedule, middlewares.JWTMiddleware())
	e.GET("/schedule/:id", scheduleHandlerAPI.GetScheduleById, middlewares.JWTMiddleware())
	e.PUT("/schedule/:id", scheduleHandlerAPI.UpdateSchedule, middlewares.JWTMiddleware())
	e.DELETE("/schedule/:id", scheduleHandlerAPI.DeleteSchedule, middlewares.JWTMiddleware())

	//handler leaves
	e.POST("/leaves", leavesHandlerAPI.RequestLeave, middlewares.JWTMiddleware())
	e.PUT("/leaves/:id", leavesHandlerAPI.UpdateLeaveStatus, middlewares.JWTMiddleware())
	e.GET("/leaves", leavesHandlerAPI.ViewLeaveHistory, middlewares.JWTMiddleware())
	e.GET("/leaves/:id", leavesHandlerAPI.GetLeavesByID, middlewares.JWTMiddleware())
  
	// handler announcement
	e.POST("/announcement", annoHandler.AddAnnouncement, middlewares.JWTMiddleware())
	e.GET("/announcement", annoHandler.GetAnno, middlewares.JWTMiddleware())
}
