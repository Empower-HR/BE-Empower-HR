package service

import (
	"be-empower-hr/app/middlewares"
	companies "be-empower-hr/features/Companies"
	users "be-empower-hr/features/Users"
	"be-empower-hr/utils"
	"be-empower-hr/utils/encrypts"
	"errors"
	"log"
)

type userService struct {
	userData          users.DataUserInterface
	hashService       encrypts.HashInterface
	middlewareservice middlewares.MiddlewaresInterface
	accountUtility    utils.AccountUtilityInterface
	company           companies.Query
}

func New(ud users.DataUserInterface, hash encrypts.HashInterface, mi middlewares.MiddlewaresInterface, au utils.AccountUtilityInterface, cm companies.Query) users.ServiceUserInterface {
	return &userService{
		userData:          ud,
		hashService:       hash,
		middlewareservice: mi,
		accountUtility:    au,
		company:           cm,
	}
}

// DeleteAccountAdmin implements users.ServiceUserInterface.
func (us *userService) DeleteAccountAdmin(userid uint) error {
	err := us.userData.DeleteAccountAdmin(userid)
	if err != nil {
		log.Println("Error deleting admin account:", err)
		return err
	}
	return nil
}

// DeleteAccountEmployee implements users.ServiceUserInterface.
func (us *userService) DeleteAccountEmployeeByAdmin(userid uint) error {
	err := us.userData.DeleteAccountEmployeeByAdmin(userid)
	if err != nil {
		log.Println("Error deleting account:", err)
		return err
	}
	return nil

}

// GetProfile implements users.ServiceUserInterface.
func (us *userService) GetProfile(userid uint) (data *users.PersonalDataEntity, err error) {
	data, err = us.userData.AccountById(userid)
	if err != nil {
		log.Println("Error getting profile:", err)
		return nil, err
	}
	return data, nil
}

func (us *userService) GetProfileById(userid uint) (data *users.PersonalDataEntity, err error) {
	data, err = us.userData.AccountById(userid)
	if err != nil {
		log.Println("Error getting profile:", err)
		return nil, err
	}
	return data, nil
}

// LoginAccount implements users.ServiceUserInterface.
func (us *userService) LoginAccount(email string, password string) (data *users.PersonalDataEntity, token string, err error) {
	data, err = us.userData.AccountByEmail(email)
	if err != nil {
		log.Println("Error logging in:", err)
		return nil, "", err
	}

	isLoginValid := us.hashService.CheckPasswordHash(data.Password, password)
	log.Printf("LoginAccount: Checking password for user %s. Stored hash: %s, Provided password: %s, IsValid: %v", email, data.Password, password, isLoginValid)
	if !isLoginValid {
		return nil, "", errors.New("email atau password tidak sesuai")
	}

	token, errJWT := us.middlewareservice.CreateToken(int(data.PersonalDataID))
	if errJWT != nil {
		log.Println("Error creating token:", errJWT)
		return nil, "", errJWT
	}
	return data, token, nil
}

// RegistrasiAccountAdmin implements users.ServiceUserInterface.
func (us *userService) RegistrasiAccountAdmin(accounts users.PersonalDataEntity, companyName string, department string, jobPosition string) (uint, uint, error) {
	if accounts.Name == "" || accounts.Email == "" || accounts.Password == "" {
		log.Println("Nama, email, atau password tidak boleh kosong")
		return 0, 0, errors.New("nama/email/password tidak boleh kosong")
	}

	if err := us.accountUtility.EmailValidator(accounts.Email); err != nil {
		log.Println("Validasi email gagal:", err)
		return 0, 0, err
	}
	if err := us.accountUtility.PasswordValidator(accounts.Password); err != nil {
		log.Println("Validasi password gagal:", err)
		return 0, 0, err
	}
	if err := us.accountUtility.PhoneNumberValidator(accounts.PhoneNumber); err != nil {
		log.Println("Validasi nomor telepon gagal:", err)
		return 0, 0, err
	}

	// Hash password
	var errHash error
	if accounts.Password, errHash = us.hashService.HashPassword(accounts.Password); errHash != nil {
		log.Println("Hash password gagal:", errHash)
		return 0, 0, errHash
	}

	personalDataID, companyID, err := us.userData.CreateAccountAdmin(accounts, companyName, department, jobPosition)
	if err != nil {
		log.Println("Error mendaftarkan akun:", err)
		return 0, 0, err
	}

	log.Println("Akun admin berhasil didaftarkan dengan ID:", personalDataID)
	return personalDataID, companyID, nil
}

