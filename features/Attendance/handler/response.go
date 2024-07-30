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
	EmploymentData []EmployeDataResponse
}

type EmployeDataResponse struct {
	ID			 uint	   `json:"id_personal"`
	Name         string    `json:"name"`
}
type AttDetailResponse struct {
	ID			 uint	`json:"id_att"`
	Name         string    `json:"name"`
	PersonalDataID 	uint	`json:"personal_id"`
    ClockIn      string `json:"clock_in"`
    ClockOut     string `json:"clock_out"`
	Date		 string `json: "date"`
}


func ToGetAttendanceResponse(attendance attendance.AttendanceDetail) AttResponse {
	employmentData := FetchEmploymentData(attendance.PersonalDataID, attendance.Name)
	return AttResponse{
		ID:       attendance.ID,
		ClockIn:  attendance.ClockIn,
		ClockOut: attendance.ClockOut,
		Status:   attendance.Status,
		Date:     attendance.Date,
		Long:     attendance.Long,
		Lat:      attendance.Lat,
		Notes:    attendance.Notes,
		EmploymentData: employmentData,
	}

}

func FetchEmploymentData(id uint, name string) []EmployeDataResponse {
    // Implementasi logika untuk mengambil data karyawan berdasarkan ID
    // Misalnya query database atau API call
    return []EmployeDataResponse{
        // Contoh data dummy
        {ID: id, Name: name},
    }
}

func ToGetAttendanceDetailResponse(attendance attendance.AttendanceDetail) AttDetailResponse {
	return AttDetailResponse{
		ID: attendance.ID,
		Name : attendance.Name,
		PersonalDataID: attendance.PersonalDataID,
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