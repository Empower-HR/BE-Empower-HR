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

type DataUserInterface interface {
	CreateAccountAdmin(account PersonalDataEntity, companyName, department, jobPosition string) (uint, error)
	GetAll(page, pageSize int) ([]PersonalDataEntity, error)
	GetAccountByName(accountName string) ([]PersonalDataEntity, error)
	GetAccountByDepartment(Department string) ([]PersonalDataEntity, error)
	CreateAccountEmployee(account PersonalDataEntity) (uint, error)
	AccountByEmail(email string) (*PersonalDataEntity, error)
	AccountById(userid uint) (*PersonalDataEntity, error)
	UpdateAccountEmployees(userid uint, account PersonalDataEntity) error
	UpdateAccountAdmins(userid uint, account PersonalDataEntity) error
	DeleteAccountAdmin(userid uint) error
	DeleteAccountEmployee(userid uint) error
}

type ServiceUserInterface interface {
	RegistrasiAccountAdmin(accounts PersonalDataEntity, companyName, department, jobPosition string) (uint, error)
	GetAllAccount(name, department string, page, pageSize int) ([]PersonalDataEntity, error)
	RegistrasiAccountEmployee(personalData PersonalDataEntity) (uint, error)
	LoginAccount(email string, password string) (data *PersonalDataEntity, token string, err error)
	GetProfile(userid uint) (data *PersonalDataEntity, err error)
	UpdateProfileEmployees(userid uint, accounts PersonalDataEntity) error
	UpdateProfileAdmins(userid uint, accounts PersonalDataEntity) error
	DeleteAccountAdmin(userid uint) error
	DeleteAccountEmployee(userid uint) error
}
