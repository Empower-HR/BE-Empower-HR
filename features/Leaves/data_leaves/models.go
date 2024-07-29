package dataleaves

import (
	"time"

	"gorm.io/gorm"
)

type LeavesData struct {
	gorm.Model
	StartDate      time.Time
	EndDate        time.Time
	Reason         string
	Status         string
	TotalLeave     int
	PersonalDataID uint
}

type PersonalData struct {
	gorm.Model
	CompanyID      uint
	ProfilePicture string
	Name           string
	Email          string
	Password       string
	PhoneNumber    string
	PlaceBirth     string
	BirthDate      string
	Gender         string
	Religion       string
	NIK            string
	Address        string
	Role           string
}

type EmploymentData struct {
	gorm.Model
	PersonalDataID   uint
	EmploymentStatus string
	JoinDate         string
	Department       string
	JobPosition      string
	JobLevel         string
	Schedule         string
	ApprovalLine     string
	Manager          string
}
