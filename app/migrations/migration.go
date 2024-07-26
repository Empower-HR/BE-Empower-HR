package migrations

import (
	datacompanies "be-empower-hr/features/Companies/data_companies"
	dataleaves "be-empower-hr/features/Leaves/data_leaves"
	datapayroll "be-empower-hr/features/Payroll/data_payroll"
	dataschedule "be-empower-hr/features/Schedule/data_schedule"
	datausers "be-empower-hr/features/Users/data_users"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&datausers.EmploymentData{})
	db.AutoMigrate(&datausers.PersonalData{})
	db.AutoMigrate(&datacompanies.CompanyData{})
	db.AutoMigrate(&dataleaves.LeavesData{})
	db.AutoMigrate(&datapayroll.PayrollData{})
	db.AutoMigrate(&dataschedule.ScheduleData{})
}
