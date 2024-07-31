package dataattendance

import (
	attendance "be-empower-hr/features/Attendance"

	"gorm.io/gorm"
)

type Attandance struct {
	gorm.Model
	PersonalDataID   uint					  `gorm:"foreignKey:PersonalDataID"`
	Clock_in       string   		          `json:"clock_in"`
	Clock_out      string                  `json:"clock_out"`
	Status         string                     `json:"status"`
	Date    	   string                  `json:"date"`
	Long       	   string                     `json:"long"`
	Lat	   		   string					  `json:"lat"`
	Notes		   string					  `json:"notes"`
}

type PersonalData struct {
	ID   uint
	Name string
	CompanyID   uint
}

type ScheduleData struct {
	ID   uint
	ScheduleIn    string
	ScheduleOut   string
	CompanyID     uint
}
type CompanyData struct {
	ID   uint
	CompanyAddress string
}

func AttandanceInput(input attendance.Attandance) Attandance{
	return Attandance{
	 	PersonalDataID:    input.PersonalDataID,
		Clock_in:	input.Clock_in,
		Clock_out: 	input.Clock_out,
		Status:    	input.Status,
		Date:		input.Date,
		Long:		input.Long,
		Lat:		input.Lat,
		Notes:		input.Notes,
	}
}

func (at *Attandance) ToAttEntity(pd *PersonalData, sc *ScheduleData) attendance.AttendanceDetail{
	return attendance.AttendanceDetail{
		Name : pd.Name,
		PersonalDataID: at.PersonalDataID,
		ScheduleIn   : sc.ScheduleIn,
		ScheduleOut  : sc.ScheduleOut,
		ClockIn     : at.Clock_in,
		ClockOut  : at.Clock_out,
		Date	: at.Date,
	}
}

func (at *Attandance) ToCompanyEntity(cm *CompanyData) attendance.CompanyDataEntity {
	return attendance.CompanyDataEntity{
		ID : cm.ID,
		CompanyAddress: cm.CompanyAddress,
	}
}