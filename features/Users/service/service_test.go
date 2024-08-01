package service_test

import (
	"errors"
	"testing"

	users "be-empower-hr/features/Users"
	"be-empower-hr/features/Users/service"
	"be-empower-hr/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteAccountAdmin(t *testing.T) {
	qry := mocks.NewDataUserInterface(t)
	srv := service.New(qry, nil, nil, nil, nil)

	t.Run("Success Delete Account Admin", func(t *testing.T) {
		userID := uint(1)

		qry.On("DeleteAccountAdmin", userID).Return(nil).Once()

		err := srv.DeleteAccountAdmin(userID)

		qry.AssertExpectations(t)
		assert.Nil(t, err)
	})

	t.Run("Error Delete Account Admin", func(t *testing.T) {
		userID := uint(1)

		qry.On("DeleteAccountAdmin", userID).Return(errors.New("internal server error")).Once()

		err := srv.DeleteAccountAdmin(userID)

		qry.AssertExpectations(t)
		assert.Error(t, err)
		assert.ErrorContains(t, err, "internal server error")
	})
}

func TestDeleteAccountEmployeeByAdmin(t *testing.T) {
	qry := mocks.NewDataUserInterface(t)
	srv := service.New(qry, nil, nil, nil, nil)

	t.Run("Success Delete Account Employee By Admin", func(t *testing.T) {
		userID := uint(1)

		qry.On("DeleteAccountEmployeeByAdmin", userID).Return(nil).Once()

		err := srv.DeleteAccountEmployeeByAdmin(userID)

		qry.AssertExpectations(t)
		assert.Nil(t, err)
	})

	t.Run("Error Delete Account Employee By Admin", func(t *testing.T) {
		userID := uint(1)

		qry.On("DeleteAccountEmployeeByAdmin", userID).Return(errors.New("internal server error")).Once()

		err := srv.DeleteAccountEmployeeByAdmin(userID)

		qry.AssertExpectations(t)
		assert.Error(t, err)
		assert.ErrorContains(t, err, "internal server error")
	})
}

