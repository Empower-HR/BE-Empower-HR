package leaves

import "time"

type LeavesDataEntity struct {
	LeavesID       uint
	StartDate      time.Time
	EndDate        time.Time
	Reason         string
	Status         string
	TotalLeave     int
	PersonalDataID uint
}

type PersonalDataEntity struct {
	PersonalDataID uint
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

type EmploymentDataEntity struct {
	EmploymentDataID uint
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

type DataLeavesInterface interface {
	RequestLeave(leave LeavesDataEntity) error
	UpdateLeaveStatus(leaveID uint, status string) error
	GetLeaveHistory(personalDataID uint, page, pageSize int) ([]LeavesDataEntity, error)
	GetLeavesByStatus(personalDataID uint, status string) ([]LeavesDataEntity, error)
	GetLeavesByDateRange(personalDataID uint, startDate, endDate string) ([]LeavesDataEntity, error)
}

type ServiceLeavesInterface interface {
	RequestLeave(leave LeavesDataEntity) error
	ViewLeaveHistory(personalDataID uint, page, pageSize int) ([]LeavesDataEntity, error)
	UpdateLeaveStatus(leaveID uint, status string) error
}
