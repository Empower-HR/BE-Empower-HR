package attendance

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Attandance struct {
	ID             uint
 	PersonalDataID 	uint
	Clock_in       	string
	Clock_out	    string
	Status	        string
	Date		   	time.Time
	Long       		string
	Lat				string
	Notes			string
}
type AttendanceDetail struct {
	ID			 uint
    Name         string 
	PersonalDataID 	uint
    ScheduleIn   string
    ScheduleOut  string
    ClockIn      string
	ClockOut     string 
	Status	        string
	Date		   	string
	Long       		string
	Lat				string
	Notes			string
}
type AttendanceEntity struct {
	ID			 uint
    Name         string 
	PersonalDataID 	uint
    ScheduleIn   string
    ScheduleOut  string
    ClockIn      string
	ClockOut     string 
	Status	        string
	Date		   	time.Time
	Long       		string
	Lat				string
	Notes			string
}

type CompanyDataEntity struct {
	ID	uint
	CompanyAddress string
}

type AHandler interface {
	AddAttendance(c echo.Context) error 
	UpdateAttendance(c echo.Context) error
	DeleteAttendance(c echo.Context) error
	GetAttendancesHandler(c echo.Context) error
	GetAllAttendancesHandler(c echo.Context) error
	DownloadPdf(c echo.Context) error
	GetAttendancesbyID(c echo.Context) error
}

type AServices interface {
	AddAtt(newAtt Attandance) error
	UpdateAtt(id uint, updateAtt Attandance) error
	DeleteAttByID(attID uint) error
	GetAttByPersonalID(personalID uint,searchBox string, limit int, offset int) ([]AttendanceDetail, error)
	GetAllAtt(search string, limit int, offset int) ([]AttendanceDetail, error)
	CountAllAtt() (int64, error)
	CountAllAttbyPerson(personID uint) (int64, error)
	CountAllAttbyDate(date string) (int64, error)
	CountAllAttbyDateandPerson(date string, personID uint) (int64, error)
	DownloadAllAtt() error
	GetAllAttbyDate(date string, limit int, offset int) ([]AttendanceDetail, error)
	GetAllAttbyStatus(status string, limit int, offset int) ([]AttendanceDetail, error)
	GetAttByIdAtt(idAtt uint) ([]AttendanceDetail, error)
	GetAttByPersonalIDandStatus(id uint, status string, limit int, offset int) ([]AttendanceDetail, error)
	CountAllAttbyStatus(status string) (int64, error)
	CountAllAttbyStatusandPerson(status string, personID uint) (int64, error)
	GetAllAttbyDateandPerson(date string,limit int, offset int, personId uint) ([]AttendanceDetail, error)
}

type AQuery interface {
	Create(newAtt Attandance) error
	IsDateExists(personalID uint, date time.Time) (bool, error)
	Update(id uint, updatedAtt Attandance) error
	GetAttByPersonalID(personalID uint, term string, limit int, offset int) ([]AttendanceDetail, error)
	DeleteAttbyId(attId uint) error
	GetTotalAttendancesCount() (int64, error)
	GetTotalAttendancesCountbyPerson(personID uint) (int64, error)
	GetTotalAttendancesCountbyDate(date time.Time) (int64, error)
	GetTotalAttendancesCountbyDateandPerson(date time.Time, personID uint) (int64, error)
	GetTotalAttendancesCountByStatus(status string) (int64, error)
	GetTotalAttendancesCountByStatusandPerson(status string, personID uint) (int64, error)
	GetAllAttDownload() ([]Attandance, error)
	GetAllAttbyDate(date time.Time, limit int, offset int) ([]AttendanceDetail, error)
	GetAllAttbyDateandPerson(perseonID uint, date time.Time, limit int, offset int) ([]AttendanceDetail, error)
	GetAttendanceDetails(searchTerm string, limit int, offset int) ([]AttendanceDetail, error)
	GetAttByIdAtt(idAtt uint) ([]AttendanceDetail, error)
	GetAllAttbyStatus(status string, limit int, offset int) ([]AttendanceDetail, error)
	GetAllAttbyIdPersonAndStatus(id uint, status string, limit int, offset int) ([]AttendanceDetail, error)
	GetCompany(idPerson uint) ([]CompanyDataEntity, error)
}
