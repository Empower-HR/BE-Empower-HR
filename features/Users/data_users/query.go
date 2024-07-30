package datausers

import (
	users "be-empower-hr/features/Users"
	"be-empower-hr/utils"
	"log"
	"time"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.DataUserInterface {
	return &userQuery{
		db: db,
	}
}

// AccountByEmail implements users.DataUserInterface.
func (uq *userQuery) AccountByEmail(email string) (*users.PersonalDataEntity, error) {
	var userData PersonalData
	tx := uq.db.Where("email = ?", email).First(&userData)
	if tx.Error != nil {
		log.Printf("AccountByEmail: Error finding account: %v", tx.Error)
		return nil, tx.Error
	}
	// mapping
	var users = users.PersonalDataEntity{
		PersonalDataID: userData.ID,
		CompanyID:      userData.CompanyID,
		ProfilePicture: userData.ProfilePicture,
		Name:           userData.Name,
		Email:          userData.Email,
		Password:       userData.Password,
		PhoneNumber:    userData.PhoneNumber,
		PlaceBirth:     userData.PlaceBirth,
		BirthDate:      userData.BirthDate,
		Gender:         userData.Gender,
		Religion:       userData.Religion,
		NIK:            userData.NIK,
		Address:        userData.Address,
		Role:           userData.Role,
	}
	return &users, nil
}

// AccountById implements users.DataUserInterface.
func (uq *userQuery) AccountById(userid uint) (*users.PersonalDataEntity, error) {
	var personalData PersonalData
	if err := uq.db.Preload("EmploymentData").Where("id = ?", userid).First(&personalData).Error; err != nil {
		return nil, err
	}

	// Mapping PersonalDataEntity from PersonalData
	personalDataEntity := &users.PersonalDataEntity{
		PersonalDataID: personalData.ID,
		CompanyID:      personalData.CompanyID,
		ProfilePicture: personalData.ProfilePicture,
		Name:           personalData.Name,
		Email:          personalData.Email,
		PhoneNumber:    personalData.PhoneNumber,
		PlaceBirth:     personalData.PlaceBirth,
		BirthDate:      personalData.BirthDate,
		Gender:         personalData.Gender,
		Religion:       personalData.Religion,
		NIK:            personalData.NIK,
		Address:        personalData.Address,
		Role:           personalData.Role,
	}

	// Mapping EmploymentDataEntities from EmploymentData
	for _, employment := range personalData.EmploymentData {
		employmentEntity := users.EmploymentDataEntity{
			EmploymentDataID: employment.ID,
			PersonalDataID:   employment.PersonalDataID,
			EmploymentStatus: employment.EmploymentStatus,
			JoinDate:         employment.JoinDate,
			Department:       employment.Department,
			JobPosition:      employment.JobPosition,
			JobLevel:         employment.JobLevel,
			Schedule:         employment.Schedule,
			ApprovalLine:     employment.ApprovalLine,
			Manager:          employment.Manager,
		}
		// Append EmploymentDataEntity to PersonalDataEntity
		personalDataEntity.EmploymentData = append(personalDataEntity.EmploymentData, employmentEntity)
	}

	return personalDataEntity, nil
}

// CreateAccount implements users.DataUserInterface.
func (uq *userQuery) CreateAccountAdmin(account users.PersonalDataEntity, companyName, department, jobPosition string) (uint, uint, error) {
	var companyData CompanyData

	// Check if company already exists
	err := uq.db.Where("company_name = ?", companyName).First(&companyData).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create new company data if not exists
			companyData = CompanyData{
				CompanyName: companyName,
			}
			if err := uq.db.Create(&companyData).Error; err != nil {
				return 0, 0, err
			}
		} else {
			return 0, 0, err
		}
	}

	// Create PersonalData with the valid CompanyID
	personalData := PersonalData{
		CompanyID:   companyData.ID,
		Name:        account.Name,
		Email:       account.Email,
		Password:    account.Password,
		PhoneNumber: account.PhoneNumber,
		Role:        "admin",
		EmploymentData: []EmploymentData{
			{
				Department:  department,
				JobPosition: jobPosition,
			},
		},
	}
	if err := uq.db.Create(&personalData).Error; err != nil {
		return 0, 0, err
	}

	return personalData.ID, companyData.ID, nil
}

