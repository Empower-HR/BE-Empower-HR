package migrations

import (
	datacompanies "be-empower-hr/features/Companies/data_companies"
	dataleaves "be-empower-hr/features/Leaves/data_leaves"
	datapayroll "be-empower-hr/features/Payroll/data_payroll"
	dataschedule "be-empower-hr/features/Schedule/data_schedule"
	datausers "be-empower-hr/features/Users/data_users"
	dataattendance "be-empower-hr/features/Attendance/data_attendance"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&datausers.EmploymentData{})
	db.AutoMigrate(&datausers.PersonalData{})
	db.AutoMigrate(&datacompanies.CompanyData{})
	db.AutoMigrate(&dataleaves.LeavesData{})
	db.AutoMigrate(&datapayroll.PayrollData{})
	db.AutoMigrate(&dataschedule.ScheduleData{})
	db.AutoMigrate(&dataattendance.Attandance{})
}
