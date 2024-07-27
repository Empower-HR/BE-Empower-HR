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