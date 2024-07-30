package leaves

type LeavesDataEntity struct {
	LeavesID       uint
	Name           string
	JobPosition    string
	StartDate      string
	EndDate        string
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
	UpdateLeaveStatus(leaveID uint, updatesleaves LeavesDataEntity) error
	GetLeaveHistory(personalDataID uint, page, pageSize int) ([]LeavesDataEntity, error)
	GetLeavesByStatus(personalDataID uint, status string) ([]LeavesDataEntity, error)
	GetLeavesByDateRange(personalDataID uint, startDate, endDate string) ([]LeavesDataEntity, error)
	GetLeavesDetail(leaveID uint) (*LeavesDataEntity, error)
	GetLeaveHistoryEmployee(personalDataID uint, page, pageSize int) ([]LeavesDataEntity, error)
}

type ServiceLeavesInterface interface {
	RequestLeave(leave LeavesDataEntity) error
	ViewLeaveHistory(personalDataID uint, page, pageSize int, status string, startDate, endDate string) ([]LeavesDataEntity, error)
	UpdateLeaveStatus(leaveID uint, updatesleaves LeavesDataEntity) error
	GetLeavesByID(leaveID uint) (leaves *LeavesDataEntity, err error)
	ViewLeaveHistoryEmployee(personalDataID uint, page, pageSize int, status string, startDate, endDate string) ([]LeavesDataEntity, error)
}
