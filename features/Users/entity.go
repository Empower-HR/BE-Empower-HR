package users

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
	EmploymentData []EmploymentDataEntity
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

type PayrollDataEntity struct {
	PayrollDataID    uint
	EmploymentDataID uint
	Salary           float64
	BankName         string
	AccountNumber    int
}

type LeavesDataEntity struct {
	LeavesID       uint
	StartDate      string
	EndDate        string
	Reason         string
	Status         string
	TotalLeave     int
	PersonalDataID uint
}

type Attandance struct {
	ID             uint
	PersonalDataID uint
	Clock_in       string
	Clock_out      string
	Status         string
	Date           string
	Long           string
	Lat            string
	Notes          string
}

type DashboardStats struct {
	TotalUsers               int64
	MaleUsers                int64
	FemaleUsers              int64
	ContractUsers            int64
	PermanentUsers           int64
	PayrollRecords           int64
	LeavesPending            int64
	MalePercentage           float64
	FemalePercentage         float64
	ContractUsersPercentage  float64
	PermanentUsersPercentage float64
	PersonalDataNames        string
	AttendanceHadir          int64
	CurrentDate              string
}

type DataUserInterface interface {
	CreateAccountAdmin(account PersonalDataEntity, companyName, department, jobPosition string) (uint, uint, error)
	GetAll(page, pageSize int) ([]PersonalDataEntity, error)
	GetAccountByName(accountName string) ([]PersonalDataEntity, error)
	GetAccountByJobLevel(jobLevel string) ([]PersonalDataEntity, error)
	AccountByEmail(email string) (*PersonalDataEntity, error)
	AccountById(userid uint) (*PersonalDataEntity, error)
	UpdateAccountEmployees(userid uint, account PersonalDataEntity) error
	UpdateAccountAdmins(userid uint, account PersonalDataEntity) error
	UpdateProfileEmployments(userid uint, accounts EmploymentDataEntity) error
	DeleteAccountAdmin(userid uint) error
	DeleteAccountEmployeeByAdmin(userid uint) error
	UpdateEmploymentEmployee(ID, employeID uint, updateEmploymentEmployee EmploymentDataEntity) error
	CreatePersonal(CompanyID uint, addPersonal PersonalDataEntity) (uint, error)
	CreateEmployment(personalID uint, addEmployment EmploymentDataEntity) (uint, error)
	CreatePayroll(employmentID uint, addPayroll PayrollDataEntity) error
	CountTotalUsers(companyID uint) (int64, error)
	CountMaleUsers(companyID uint) (int64, error)
	CountFemaleUsers(companyID uint) (int64, error)
	CountContractUsers(companyID uint) (int64, error)
	CountPermanentUsers(companyID uint) (int64, error)
	CountPayrollUsers(companyID uint) (int64, error)
	GetCompanyIDByName(companyName string) (uint, error)
	Dashboard(companyID uint) (*DashboardStats, error)
	CreateLeaves(PersonalID uint, addLeaves LeavesDataEntity) (uint, error)
}

type ServiceUserInterface interface {
	RegistrasiAccountAdmin(accounts PersonalDataEntity, companyName, department, jobPosition string) (uint, uint, error)
	GetAllAccount(name, jobLevel string, page, pageSize int) ([]PersonalDataEntity, error)
	LoginAccount(email string, password string) (data *PersonalDataEntity, token string, err error)
	GetProfile(userid uint) (data *PersonalDataEntity, err error)
	GetProfileById(userid uint) (data *PersonalDataEntity, err error)
	UpdateProfileEmployees(userid uint, accounts PersonalDataEntity) error
	UpdateProfileAdmins(userid uint, accounts PersonalDataEntity) error
	UpdateProfileEmployments(userid uint, accounts EmploymentDataEntity) error
	DeleteAccountAdmin(userid uint) error
	DeleteAccountEmployeeByAdmin(userid uint) error
	UpdateEmploymentEmployee(ID, employeID uint, updateEmploymentEmployee EmploymentDataEntity) error
	CreateNewEmployee(addPersonal PersonalDataEntity, addEmployment EmploymentDataEntity, addPayroll PayrollDataEntity, addLeaves LeavesDataEntity) error
	Dashboard(companyID uint) (*DashboardStats, error)
}