// UpdateProfileAdmins implements users.ServiceUserInterface.
func (us *userService) UpdateProfileAdmins(userid uint, accounts users.PersonalDataEntity) error {
	if accounts.Email != "" {
		if err := us.accountUtility.EmailValidator(accounts.Email); err != nil {
			return err
		}
	}

	if accounts.PhoneNumber != "" {
		if err := us.accountUtility.PhoneNumberValidator(accounts.PhoneNumber); err != nil {
			return err
		}
	}

	if accounts.Password != "" {
		var errHash error
		if accounts.Password, errHash = us.hashService.HashPassword(accounts.Password); errHash != nil {
			return errHash
		}
	}

	err := us.userData.UpdateAccountAdmins(userid, accounts)
	if err != nil {
		log.Println("Error updating admin profile:", err)
		return err
	}

	return nil
}

func (us *userService) UpdateProfileEmployments(userid uint, accounts users.EmploymentDataEntity) error {
	if userid == 0 {
		return errors.New("invalid user ID")
	}
	err := us.userData.UpdateProfileEmployments(userid, accounts)
	if err != nil {
		log.Printf("failed to update profile employments: %v", err)
		return err
	}

	return nil
}

// UpdateProfileEmployees implements users.ServiceUserInterface.
func (us *userService) UpdateProfileEmployees(userid uint, accounts users.PersonalDataEntity) error {
	if accounts.Email != "" {
		if err := us.accountUtility.EmailValidator(accounts.Email); err != nil {
			return err
		}
	}

	if accounts.PhoneNumber != "" {
		if err := us.accountUtility.PhoneNumberValidator(accounts.PhoneNumber); err != nil {
			return err
		}
	}

	if accounts.Password != "" {
		var errHash error
		if accounts.Password, errHash = us.hashService.HashPassword(accounts.Password); errHash != nil {
			return errHash
		}
	}

	err := us.userData.UpdateAccountEmployees(userid, accounts)
	if err != nil {
		log.Println("Error updating employee profile:", err)
		return err
	}

	return nil
}

// GetAllAccount implements users.ServiceUserInterface.
func (us *userService) GetAllAccount(name, jobLevel string, page int, pageSize int) ([]users.PersonalDataEntity, error) {
	if name != "" {
		product, err := us.userData.GetAccountByName(name)
		if err != nil {
			log.Println("Error retrieving account by name:", err)
			return nil, err
		}
		return product, nil
	}
	if jobLevel != "" {
		product, err := us.userData.GetAccountByJobLevel(jobLevel)
		if err != nil {
			log.Println("Error retrieving account by department:", err)
			return nil, err
		}
		return product, nil
	}

	allAccount, err := us.userData.GetAll(page, pageSize)
	if err != nil {
		log.Println("Error retrieving all account:", err)
		return nil, err
	}
	return allAccount, nil

}

// update employment employee
func (us *userService) UpdateEmploymentEmployee(ID uint, employeID uint, updateEmploymentEmployee users.EmploymentDataEntity) error {
	err := us.userData.UpdateEmploymentEmployee(ID, employeID, updateEmploymentEmployee)

	if err != nil {
		log.Println("Error update employment account:", err)
		return err
	}

	return nil
}

