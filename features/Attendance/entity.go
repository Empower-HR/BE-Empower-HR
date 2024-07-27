package attendance

import (
	"github.com/labstack/echo/v4"
)

type Attandance struct {
	ID             uint
 	PersonalDataID 	uint
	Clock_in       	string
	Clock_out	    string
	Status	        string
	Date		   	string
	Long       		string
	Lat				string
	Notes			string
}

type AHandler interface {
	AddAttendance(c echo.Context) error 
	UpdateAttendance(c echo.Context) error
	DeleteAttendance(c echo.Context) error
	GetAttendancesHandler(c echo.Context) error
	GetAllAttendancesHandler(c echo.Context) error
}

type AServices interface {
	AddAtt(newAtt Attandance) error
	UpdateAtt(id uint, updateAtt Attandance) error
	DeleteAttByID(attID uint) error
	GetAttByPersonalID(personalID uint) ([]Attandance, error)
	GetAllAtt() ([]Attandance, error)
}

type AQuery interface {
	Create(newAtt Attandance) error
	IsDateExists(personalID uint, date string) (bool, error)
	Update(id uint, updatedAtt Attandance) error
	GetAttByPersonalID(personalID uint) ([]Attandance, error)
	GetAllAtt() ([]Attandance, error)
	DeleteAttbyId(attId uint) error
}