// DeleteAccountAdmin implements users.DataUserInterface.
func (uq *userQuery) DeleteAccountAdmin(userid uint) error {
	err := uq.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("employment_data_id IN (SELECT id FROM employment_data WHERE personal_data_id = ?)", userid).Delete(&PayrollData{}).Error; err != nil {
			return err
		}

		// Delete EmploymentData related to the PersonalData
		if err := tx.Where("personal_data_id = ?", userid).Delete(&EmploymentData{}).Error; err != nil {
			return err
		}

		// Delete PersonalData itself
		if err := tx.Where("id = ?", userid).Delete(&PersonalData{}).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}

// DeleteAccountEmployee implements users.DataUserInterface.
func (uq *userQuery) DeleteAccountEmployeeByAdmin(userid uint) error {
	err := uq.db.Transaction(func(tx *gorm.DB) error {
		// Delete PayrollData first to avoid foreign key constraint issues
		if err := tx.Where("employment_data_id IN (SELECT id FROM employment_data WHERE personal_data_id = ?)", userid).Delete(&PayrollData{}).Error; err != nil {
			return err
		}

		// Delete EmploymentData related to the PersonalData
		if err := tx.Where("personal_data_id = ?", userid).Delete(&EmploymentData{}).Error; err != nil {
			return err
		}

		// Delete PersonalData itself
		if err := tx.Where("id = ?", userid).Delete(&PersonalData{}).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}

// UpdateAccount implements users.DataUserInterface.
func (uq *userQuery) UpdateAccountEmployees(userid uint, account users.PersonalDataEntity) error {
	updates := map[string]interface{}{
		"ProfilePicture": account.ProfilePicture,
		"Name":           account.Name,
		"Email":          account.Email,
		"Password":       account.Password,
		"PhoneNumber":    account.PhoneNumber,
		"PlaceBirth":     account.PlaceBirth,
		"BirthDate":      account.BirthDate,
		"Gender":         account.Gender,
		"Religion":       account.Religion,
		"NIK":            account.NIK,
		"Address":        account.Address,
	}

	// Update the personal data fields
	if err := uq.db.Model(&PersonalData{}).Where("id = ?", userid).Updates(updates).Error; err != nil {
		return err
	}

	return nil
}

// UpdateAccountAdmins implements users.DataUserInterface.
func (uq *userQuery) UpdateAccountAdmins(userid uint, account users.PersonalDataEntity) error {
	updates := map[string]interface{}{
		"ProfilePicture": account.ProfilePicture,
		"Name":           account.Name,
		"Email":          account.Email,
		"Password":       account.Password,
		"PhoneNumber":    account.PhoneNumber,
		"PlaceBirth":     account.PlaceBirth,
		"BirthDate":      account.BirthDate,
		"Gender":         account.Gender,
		"Religion":       account.Religion,
		"NIK":            account.NIK,
		"Address":        account.Address,
	}

	// Update the personal data fields
	if err := uq.db.Model(&PersonalData{}).Where("id = ?", userid).Updates(updates).Error; err != nil {
		return err
	}

	return nil
}

// UpdateProfileEmployments implements users.DataUserInterface.
func (uq *userQuery) UpdateProfileEmployments(userid uint, accounts users.EmploymentDataEntity) error {
	updates := map[string]interface{}{
		"EmploymentStatus": accounts.EmploymentStatus,
		"JoinDate":         accounts.JoinDate,
		"Department":       accounts.Department,
		"JobPosition":      accounts.JobPosition,
		"JobLevel":         accounts.JobLevel,
		"Schedule":         accounts.Schedule,
		"ApprovalLine":     accounts.ApprovalLine,
		"Manager":          accounts.Manager,
	}

	// Update the employment data fields
	if err := uq.db.Model(&EmploymentData{}).Where("personal_data_id = ?", userid).Updates(updates).Error; err != nil {
		return err
	}

	return nil
}

// update employment employee
func (uq *userQuery) UpdateEmploymentEmployee(ID, employeID uint, updateEmploymentEmployee users.EmploymentDataEntity) error {
	cnvQueryModel := ToQueryEmploymentEmployee(updateEmploymentEmployee)
	qry := uq.db.Where("id = ? AND personal_data_id = ? AND deleted_at IS NULL", ID, employeID).Updates(&cnvQueryModel)

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// GetAccountByName implements users.DataUserInterface.
func (uq *userQuery) GetAccountByName(accountName string) ([]users.PersonalDataEntity, error) {
	var personalDataList []PersonalData
	if err := uq.db.Preload("EmploymentData").Where("name LIKE ?", "%"+accountName+"%").Find(&personalDataList).Error; err != nil {
		return nil, err
	}

	var result []users.PersonalDataEntity
	for _, personalData := range personalDataList {
		personalDataEntity := users.PersonalDataEntity{
			PersonalDataID: personalData.ID,
			Name:           personalData.Name,
		}

		for _, employment := range personalData.EmploymentData {
			employmentEntity := users.EmploymentDataEntity{
				EmploymentDataID: employment.ID,
				JobPosition:      employment.JobPosition,
				EmploymentStatus: employment.EmploymentStatus,
			}
			personalDataEntity.EmploymentData = append(personalDataEntity.EmploymentData, employmentEntity)
		}

		result = append(result, personalDataEntity)
	}

	return result, nil
}

// GetAll implements users.DataUserInterface.
func (uq *userQuery) GetAll(page int, pageSize int) ([]users.PersonalDataEntity, error) {
	var personalDataList []PersonalData
	if err := uq.db.Preload("EmploymentData").Find(&personalDataList).Error; err != nil {
		return nil, err
	}
	pagination := utils.NewPagination(page, pageSize)

	tx := uq.db.Limit(pagination.PageSize).Offset(pagination.Offset()).Preload("EmploymentData").Find(&personalDataList)
	if tx.Error != nil {
		log.Printf("Error fetching all accounts: %v", tx.Error)
		return nil, tx.Error
	}

	var result []users.PersonalDataEntity
	for _, personalData := range personalDataList {
		personalDataEntity := users.PersonalDataEntity{
			PersonalDataID: personalData.ID,
			Name:           personalData.Name,
			EmploymentData: []users.EmploymentDataEntity{},
		}

		for _, employment := range personalData.EmploymentData {
			employmentEntity := users.EmploymentDataEntity{
				EmploymentDataID: employment.ID,
				JobPosition:      employment.JobPosition,
				JobLevel:         employment.JobLevel,
				EmploymentStatus: employment.EmploymentStatus,
				JoinDate:         employment.JoinDate,
			}
			personalDataEntity.EmploymentData = append(personalDataEntity.EmploymentData, employmentEntity)
		}

		result = append(result, personalDataEntity)
	}

	return result, nil
}

func (uq *userQuery) GetAccountByJobLevel(jobLevel string) ([]users.PersonalDataEntity, error) {
	var personalDataList []PersonalData
	if err := uq.db.Preload("EmploymentData").
		Joins("JOIN employment_data ON employment_data.personal_data_id = personal_data.id").
		Where("employment_data.job_level = ?", jobLevel).
		Find(&personalDataList).Error; err != nil {
		return nil, err
	}

	var result []users.PersonalDataEntity
	for _, personalData := range personalDataList {
		personalDataEntity := users.PersonalDataEntity{
			PersonalDataID: personalData.ID,
			Name:           personalData.Name,
		}

		for _, employment := range personalData.EmploymentData {
			if employment.JobLevel == jobLevel {
				employmentEntity := users.EmploymentDataEntity{
					EmploymentDataID: employment.ID,
					JobPosition:      employment.JobPosition,
					EmploymentStatus: employment.EmploymentStatus,
					JoinDate:         employment.JoinDate,
				}
				personalDataEntity.EmploymentData = append(personalDataEntity.EmploymentData, employmentEntity)
			}
		}

		result = append(result, personalDataEntity)
	}

	return result, nil
}

// Add Employe
func (uq *userQuery) CreatePersonal(CompanyID uint, addPersonal users.PersonalDataEntity) (uint, error) {
	cnvQuery := ToPersonalDataQuery(addPersonal)
	cnvQuery.CompanyID = CompanyID
	err := uq.db.Create(&cnvQuery).Error

	if err != nil {
		return 0, err
	}

	return cnvQuery.ID, nil
}

// Add employment data
func (uq *userQuery) CreateEmployment(personalID uint, addEmployment users.EmploymentDataEntity) (uint, error) {
	cnvQuery := ToEmploymentQuery(addEmployment)
	cnvQuery.PersonalDataID = personalID
	err := uq.db.Create(&cnvQuery).Error

	if err != nil {
		return 0, err
	}

	return cnvQuery.ID, nil
}

// Add Payroll data
func (uq *userQuery) CreatePayroll(employmentID uint, addPayroll users.PayrollDataEntity) error {
	cnvQuery := ToPayrollQuery(addPayroll)
	cnvQuery.EmploymentDataID = employmentID
	err := uq.db.Create(&cnvQuery).Error

	if err != nil {
		return err
	}

	return nil
}

// CreateLeaves implements users.DataUserInterface.
func (uq *userQuery) CreateLeaves(PersonalID uint, addLeaves users.LeavesDataEntity) (uint, error) {
	cnvQuery := ToLeavesQuery(addLeaves)
	cnvQuery.PersonalDataID = PersonalID
	err := uq.db.Create(&cnvQuery).Error

	if err != nil {
		return 0, err
	}

	return cnvQuery.ID, nil
}

// CountTotalUsers menghitung jumlah total users berdasarkan CompanyID
func (uq *userQuery) CountTotalUsers(companyID uint) (int64, error) {
	var count int64
	if err := uq.db.Model(&PersonalData{}).Where("company_id = ?", companyID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// CountMaleUsers menghitung jumlah gender laki-laki berdasarkan CompanyID
func (uq *userQuery) CountMaleUsers(companyID uint) (int64, error) {
	var count int64
	if err := uq.db.Model(&PersonalData{}).Where("company_id = ? AND gender = ?", companyID, "male").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// CountFemaleUsers menghitung jumlah gender perempuan berdasarkan CompanyID
func (uq *userQuery) CountFemaleUsers(companyID uint) (int64, error) {
	var count int64
	if err := uq.db.Model(&PersonalData{}).Where("company_id = ? AND gender = ?", companyID, "female").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (uq *userQuery) CountContractUsers(companyID uint) (int64, error) {
	var count int64
	if err := uq.db.Model(&EmploymentData{}).
		Where("personal_data_id IN (SELECT id FROM personal_data WHERE company_id = ?)", companyID).
		Where("employment_status = ?", "contract").
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// CountPermanentUsers menghitung jumlah pengguna dengan status pekerjaan "permanent" berdasarkan CompanyID
func (uq *userQuery) CountPermanentUsers(companyID uint) (int64, error) {
	var count int64
	if err := uq.db.Model(&EmploymentData{}).
		Where("personal_data_id IN (SELECT id FROM personal_data WHERE company_id = ?)", companyID).
		Where("employment_status = ?", "permanent").
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (uq *userQuery) CountPayrollUsers(companyID uint) (int64, error) {
	var count int64
	if err := uq.db.Model(&PayrollData{}).
		Where("employment_data_id IN (SELECT id FROM employment_data WHERE personal_data_id IN (SELECT id FROM personal_data WHERE company_id = ?))", companyID).
		Where("deleted_at IS NULL").
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (uq *userQuery) GetCompanyIDByName(companyName string) (uint, error) {
	var companyData CompanyData
	if err := uq.db.Where("company_name = ?", companyName).First(&companyData).Error; err != nil {
		return 0, err
	}
	return companyData.ID, nil
}

func (uq *userQuery) CountPendingLeaves(companyID uint) (int64, error) {
	var count int64
	if err := uq.db.Model(&LeavesData{}).
		Where("personal_data_id IN (SELECT id FROM personal_data WHERE company_id = ?)", companyID).
		Where("status = ?", "pending").
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (uq *userQuery) CountAttendanceHadir(companyID uint) (int64, error) {
	var count int64
	if err := uq.db.Model(&Attandance{}).
		Where("personal_data_id IN (SELECT id FROM personal_data WHERE company_id = ?)", companyID).
		Where("status = ?", "hadir").
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (uq *userQuery) Dashboard(companyID uint) (*users.DashboardStats, error) {
	var stats users.DashboardStats

	// Fetch statistics
	totalUsers, err := uq.CountTotalUsers(companyID)
	if err != nil {
		log.Printf("Error counting total users: %v", err)
		return nil, err
	}
	stats.TotalUsers = totalUsers

	maleUsers, err := uq.CountMaleUsers(companyID)
	if err != nil {
		log.Printf("Error counting male users: %v", err)
		return nil, err
	}
	stats.MaleUsers = maleUsers

	femaleUsers, err := uq.CountFemaleUsers(companyID)
	if err != nil {
		log.Printf("Error counting female users: %v", err)
		return nil, err
	}
	stats.FemaleUsers = femaleUsers

	contractUsers, err := uq.CountContractUsers(companyID)
	if err != nil {
		log.Printf("Error counting contract users: %v", err)
		return nil, err
	}
	stats.ContractUsers = contractUsers

	permanentUsers, err := uq.CountPermanentUsers(companyID)
	if err != nil {
		log.Printf("Error counting permanent users: %v", err)
		return nil, err
	}
	stats.PermanentUsers = permanentUsers

	// menghitung presentase
	if totalUsers > 0 {
		stats.MalePercentage = (float64(maleUsers) / float64(totalUsers)) * 100
		stats.FemalePercentage = (float64(femaleUsers) / float64(totalUsers)) * 100
		stats.ContractUsersPercentage = (float64(contractUsers) / float64(totalUsers)) * 100
		stats.PermanentUsersPercentage = (float64(permanentUsers) / float64(totalUsers)) * 100
	}

	payrollRecords, err := uq.CountPayrollUsers(companyID)
	if err != nil {
		log.Printf("Error counting payroll records: %v", err)
		return nil, err
	}
	stats.PayrollRecords = payrollRecords

	pendingLeaves, err := uq.CountPendingLeaves(companyID)
	if err != nil {
		log.Printf("Error counting pending leaves: %v", err)
		return nil, err
	}
	stats.LeavesPending = pendingLeaves

	attendanceHadir, err := uq.CountAttendanceHadir(companyID)
	if err != nil {
		log.Printf("Error counting attendance with status 'hadir': %v", err)
		return nil, err
	}
	stats.AttendanceHadir = attendanceHadir

	var name string
	if err := uq.db.Model(&PersonalData{}).
		Where("company_id = ? AND deleted_at IS NULL", companyID).
		Select("name").
		Limit(1).
		Pluck("name", &name).Error; err != nil {
		log.Printf("Error fetching user name: %v", err)
		return nil, err
	}
	stats.PersonalDataNames = name

	stats.CurrentDate = time.Now().Format("Monday, 02 January 2006")

	return &stats, nil
}

func (uq *userQuery) DashboardEmployees(companyID uint) (*users.DashboardStats, error) {
	var stats users.DashboardStats

	// Fetch statistics
	totalUsers, err := uq.CountTotalUsers(companyID)
	if err != nil {
		log.Printf("Error counting total users: %v", err)
		return nil, err
	}
	stats.TotalUsers = totalUsers

	maleUsers, err := uq.CountMaleUsers(companyID)
	if err != nil {
		log.Printf("Error counting male users: %v", err)
		return nil, err
	}
	stats.MaleUsers = maleUsers

	femaleUsers, err := uq.CountFemaleUsers(companyID)
	if err != nil {
		log.Printf("Error counting female users: %v", err)
		return nil, err
	}
	stats.FemaleUsers = femaleUsers

	contractUsers, err := uq.CountContractUsers(companyID)
	if err != nil {
		log.Printf("Error counting contract users: %v", err)
		return nil, err
	}
	stats.ContractUsers = contractUsers

	permanentUsers, err := uq.CountPermanentUsers(companyID)
	if err != nil {
		log.Printf("Error counting permanent users: %v", err)
		return nil, err
	}
	stats.PermanentUsers = permanentUsers

	// menghitung presentase
	if totalUsers > 0 {
		stats.MalePercentage = (float64(maleUsers) / float64(totalUsers)) * 100
		stats.FemalePercentage = (float64(femaleUsers) / float64(totalUsers)) * 100
		stats.ContractUsersPercentage = (float64(contractUsers) / float64(totalUsers)) * 100
		stats.PermanentUsersPercentage = (float64(permanentUsers) / float64(totalUsers)) * 100
	}

	var name string
	if err := uq.db.Model(&PersonalData{}).
		Where("company_id = ? AND deleted_at IS NULL", companyID).
		Select("name").
		Limit(1).
		Pluck("name", &name).Error; err != nil {
		log.Printf("Error fetching user name: %v", err)
		return nil, err
	}
	stats.PersonalDataNames = name

	stats.CurrentDate = time.Now().Format("Monday, 02 January 2006")

	return &stats, nil
}