// Create Employment
func (us *userService) CreateNewEmployee(addPersonal users.PersonalDataEntity, addEmployment users.EmploymentDataEntity, addPayroll users.PayrollDataEntity, addLeaves users.LeavesDataEntity) error {
	// get company ID
	result, err := us.company.GetCompany()
	if err != nil {
		log.Println("Error get company account:", err)
	}

	if addPersonal.Name == "" || addPersonal.Email == "" {
		log.Println("Nama atau email tidak boleh kosong")
		return errors.New("nama/email tidak boleh kosong")
	}

	if err := us.accountUtility.EmailValidator(addPersonal.Email); err != nil {
		log.Println("Validasi email gagal:", err)
		return err
	}
	if err := us.accountUtility.PhoneNumberValidator(addPersonal.PhoneNumber); err != nil {
		log.Println("Validasi nomor telepon gagal:", err)
		return err
	}
	if err := us.accountUtility.GenderValidator(addPersonal.Gender); err != nil {
		log.Println("Validasi jenis kelamin gagal:", err)
		return err
	}
	if err := us.accountUtility.ReligionValidator(addPersonal.Religion); err != nil {
		log.Println("Validasi agama gagal:", err)
		return err
	}

	if err := us.accountUtility.EmploymentStatusValidator(addEmployment.EmploymentStatus); err != nil {
		log.Println("Validasi status pekerjaan gagal:", err)
		return err
	}
	if err := us.accountUtility.JobLevelValidator(addEmployment.JobLevel); err != nil {
		log.Println("Validasi level pekerjaan gagal:", err)
		return err
	}

	// Generate and hash password
	var errHash error
	password, err := us.accountUtility.GeneratePassword(8)
	if err != nil {
		log.Println("Generate password gagal:", err)
		return err
	}
	if addPersonal.Password, errHash = us.hashService.HashPassword(password); errHash != nil {
		log.Println("Hash password gagal:", errHash)
		return errHash
	}

	// Print generated password (optional)
	log.Printf("Generated password for %s: %s", addPersonal.Email, password)

	addPersonal.Role = "employees"
	personalID, err := us.userData.CreatePersonal(result.ID, addPersonal)
	if err != nil {
		log.Println("Error add personal account:", err)
	}

	employmentID, err := us.userData.CreateEmployment(personalID, addEmployment)
	if err != nil {
		log.Println("Error add employment account:", err)
	}

	err = us.userData.CreatePayroll(employmentID, addPayroll)
	if err != nil {
		log.Println("Error add payroll account:", err)
	}

	addLeaves.TotalLeave = 12
	_, err = us.userData.CreateLeaves(personalID, addLeaves)
	if err != nil {
		log.Println("Error add leaves account:", err)
		return err
	}

	// Send email
	subject := "Selamat Bergabung di EmpowerHR!"

	body := "Halo " + addPersonal.Name + ",\n\n" +
		"Selamat datang di EmpowerHR, tempat di mana kami menyambut Anda ke dalam tim kami dengan hangat! 🎉\n\n" +
		"Berikut adalah kredensial login Anda:\n" +
		"Email: " + addPersonal.Email + "\n" +
		"Password: " + password + "\n\n" +
		"EmpowerHR adalah startup yang inovatif di bidang penyedia software SDM, dan kami sangat senang memiliki Anda sebagai bagian dari tim kami.\n\n" +
		"Untuk memulai, silakan login ke sistem kami menggunakan kredensial di atas. Kami percaya Anda akan membawa banyak kontribusi berharga untuk perusahaan kami. Jika Anda memiliki pertanyaan atau membutuhkan bantuan, jangan ragu untuk menghubungi tim kami.\n\n" +
		"Selamat bergabung, dan semoga perjalanan Anda bersama EmpowerHR penuh dengan pengalaman yang menyenangkan dan sukses!\n\n" +
		"Terima kasih dan salam hangat,\n" +
		"Tim EmpowerHR"

	// Kirim email
	if err := us.accountUtility.SendEmail(addPersonal.Email, subject, body); err != nil {
		log.Println("Error mengirim email:", err)
		return err
	}

	log.Println("Akun employee berhasil didaftarkan dengan ID:", employmentID)

	return err
}