func TestGetProfile(t *testing.T) {
	qry := mocks.NewDataUserInterface(t)
	srv := service.New(qry, nil, nil, nil, nil)

	t.Run("Success Get Profile", func(t *testing.T) {
		userID := uint(1)
		expectedResult := &users.PersonalDataEntity{
			Name:  "John Doe",
			Email: "johndoe@example.com",
		}

		qry.On("AccountById", userID).Return(expectedResult, nil).Once()

		result, err := srv.GetProfile(userID)

		qry.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("Error Get Profile", func(t *testing.T) {
		userID := uint(1)

		qry.On("AccountById", userID).Return(nil, errors.New("internal server error")).Once()

		result, err := srv.GetProfile(userID)

		qry.AssertExpectations(t)
		assert.Error(t, err)
		assert.ErrorContains(t, err, "internal server error")
		assert.Nil(t, result)
	})
}

func TestGetProfileById(t *testing.T) {
	qry := mocks.NewDataUserInterface(t)
	srv := service.New(qry, nil, nil, nil, nil)

	t.Run("Success Get Profile By Id", func(t *testing.T) {
		userID := uint(1)
		expectedResult := &users.PersonalDataEntity{
			Name:  "John Doe",
			Email: "johndoe@example.com",
		}

		qry.On("AccountById", userID).Return(expectedResult, nil).Once()

		result, err := srv.GetProfileById(userID)

		qry.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("Error Get Profile By Id", func(t *testing.T) {
		userID := uint(1)

		qry.On("AccountById", userID).Return(nil, errors.New("internal server error")).Once()

		result, err := srv.GetProfileById(userID)

		qry.AssertExpectations(t)
		assert.Error(t, err)
		assert.ErrorContains(t, err, "internal server error")
		assert.Nil(t, result)
	})
}

func TestLoginAccount(t *testing.T) {
	qry := mocks.NewDataUserInterface(t)
	hash := mocks.NewHashInterface(t)
	mw := mocks.NewMiddlewaresInterface(t)
	srv := service.New(qry, hash, mw, nil, nil)

	t.Run("Success Login Account", func(t *testing.T) {
		email := "johndoe@example.com"
		password := "password123"
		expectedResult := &users.PersonalDataEntity{
			PersonalDataID: 1,
			Email:          email,
			Password:       "hashedpassword",
		}
		token := "valid.token.here"

		qry.On("AccountByEmail", email).Return(expectedResult, nil).Once()
		hash.On("CheckPasswordHash", "hashedpassword", password).Return(true).Once()
		mw.On("CreateToken", 1).Return(token, nil).Once()

		result, tokenResult, err := srv.LoginAccount(email, password)

		qry.AssertExpectations(t)
		hash.AssertExpectations(t)
		mw.AssertExpectations(t)

		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
		assert.Equal(t, token, tokenResult)
	})

	t.Run("Error Login Account - Invalid Password", func(t *testing.T) {
		email := "johndoe@example.com"
		password := "wrongpassword"
		expectedResult := &users.PersonalDataEntity{
			PersonalDataID: 1,
			Email:          email,
			Password:       "hashedpassword",
		}

		qry.On("AccountByEmail", email).Return(expectedResult, nil).Once()
		hash.On("CheckPasswordHash", "hashedpassword", password).Return(false).Once()

		result, tokenResult, err := srv.LoginAccount(email, password)

		qry.AssertExpectations(t)
		hash.AssertExpectations(t)

		assert.Error(t, err)
		assert.Equal(t, "email atau password tidak sesuai", err.Error())
		assert.Nil(t, result)
		assert.Empty(t, tokenResult)
	})

	t.Run("Error Login Account - User Not Found", func(t *testing.T) {
		email := "notfound@example.com"
		password := "password123"

		qry.On("AccountByEmail", email).Return(nil, errors.New("user not found")).Once()

		result, tokenResult, err := srv.LoginAccount(email, password)

		qry.AssertExpectations(t)

		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
		assert.Nil(t, result)
		assert.Empty(t, tokenResult)
	})

	t.Run("Error Login Account - JWT Creation Error", func(t *testing.T) {
		email := "johndoe@example.com"
		password := "password123"
		expectedResult := &users.PersonalDataEntity{
			PersonalDataID: 1,
			Email:          email,
			Password:       "hashedpassword",
		}

		qry.On("AccountByEmail", email).Return(expectedResult, nil).Once()
		hash.On("CheckPasswordHash", "hashedpassword", password).Return(true).Once()
		mw.On("CreateToken", 1).Return("", errors.New("jwt error")).Once()

		result, tokenResult, err := srv.LoginAccount(email, password)

		qry.AssertExpectations(t)
		hash.AssertExpectations(t)
		mw.AssertExpectations(t)

		assert.Error(t, err)
		assert.Equal(t, "jwt error", err.Error())
		assert.Nil(t, result)
		assert.Empty(t, tokenResult)
	})
}

func TestRegistrasiAccountAdmin(t *testing.T) {
	qry := mocks.NewDataUserInterface(t)
	hash := mocks.NewHashInterface(t)
	au := mocks.NewAccountUtilityInterface(t)
	srv := service.New(qry, hash, nil, au, nil)

	t.Run("Success Registrasi Account Admin", func(t *testing.T) {
		accounts := users.PersonalDataEntity{
			Name:        "John Doe",
			Email:       "johndoe@example.com",
			Password:    "password123",
			PhoneNumber: "081234567890",
		}
		companyName := "Test Company"
		department := "IT"
		jobPosition := "Developer"
		expectedPersonalID := uint(1)
		expectedCompanyID := uint(1)
		hashedPassword := "hashedpassword"

		au.On("EmailValidator", accounts.Email).Return(nil).Once()
		au.On("PasswordValidator", accounts.Password).Return(nil).Once()
		au.On("PhoneNumberValidator", accounts.PhoneNumber).Return(nil).Once()
		hash.On("HashPassword", accounts.Password).Return(hashedPassword, nil).Once()

		qry.On("CreateAccountAdmin", mock.MatchedBy(func(a users.PersonalDataEntity) bool {
			return a.Name == accounts.Name && a.Email == accounts.Email && a.Password == hashedPassword && a.PhoneNumber == accounts.PhoneNumber
		}), companyName, department, jobPosition).Return(expectedPersonalID, expectedCompanyID, nil).Once()

		personalID, companyID, err := srv.RegistrasiAccountAdmin(accounts, companyName, department, jobPosition)

		qry.AssertExpectations(t)
		hash.AssertExpectations(t)
		au.AssertExpectations(t)

		assert.Nil(t, err)
		assert.Equal(t, expectedPersonalID, personalID)
		assert.Equal(t, expectedCompanyID, companyID)
	})

	t.Run("Error Registrasi Account Admin - Validation Error", func(t *testing.T) {
		accounts := users.PersonalDataEntity{
			Name:        "John Doe",
			Email:       "invalidemail",
			Password:    "short",
			PhoneNumber: "081234567890",
		}
		companyName := "Test Company"
		department := "IT"
		jobPosition := "Developer"

		au.On("EmailValidator", accounts.Email).Return(errors.New("invalid email")).Once()

		personalID, companyID, err := srv.RegistrasiAccountAdmin(accounts, companyName, department, jobPosition)

		au.AssertExpectations(t)

		assert.Error(t, err)
		assert.Equal(t, uint(0), personalID)
		assert.Equal(t, uint(0), companyID)
	})

	t.Run("Error Registrasi Account Admin - Empty Fields", func(t *testing.T) {
		accounts := users.PersonalDataEntity{
			Name:        "",
			Email:       "",
			Password:    "",
			PhoneNumber: "081234567890",
		}
		companyName := "Test Company"
		department := "IT"
		jobPosition := "Developer"

		personalID, companyID, err := srv.RegistrasiAccountAdmin(accounts, companyName, department, jobPosition)

		assert.Error(t, err)
		assert.Equal(t, "nama/email/password tidak boleh kosong", err.Error())
		assert.Equal(t, uint(0), personalID)
		assert.Equal(t, uint(0), companyID)
	})

	t.Run("Error Registrasi Account Admin - Hash Password Error", func(t *testing.T) {
		accounts := users.PersonalDataEntity{
			Name:        "John Doe",
			Email:       "johndoe@example.com",
			Password:    "password123",
			PhoneNumber: "081234567890",
		}
		companyName := "Test Company"
		department := "IT"
		jobPosition := "Developer"

		au.On("EmailValidator", accounts.Email).Return(nil).Once()
		au.On("PasswordValidator", accounts.Password).Return(nil).Once()
		au.On("PhoneNumberValidator", accounts.PhoneNumber).Return(nil).Once()
		hash.On("HashPassword", accounts.Password).Return("", errors.New("hash password error")).Once()

		personalID, companyID, err := srv.RegistrasiAccountAdmin(accounts, companyName, department, jobPosition)

		hash.AssertExpectations(t)
		au.AssertExpectations(t)

		assert.Error(t, err)
		assert.Equal(t, "hash password error", err.Error())
		assert.Equal(t, uint(0), personalID)
		assert.Equal(t, uint(0), companyID)
	})

	t.Run("Error Registrasi Account Admin - Create Account Error", func(t *testing.T) {
		accounts := users.PersonalDataEntity{
			Name:        "John Doe",
			Email:       "johndoe@example.com",
			Password:    "password123",
			PhoneNumber: "081234567890",
		}
		companyName := "Test Company"
		department := "IT"
		jobPosition := "Developer"
		hashedPassword := "hashedpassword"

		au.On("EmailValidator", accounts.Email).Return(nil).Once()
		au.On("PasswordValidator", accounts.Password).Return(nil).Once()
		au.On("PhoneNumberValidator", accounts.PhoneNumber).Return(nil).Once()
		hash.On("HashPassword", accounts.Password).Return(hashedPassword, nil).Once()

		qry.On("CreateAccountAdmin", mock.MatchedBy(func(a users.PersonalDataEntity) bool {
			return a.Name == accounts.Name && a.Email == accounts.Email && a.Password == hashedPassword && a.PhoneNumber == accounts.PhoneNumber
		}), companyName, department, jobPosition).Return(uint(0), uint(0), errors.New("create account error")).Once()

		personalID, companyID, err := srv.RegistrasiAccountAdmin(accounts, companyName, department, jobPosition)

		qry.AssertExpectations(t)
		hash.AssertExpectations(t)
		au.AssertExpectations(t)

		assert.Error(t, err)
		assert.Equal(t, "create account error", err.Error())
		assert.Equal(t, uint(0), personalID)
		assert.Equal(t, uint(0), companyID)
	})
	
	t.Run("Error Registrasi Account Admin - PhoneNumber Validation Error", func(t *testing.T) {
		accounts := users.PersonalDataEntity{
			Name:        "John Doe",
			Email:       "johndoe@example.com",
			Password:    "password123",
			PhoneNumber: "081234567890",
		}
		companyName := "Test Company"
		department := "IT"
		jobPosition := "Developer"

		au.On("EmailValidator", accounts.Email).Return(nil).Once()
		au.On("PasswordValidator", accounts.Password).Return(nil).Once()
		au.On("PhoneNumberValidator", accounts.PhoneNumber).Return(errors.New("invalid phone number")).Once()

		personalID, companyID, err := srv.RegistrasiAccountAdmin(accounts, companyName, department, jobPosition)

		au.AssertExpectations(t)

		assert.Error(t, err)
		assert.Equal(t, "invalid phone number", err.Error())
		assert.Equal(t, uint(0), personalID)
		assert.Equal(t, uint(0), companyID)
	})
}

func TestUpdateProfileAdmins(t *testing.T) {
	qry := mocks.NewDataUserInterface(t)
	hash := mocks.NewHashInterface(t)
	au := mocks.NewAccountUtilityInterface(t)
	srv := service.New(qry, hash, nil, au, nil)

	userID := uint(1)
	updateData := users.PersonalDataEntity{
		Email:       "johndoe@example.com",
		PhoneNumber: "081234567890",
		Password:    "newpassword",
	}

	t.Run("Success Update Profile Admins", func(t *testing.T) {
		hashedPassword := "hashedpassword"

		au.On("EmailValidator", updateData.Email).Return(nil).Once()
		au.On("PhoneNumberValidator", updateData.PhoneNumber).Return(nil).Once()
		hash.On("HashPassword", updateData.Password).Return(hashedPassword, nil).Once()

		qry.On("UpdateAccountAdmins", userID, mock.MatchedBy(func(data users.PersonalDataEntity) bool {
			return data.Email == updateData.Email && data.PhoneNumber == updateData.PhoneNumber && data.Password == hashedPassword
		})).Return(nil).Once()

		err := srv.UpdateProfileAdmins(userID, updateData)

		au.AssertExpectations(t)
		hash.AssertExpectations(t)
		qry.AssertExpectations(t)
		assert.Nil(t, err)
	})

	t.Run("Error Update Profile Admins - Invalid Email", func(t *testing.T) {
		au.On("EmailValidator", updateData.Email).Return(errors.New("invalid email")).Once()

		err := srv.UpdateProfileAdmins(userID, updateData)

		au.AssertExpectations(t)
		assert.Error(t, err)
		assert.Equal(t, "invalid email", err.Error())
	})

	t.Run("Error Update Profile Admins - Invalid Phone Number", func(t *testing.T) {
		au.On("EmailValidator", updateData.Email).Return(nil).Once()
		au.On("PhoneNumberValidator", updateData.PhoneNumber).Return(errors.New("invalid phone number")).Once()

		err := srv.UpdateProfileAdmins(userID, updateData)

		au.AssertExpectations(t)
		assert.Error(t, err)
		assert.Equal(t, "invalid phone number", err.Error())
	})

	t.Run("Error Update Profile Admins - Hash Password Error", func(t *testing.T) {
		au.On("EmailValidator", updateData.Email).Return(nil).Once()
		au.On("PhoneNumberValidator", updateData.PhoneNumber).Return(nil).Once()
		hash.On("HashPassword", updateData.Password).Return("", errors.New("hash password error")).Once()

		err := srv.UpdateProfileAdmins(userID, updateData)

		au.AssertExpectations(t)
		hash.AssertExpectations(t)
		assert.Error(t, err)
		assert.Equal(t, "hash password error", err.Error())
	})

	t.Run("Error Update Profile Admins - Database Error", func(t *testing.T) {
		hashedPassword := "hashedpassword"

		au.On("EmailValidator", updateData.Email).Return(nil).Once()
		au.On("PhoneNumberValidator", updateData.PhoneNumber).Return(nil).Once()
		hash.On("HashPassword", updateData.Password).Return(hashedPassword, nil).Once()

		qry.On("UpdateAccountAdmins", userID, mock.MatchedBy(func(data users.PersonalDataEntity) bool {
			return data.Email == updateData.Email && data.PhoneNumber == updateData.PhoneNumber && data.Password == hashedPassword
		})).Return(errors.New("database error")).Once()

		err := srv.UpdateProfileAdmins(userID, updateData)

		au.AssertExpectations(t)
		hash.AssertExpectations(t)
		qry.AssertExpectations(t)
		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
	})
}

func TestUpdateProfileEmployments(t *testing.T) {
	qry := mocks.NewDataUserInterface(t)
	srv := service.New(qry, nil, nil, nil, nil)

	userID := uint(1)
	updateData := users.EmploymentDataEntity{
		JobPosition: "Manager",
	}

	t.Run("Success Update Profile Employments", func(t *testing.T) {
		qry.On("UpdateProfileEmployments", userID, updateData).Return(nil).Once()

		err := srv.UpdateProfileEmployments(userID, updateData)

		qry.AssertExpectations(t)
		assert.Nil(t, err)
	})

	t.Run("Error Update Profile Employments - Invalid User ID", func(t *testing.T) {
		err := srv.UpdateProfileEmployments(0, updateData)

		assert.Error(t, err)
		assert.Equal(t, "invalid user ID", err.Error())
	})

	t.Run("Error Update Profile Employments - Database Error", func(t *testing.T) {
		qry.On("UpdateProfileEmployments", userID, updateData).Return(errors.New("database error")).Once()

		err := srv.UpdateProfileEmployments(userID, updateData)

		qry.AssertExpectations(t)
		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
	})
}

func TestUpdateProfileEmployees(t *testing.T) {
	qry := mocks.NewDataUserInterface(t)
	hash := mocks.NewHashInterface(t)
	au := mocks.NewAccountUtilityInterface(t)
	srv := service.New(qry, hash, nil, au, nil)

	userID := uint(1)
	updateData := users.PersonalDataEntity{
		Email:       "johndoe@example.com",
		PhoneNumber: "081234567890",
		Password:    "newpassword",
	}

	t.Run("Success Update Profile Employees", func(t *testing.T) {
		hashedPassword := "hashedpassword"

		au.On("EmailValidator", updateData.Email).Return(nil).Once()
		au.On("PhoneNumberValidator", updateData.PhoneNumber).Return(nil).Once()
		hash.On("HashPassword", updateData.Password).Return(hashedPassword, nil).Once()

		qry.On("UpdateAccountEmployees", userID, mock.MatchedBy(func(data users.PersonalDataEntity) bool {
			return data.Email == updateData.Email && data.PhoneNumber == updateData.PhoneNumber && data.Password == hashedPassword
		})).Return(nil).Once()

		err := srv.UpdateProfileEmployees(userID, updateData)

		au.AssertExpectations(t)
		hash.AssertExpectations(t)
		qry.AssertExpectations(t)
		assert.Nil(t, err)
	})

	t.Run("Error Update Profile Employees - Invalid Email", func(t *testing.T) {
		au.On("EmailValidator", updateData.Email).Return(errors.New("invalid email")).Once()

		err := srv.UpdateProfileEmployees(userID, updateData)

		au.AssertExpectations(t)
		assert.Error(t, err)
		assert.Equal(t, "invalid email", err.Error())
	})

	t.Run("Error Update Profile Employees - Invalid Phone Number", func(t *testing.T) {
		au.On("EmailValidator", updateData.Email).Return(nil).Once()
		au.On("PhoneNumberValidator", updateData.PhoneNumber).Return(errors.New("invalid phone number")).Once()

		err := srv.UpdateProfileEmployees(userID, updateData)

		au.AssertExpectations(t)
		assert.Error(t, err)
		assert.Equal(t, "invalid phone number", err.Error())
	})

	t.Run("Error Update Profile Employees - Hash Password Error", func(t *testing.T) {
		au.On("EmailValidator", updateData.Email).Return(nil).Once()
		au.On("PhoneNumberValidator", updateData.PhoneNumber).Return(nil).Once()
		hash.On("HashPassword", updateData.Password).Return("", errors.New("hash password error")).Once()

		err := srv.UpdateProfileEmployees(userID, updateData)

		au.AssertExpectations(t)
		hash.AssertExpectations(t)
		assert.Error(t, err)
		assert.Equal(t, "hash password error", err.Error())
	})

	t.Run("Error Update Profile Employees - Database Error", func(t *testing.T) {
		hashedPassword := "hashedpassword"

		au.On("EmailValidator", updateData.Email).Return(nil).Once()
		au.On("PhoneNumberValidator", updateData.PhoneNumber).Return(nil).Once()
		hash.On("HashPassword", updateData.Password).Return(hashedPassword, nil).Once()

		qry.On("UpdateAccountEmployees", userID, mock.MatchedBy(func(data users.PersonalDataEntity) bool {
			return data.Email == updateData.Email && data.PhoneNumber == updateData.PhoneNumber && data.Password == hashedPassword
		})).Return(errors.New("database error")).Once()

		err := srv.UpdateProfileEmployees(userID, updateData)

		au.AssertExpectations(t)
		hash.AssertExpectations(t)
		qry.AssertExpectations(t)
		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
	})
}

func TestGetAllAccount(t *testing.T) {
	qry := mocks.NewDataUserInterface(t)
	srv := service.New(qry, nil, nil, nil, nil)

	t.Run("Success Get All Account without filters", func(t *testing.T) {
		expectedResult := []users.PersonalDataEntity{
			{Name: "John Doe", Email: "johndoe@example.com"},
		}

		qry.On("GetAll", 1, 10).Return(expectedResult, nil).Once()

		result, err := srv.GetAllAccount("", "", 1, 10)

		qry.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("Success Get All Account by Name", func(t *testing.T) {
		name := "John"
		expectedResult := []users.PersonalDataEntity{
			{Name: "John Doe", Email: "johndoe@example.com"},
		}

		qry.On("GetAccountByName", name).Return(expectedResult, nil).Once()

		result, err := srv.GetAllAccount(name, "", 1, 10)

		qry.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("Success Get All Account by Job Level", func(t *testing.T) {
		jobLevel := "Manager"
		expectedResult := []users.PersonalDataEntity{
			{Name: "Jane Doe", Email: "janedoe@example.com"},
		}

		qry.On("GetAccountByJobLevel", jobLevel).Return(expectedResult, nil).Once()

		result, err := srv.GetAllAccount("", jobLevel, 1, 10)

		qry.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("Error Get All Account - Database Error", func(t *testing.T) {
		qry.On("GetAll", 1, 10).Return(nil, errors.New("database error")).Once()

		result, err := srv.GetAllAccount("", "", 1, 10)

		qry.AssertExpectations(t)
		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, result)
	})

	t.Run("Error Get All Account by Name - Database Error", func(t *testing.T) {
		name := "John"
		qry.On("GetAccountByName", name).Return(nil, errors.New("database error")).Once()

		result, err := srv.GetAllAccount(name, "", 1, 10)

		qry.AssertExpectations(t)
		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, result)
	})

	t.Run("Error Get All Account by Job Level - Database Error", func(t *testing.T) {
		jobLevel := "Manager"
		qry.On("GetAccountByJobLevel", jobLevel).Return(nil, errors.New("database error")).Once()

		result, err := srv.GetAllAccount("", jobLevel, 1, 10)

		qry.AssertExpectations(t)
		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, result)
	})
}

func TestUpdateEmploymentEmployee(t *testing.T) {
	qry := mocks.NewDataUserInterface(t)
	srv := service.New(qry, nil, nil, nil, nil)

	ID := uint(1)
	employeeID := uint(2)
	updateData := users.EmploymentDataEntity{
		JobPosition: "Manager",
	}

	t.Run("Success Update Employment Employee", func(t *testing.T) {
		qry.On("UpdateEmploymentEmployee", ID, employeeID, updateData).Return(nil).Once()

		err := srv.UpdateEmploymentEmployee(ID, employeeID, updateData)

		qry.AssertExpectations(t)
		assert.Nil(t, err)
	})

	t.Run("Error Update Employment Employee - Database Error", func(t *testing.T) {
		qry.On("UpdateEmploymentEmployee", ID, employeeID, updateData).Return(errors.New("database error")).Once()

		err := srv.UpdateEmploymentEmployee(ID, employeeID, updateData)

		qry.AssertExpectations(t)
		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
	})
}

func TestDashboard(t *testing.T) {
	qry := mocks.NewDataUserInterface(t)
	srv := service.New(qry, nil, nil, nil, nil)

	t.Run("Success Dashboard", func(t *testing.T) {
		companyID := uint(1)
		expectedResult := &users.DashboardStats{
			TotalUsers: 100,
		}

		qry.On("Dashboard", companyID).Return(expectedResult, nil).Once()

		result, err := srv.Dashboard(companyID)

		qry.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("Error Dashboard - Database Error", func(t *testing.T) {
		companyID := uint(1)

		qry.On("Dashboard", companyID).Return(nil, errors.New("database error")).Once()

		result, err := srv.Dashboard(companyID)

		qry.AssertExpectations(t)
		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, result)
	})
}