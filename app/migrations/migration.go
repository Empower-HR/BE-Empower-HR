package migrations

import (
	dataannouncement "be-empower-hr/features/Announcement/data_announcement"
	dataattendance "be-empower-hr/features/Attendance/data_attendance"
	datacompanies "be-empower-hr/features/Companies/data_companies"
	dataleaves "be-empower-hr/features/Leaves/data_leaves"
	dataschedule "be-empower-hr/features/Schedule/data_schedule"
	datausers "be-empower-hr/features/Users/data_users"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&datausers.EmploymentData{})
	db.AutoMigrate(&datausers.PersonalData{})
	db.AutoMigrate(&datacompanies.CompanyData{})
	db.AutoMigrate(&dataleaves.LeavesData{})
	db.AutoMigrate(&datausers.PayrollData{})
	db.AutoMigrate(&dataschedule.ScheduleData{})
	db.AutoMigrate(&dataattendance.Attandance{})
	db.AutoMigrate(&dataannouncement.Announcement{})
}
