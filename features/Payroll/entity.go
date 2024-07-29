package payroll

type PayrollDataEntity struct {
	PayrollID        uint
	EmploymentDataID uint
	Salary           float64
	BankName         string
	AccountNumber    int
}

type PayrollResponse struct {
	ID             uint   `json:"id"`
	EmploymentName string `json:"employee_name"`
	Date           string `json:"date"`
	Position       string `json:"position"`
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
	Status         string
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
	Payrolls         []PayrollDataEntity
}

type DataPayrollInterface interface {
	CreatePayroll(payroll PayrollDataEntity) (PayrollDataEntity, error)
	GetAllPayroll() ([]PayrollResponse, error)
	GetEmpById(employee uint) (EmploymentDataEntity, error)
}

type ServicePayrollInterface interface {
	CreatePayroll(payroll PayrollDataEntity) (PayrollDataEntity, error)
	GetAllPayroll() ([]PayrollResponse, error)
}
