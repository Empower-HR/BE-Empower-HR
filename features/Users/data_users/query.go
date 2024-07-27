package datausers

import (
	users "be-empower-hr/features/Users"
	"be-empower-hr/utils"
	"log"

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
func (uq *userQuery) CreateAccountAdmin(account users.PersonalDataEntity, companyName, department, jobPosition string) (uint, error) {
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
				return 0, err
			}
		} else {
			return 0, err
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
		return 0, err
	}

	return personalData.ID, nil
}

// CreateAccountEmployee implements users.DataUserInterface.
func (uq *userQuery) CreateAccountEmployee(personalData users.PersonalDataEntity) (uint, error) {
	var personalDataID uint

	err := uq.db.Transaction(func(tx *gorm.DB) error {
		personalDataModel := &PersonalData{
			CompanyID:      personalData.CompanyID,
			ProfilePicture: personalData.ProfilePicture,
			Name:           personalData.Name,
			Email:          personalData.Email,
			Password:       personalData.Password,
			PhoneNumber:    personalData.PhoneNumber,
			PlaceBirth:     personalData.PlaceBirth,
			BirthDate:      personalData.BirthDate,
			Gender:         personalData.Gender,
			Religion:       personalData.Religion,
			NIK:            personalData.NIK,
			Address:        personalData.Address,
			Role:           "employees",
		}

		if err := tx.Create(&personalDataModel).Error; err != nil {
			return err
		}

		personalDataID = personalDataModel.ID

		for _, employment := range personalData.EmploymentData {
			employmentDataModel := &EmploymentData{
				PersonalDataID:   personalDataModel.ID,
				EmploymentStatus: employment.EmploymentStatus,
				JoinDate:         employment.JoinDate,
				Department:       employment.Department,
				JobPosition:      employment.JobPosition,
				JobLevel:         employment.JobLevel,
				Schedule:         employment.Schedule,
				ApprovalLine:     employment.ApprovalLine,
				Manager:          employment.Manager,
			}

			if err := tx.Create(&employmentDataModel).Error; err != nil {
				return err
			}

			for _, payroll := range employment.Payrolls {
				payrollDataModel := &PayrollData{
					EmploymentDataID: employmentDataModel.ID,
					Salary:           payroll.Salary,
					BankName:         payroll.BankName,
					AccountNumber:    payroll.AccountNumber,
				}

				if err := tx.Create(&payrollDataModel).Error; err != nil {
					return err
				}
			}
		}

		return nil
	})

	return personalDataID, err
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
