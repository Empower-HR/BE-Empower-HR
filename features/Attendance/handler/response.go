package handler

import(
	"be-empower-hr/features/Attendance"
)

type AttAllResponse struct {
	ID       uint   `json:"id"`
	PersonalDataID uint `json:"personal_id"`
	ClockIn  string `json:"clock_in"`
	ClockOut string `json:"clock_out"`
	Status   string `json:"status"`
	Date     string `json:"date"`
	Long     string `json:"long"`
	Lat      string `json:"lat"`
	Notes    string `json:"notes"`
}
type AttResponse struct {
	ID       uint   `json:"id"`
	ClockIn  string `json:"clock_in"`
	ClockOut string `json:"clock_out"`
	Status   string `json:"status"`
	Date     string `json:"date"`
	Long     string `json:"long"`
	Lat      string `json:"lat"`
	Notes    string `json:"notes"`
}

func ToGetAttendanceResponse(attendance attendance.Attandance) AttResponse {
	return AttResponse{
		ID:       attendance.ID,
		ClockIn:  attendance.Clock_in,
		ClockOut: attendance.Clock_out,
		Status:   attendance.Status,
		Date:     attendance.Date,
		Long:     attendance.Long,
		Lat:      attendance.Lat,
		Notes:    attendance.Notes,
	}
}
func ToGetAllAttendance(attendance attendance.Attandance) AttAllResponse {
	return AttAllResponse{
		ID:       attendance.ID,
		PersonalDataID: attendance.PersonalDataID,
		ClockIn:  attendance.Clock_in,
		ClockOut: attendance.Clock_out,
		Status:   attendance.Status,
		Date:     attendance.Date,
		Long:     attendance.Long,
		Lat:      attendance.Lat,
		Notes:    attendance.Notes,
	}
}