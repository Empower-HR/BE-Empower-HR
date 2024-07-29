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
type AttDetailResponse struct {
	ID			 uint	`json:"id_att"`
	Name         string    `json:"name"`
	PersonalDataID 	uint	`json:"personal_id"`
    ScheduleIn   string `json:"schedule_in"`
    ScheduleOut  string `json:"schedule_out"`
    ClockIn      string `json:"clock_in"`
    ClockOut     string `json:"clock_out"`
	Date		 string `json: "date"`
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
func ToGetAttendanceDetailResponse(attendance attendance.AttendanceDetail) AttDetailResponse {
	return AttDetailResponse{
		ID: attendance.ID,
		Name : attendance.Name,
		PersonalDataID: attendance.PersonalDataID,
		ScheduleIn: attendance.ScheduleIn,
		ScheduleOut: attendance.ScheduleOut,
		ClockIn:  attendance.ClockIn,
		ClockOut: attendance.ClockOut,
		Date:     attendance.Date,
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