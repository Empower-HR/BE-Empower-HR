package service_test

import (
	"errors"
	"testing"

	companies "be-empower-hr/features/Companies"
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
		companyID := uint(1)
		existingUser := &users.PersonalDataEntity{
			CompanyID: companyID,
		}

		// Set up expectations
		qry.On("AccountById", userID).Return(existingUser, nil).Once()
		qry.On("DeleteAccountEmployeeByAdmin", userID).Return(nil).Once()

		err := srv.DeleteAccountEmployeeByAdmin(userID, companyID)

		// Assert expectations
		qry.AssertExpectations(t)
		assert.Nil(t, err)
	})

	t.Run("Error Fetching User Data", func(t *testing.T) {
		userID := uint(1)
		companyID := uint(1)

		// Set up expectation for AccountById to return an error
		qry.On("AccountById", userID).Return(nil, errors.New("fetch error")).Once()

		err := srv.DeleteAccountEmployeeByAdmin(userID, companyID)

		// Assert expectations
		qry.AssertExpectations(t)
		assert.Error(t, err)
		assert.ErrorContains(t, err, "fetch error")
	})

	t.Run("Error Deleting Account", func(t *testing.T) {
		userID := uint(1)
		companyID := uint(1)
		existingUser := &users.PersonalDataEntity{
			CompanyID: companyID,
		}

		// Set up expectations
		qry.On("AccountById", userID).Return(existingUser, nil).Once()
		qry.On("DeleteAccountEmployeeByAdmin", userID).Return(errors.New("delete error")).Once()

		err := srv.DeleteAccountEmployeeByAdmin(userID, companyID)

		// Assert expectations
		qry.AssertExpectations(t)
		assert.Error(t, err)
		assert.ErrorContains(t, err, "delete error")
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
	mockUserData := new(mocks.DataUserInterface)
	mockMiddleware := new(mocks.MiddlewaresInterface)
	userService := service.New(mockUserData, nil, mockMiddleware, nil, nil)

	t.Run("success", func(t *testing.T) {
		email := "johndoe@example.com"
		password := "password123"
		user := &users.PersonalDataEntity{
			PersonalDataID: 1,
			Email:          email,
			Password:       "hashedpassword",
			CompanyID:      1,
		}
		token := "sometoken"

		// Set up mocks
		mockUserData.On("AccountByEmail", email).Return(user, nil).Once()
		mockMiddleware.On("CreateToken", int(user.PersonalDataID), int(user.CompanyID)).Return(token, nil).Once()

		// Call the function
		returnedUser, returnedToken, err := userService.LoginAccount(email, password)

		// Assert results
		assert.NoError(t, err)
		assert.Equal(t, user, returnedUser)
		assert.Equal(t, token, returnedToken)
		mockUserData.AssertExpectations(t)
		mockMiddleware.AssertExpectations(t)
	})

	t.Run("account not found", func(t *testing.T) {
		email := "johndoe@example.com"
		password := "password123"

		// Set up mocks
		mockUserData.On("AccountByEmail", email).Return(nil, errors.New("account not found")).Once()

		// Call the function
		returnedUser, returnedToken, err := userService.LoginAccount(email, password)

		// Assert results
		assert.Error(t, err)
		assert.Equal(t, "account not found", err.Error())
		assert.Nil(t, returnedUser)
		assert.Empty(t, returnedToken)
		mockUserData.AssertExpectations(t)
	})

	t.Run("jwt authentication error", func(t *testing.T) {
		email := "johndoe@example.com"
		password := "password123"
		user := &users.PersonalDataEntity{
			PersonalDataID: 1,
			Email:          email,
			Password:       "hashedpassword",
			CompanyID:      1,
		}

		// Set up mocks
		mockUserData.On("AccountByEmail", email).Return(user, nil).Once()
		mockMiddleware.On("CreateToken", int(user.PersonalDataID), int(user.CompanyID)).Return("", errors.New("jwt error")).Once()

		// Call the function
		returnedUser, returnedToken, err := userService.LoginAccount(email, password)

		// Assert results
		assert.Error(t, err)
		assert.Equal(t, "jwt error", err.Error())
		assert.Nil(t, returnedUser)
		assert.Empty(t, returnedToken)
		mockUserData.AssertExpectations(t)
		mockMiddleware.AssertExpectations(t)
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

		qry.On("GetAll", 1, 10, uint(1)).Return(expectedResult, nil).Once()

		result, err := srv.GetAllAccount(1, "", "", 1, 10)

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

		result, err := srv.GetAllAccount(1, name, "", 1, 10)

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

		result, err := srv.GetAllAccount(1, "", jobLevel, 1, 10)

		qry.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("Error Get All Account - Database Error", func(t *testing.T) {
		qry.On("GetAll", 1, 10, uint(1)).Return(nil, errors.New("database error")).Once()

		result, err := srv.GetAllAccount(1, "", "", 1, 10)

		qry.AssertExpectations(t)
		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, result)
	})

	t.Run("Error Get All Account by Name - Database Error", func(t *testing.T) {
		name := "John"
		qry.On("GetAccountByName", name).Return(nil, errors.New("database error")).Once()

		result, err := srv.GetAllAccount(1, name, "", 1, 10)

		qry.AssertExpectations(t)
		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, result)
	})

	t.Run("Error Get All Account by Job Level - Database Error", func(t *testing.T) {
		jobLevel := "Manager"
		qry.On("GetAccountByJobLevel", jobLevel).Return(nil, errors.New("database error")).Once()

		result, err := srv.GetAllAccount(1, "", jobLevel, 1, 10)

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
		userID := uint(2)
		expectedResult := &users.DashboardStats{
			TotalUsers: 100,
		}

		qry.On("Dashboard", userID, companyID).Return(expectedResult, nil).Once()

		result, err := srv.Dashboard(userID, companyID)

		qry.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("Error Dashboard - Database Error", func(t *testing.T) {
		companyID := uint(1)
		userID := uint(2)

		qry.On("Dashboard", userID, companyID).Return(nil, errors.New("database error")).Once()

		result, err := srv.Dashboard(userID, companyID)

		qry.AssertExpectations(t)
		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		assert.Nil(t, result)
	})
}

func TestCreateNewEmployee(t *testing.T) {
	mockDataUser := new(mocks.DataUserInterface)
	mockHash := new(mocks.HashInterface)
	mockMiddleware := new(mocks.MiddlewaresInterface)
	mockAccountUtility := new(mocks.AccountUtilityInterface)
	mockCompany := new(mocks.Query)

	// Create service instance with mocked dependencies
	userService := service.New(mockDataUser, mockHash, mockMiddleware, mockAccountUtility, mockCompany)

	cmID := uint(1)
	addPersonal := users.PersonalDataEntity{
		Name:        "John Doe",
		Email:       "riandarmawan44@gmail.com",
		PhoneNumber: "1234567890",
		Gender:      "male",
		Religion:    "islam",
	}
	addEmployment := users.EmploymentDataEntity{
		EmploymentStatus: "permanent",
		JobLevel:         "staff",
	}
	addPayroll := users.PayrollDataEntity{}
	addLeaves := users.LeavesDataEntity{}

	companyData := companies.CompanyDataEntity{ID: 1}

	t.Run("success", func(t *testing.T) {
		mockCompany.On("GetCompanyID", cmID).Return(companyData, nil)
		mockAccountUtility.On("EmailValidator", addPersonal.Email).Return(nil)
		mockAccountUtility.On("PhoneNumberValidator", addPersonal.PhoneNumber).Return(nil)
		mockAccountUtility.On("GenderValidator", addPersonal.Gender).Return(nil)
		mockAccountUtility.On("ReligionValidator", addPersonal.Religion).Return(nil)
		mockAccountUtility.On("EmploymentStatusValidator", addEmployment.EmploymentStatus).Return(nil)
		mockAccountUtility.On("JobLevelValidator", addEmployment.JobLevel).Return(nil)

		addPersonal.Role = "employees" // Set the expected role value
		mockDataUser.On("CreatePersonal", companyData.ID, addPersonal).Return(uint(1), nil)

		mockDataUser.On("CreateEmployment", uint(1), addEmployment).Return(uint(2), nil)
		mockDataUser.On("CreatePayroll", uint(2), addPayroll).Return(nil)

		addLeaves.TotalLeave = 12 // Set the expected total leave value
		mockDataUser.On("CreateLeaves", uint(1), addLeaves).Return(uint(3), nil)

		mockAccountUtility.On("SendEmail", addPersonal.Email, mock.Anything, mock.Anything).Return(nil)

		// Call the function
		err := userService.CreateNewEmployee(cmID, addPersonal, addEmployment, addPayroll, addLeaves)

		// Assert results
		assert.NoError(t, err)
		mockCompany.AssertCalled(t, "GetCompanyID", cmID)
		mockAccountUtility.AssertExpectations(t)
		mockDataUser.AssertExpectations(t)
	})

	t.Run("empty name or email", func(t *testing.T) {
		// Test for empty name or email
		addPersonalEmpty := users.PersonalDataEntity{
			Name:  "",
			Email: "",
		}
		err := userService.CreateNewEmployee(cmID, addPersonalEmpty, addEmployment, addPayroll, addLeaves)
		assert.Error(t, err)
	})

	t.Run("email validation failed", func(t *testing.T) {
		mockCompany.On("GetCompanyID", cmID).Return(companyData, nil)
		mockAccountUtility.On("EmailValidator", addPersonal.Email).Return(errors.New("invalid email"))
		err := userService.CreateNewEmployee(cmID, addPersonal, addEmployment, addPayroll, addLeaves)
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid email")
		mockCompany.AssertCalled(t, "GetCompanyID", cmID)
		mockAccountUtility.AssertCalled(t, "EmailValidator", addPersonal.Email)
	})

	t.Run("phone number validation failed", func(t *testing.T) {
		mockCompany.On("GetCompanyID", cmID).Return(companyData, nil)
		mockAccountUtility.On("EmailValidator", addPersonal.Email).Return(nil)
		mockAccountUtility.On("PhoneNumberValidator", addPersonal.PhoneNumber).Return(errors.New("invalid phone number"))

		err := userService.CreateNewEmployee(cmID, addPersonal, addEmployment, addPayroll, addLeaves)
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid phone number")
		mockAccountUtility.AssertCalled(t, "PhoneNumberValidator", addPersonal.PhoneNumber)
	})

	t.Run("gender validation failed", func(t *testing.T) {
		mockCompany.On("GetCompanyID", cmID).Return(companyData, nil)
		mockAccountUtility.On("EmailValidator", addPersonal.Email).Return(nil)
		mockAccountUtility.On("PhoneNumberValidator", addPersonal.PhoneNumber).Return(nil)
		mockAccountUtility.On("GenderValidator", addPersonal.Gender).Return(errors.New("invalid gender"))

		err := userService.CreateNewEmployee(cmID, addPersonal, addEmployment, addPayroll, addLeaves)
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid gender")
		mockAccountUtility.AssertCalled(t, "GenderValidator", addPersonal.Gender)
	})

	t.Run("religion validation failed", func(t *testing.T) {
		mockCompany.On("GetCompanyID", cmID).Return(companyData, nil)
		mockAccountUtility.On("EmailValidator", addPersonal.Email).Return(nil)
		mockAccountUtility.On("PhoneNumberValidator", addPersonal.PhoneNumber).Return(nil)
		mockAccountUtility.On("GenderValidator", addPersonal.Gender).Return(nil)
		mockAccountUtility.On("ReligionValidator", addPersonal.Religion).Return(errors.New("invalid religion"))

		err := userService.CreateNewEmployee(cmID, addPersonal, addEmployment, addPayroll, addLeaves)
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid religion")
		mockAccountUtility.AssertCalled(t, "ReligionValidator", addPersonal.Religion)
	})

	t.Run("employment status validation failed", func(t *testing.T) {
		mockCompany.On("GetCompanyID", cmID).Return(companyData, nil)
		mockAccountUtility.On("EmailValidator", addPersonal.Email).Return(nil)
		mockAccountUtility.On("PhoneNumberValidator", addPersonal.PhoneNumber).Return(nil)
		mockAccountUtility.On("GenderValidator", addPersonal.Gender).Return(nil)
		mockAccountUtility.On("ReligionValidator", addPersonal.Religion).Return(nil)
		mockAccountUtility.On("EmploymentStatusValidator", addEmployment.EmploymentStatus).Return(errors.New("invalid employment status"))

		err := userService.CreateNewEmployee(cmID, addPersonal, addEmployment, addPayroll, addLeaves)
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid employment status")
		mockAccountUtility.AssertCalled(t, "EmploymentStatusValidator", addEmployment.EmploymentStatus)
	})

	t.Run("job level validation failed", func(t *testing.T) {
		mockCompany.On("GetCompanyID", cmID).Return(companyData, nil)
		mockAccountUtility.On("EmailValidator", addPersonal.Email).Return(nil)
		mockAccountUtility.On("PhoneNumberValidator", addPersonal.PhoneNumber).Return(nil)
		mockAccountUtility.On("GenderValidator", addPersonal.Gender).Return(nil)
		mockAccountUtility.On("ReligionValidator", addPersonal.Religion).Return(nil)
		mockAccountUtility.On("EmploymentStatusValidator", addEmployment.EmploymentStatus).Return(nil)
		mockAccountUtility.On("JobLevelValidator", addEmployment.JobLevel).Return(errors.New("invalid job level"))

		err := userService.CreateNewEmployee(cmID, addPersonal, addEmployment, addPayroll, addLeaves)
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid job level")
		mockAccountUtility.AssertCalled(t, "JobLevelValidator", addEmployment.JobLevel)
	})

	t.Run("fail to create personal data", func(t *testing.T) {
		mockAccountUtility.On("EmailValidator", addPersonal.Email).Return(nil)
		mockAccountUtility.On("PhoneNumberValidator", addPersonal.PhoneNumber).Return(nil)
		mockAccountUtility.On("GenderValidator", addPersonal.Gender).Return(nil)
		mockAccountUtility.On("ReligionValidator", addPersonal.Religion).Return(nil)
		mockAccountUtility.On("EmploymentStatusValidator", addEmployment.EmploymentStatus).Return(nil)
		mockAccountUtility.On("JobLevelValidator", addEmployment.JobLevel).Return(nil)
		mockCompany.On("GetCompanyID", cmID).Return(companyData, nil)

		addPersonal.Role = "employees" // Set the expected role value
		mockDataUser.On("CreatePersonal", companyData.ID, addPersonal).Return(uint(0), errors.New("failed to create personal data"))

		err := userService.CreateNewEmployee(cmID, addPersonal, addEmployment, addPayroll, addLeaves)

		assert.Error(t, err)
		assert.EqualError(t, err, "failed to create personal data")
		mockCompany.AssertCalled(t, "GetCompanyID", cmID)
		mockAccountUtility.AssertExpectations(t)
		mockDataUser.AssertExpectations(t)
	})

	t.Run("fail to create employment data", func(t *testing.T) {
		mockAccountUtility.On("EmailValidator", addPersonal.Email).Return(nil)
		mockAccountUtility.On("PhoneNumberValidator", addPersonal.PhoneNumber).Return(nil)
		mockAccountUtility.On("GenderValidator", addPersonal.Gender).Return(nil)
		mockAccountUtility.On("ReligionValidator", addPersonal.Religion).Return(nil)
		mockAccountUtility.On("EmploymentStatusValidator", addEmployment.EmploymentStatus).Return(nil)
		mockAccountUtility.On("JobLevelValidator", addEmployment.JobLevel).Return(nil)
		mockCompany.On("GetCompanyID", cmID).Return(companyData, nil)

		addPersonal.Role = "employees" // Set the expected role value
		mockDataUser.On("CreatePersonal", companyData.ID, addPersonal).Return(uint(1), nil)
		mockDataUser.On("CreateEmployment", uint(1), addEmployment).Return(uint(0), errors.New("failed to create employment data"))

		err := userService.CreateNewEmployee(cmID, addPersonal, addEmployment, addPayroll, addLeaves)

		assert.Error(t, err)
		assert.EqualError(t, err, "failed to create employment data")
		mockCompany.AssertCalled(t, "GetCompanyID", cmID)
		mockAccountUtility.AssertExpectations(t)
		mockDataUser.AssertExpectations(t)
	})

	t.Run("fail to create payroll data", func(t *testing.T) {
		mockAccountUtility.On("EmailValidator", addPersonal.Email).Return(nil)
		mockAccountUtility.On("PhoneNumberValidator", addPersonal.PhoneNumber).Return(nil)
		mockAccountUtility.On("GenderValidator", addPersonal.Gender).Return(nil)
		mockAccountUtility.On("ReligionValidator", addPersonal.Religion).Return(nil)
		mockAccountUtility.On("EmploymentStatusValidator", addEmployment.EmploymentStatus).Return(nil)
		mockAccountUtility.On("JobLevelValidator", addEmployment.JobLevel).Return(nil)
		mockCompany.On("GetCompanyID", cmID).Return(companyData, nil)

		addPersonal.Role = "employees" // Set the expected role value
		mockDataUser.On("CreatePersonal", companyData.ID, addPersonal).Return(uint(1), nil)
		mockDataUser.On("CreateEmployment", uint(1), addEmployment).Return(uint(2), nil)
		mockDataUser.On("CreatePayroll", uint(2), addPayroll).Return(errors.New("failed to create payroll data"))

		err := userService.CreateNewEmployee(cmID, addPersonal, addEmployment, addPayroll, addLeaves)

		assert.Error(t, err)
		assert.EqualError(t, err, "failed to create payroll data")
		mockCompany.AssertCalled(t, "GetCompanyID", cmID)
		mockAccountUtility.AssertExpectations(t)
		mockDataUser.AssertExpectations(t)
	})

	t.Run("fail to create leaves data", func(t *testing.T) {
		mockAccountUtility.On("EmailValidator", addPersonal.Email).Return(nil)
		mockAccountUtility.On("PhoneNumberValidator", addPersonal.PhoneNumber).Return(nil)
		mockAccountUtility.On("GenderValidator", addPersonal.Gender).Return(nil)
		mockAccountUtility.On("ReligionValidator", addPersonal.Religion).Return(nil)
		mockAccountUtility.On("EmploymentStatusValidator", addEmployment.EmploymentStatus).Return(nil)
		mockAccountUtility.On("JobLevelValidator", addEmployment.JobLevel).Return(nil)
		mockCompany.On("GetCompanyID", cmID).Return(companyData, nil)

		addPersonal.Role = "employees" // Set the expected role value
		mockDataUser.On("CreatePersonal", companyData.ID, addPersonal).Return(uint(1), nil)
		mockDataUser.On("CreateEmployment", uint(1), addEmployment).Return(uint(2), nil)
		mockDataUser.On("CreatePayroll", uint(2), addPayroll).Return(nil)
		addLeaves.TotalLeave = 12 // Set the expected total leave value
		mockDataUser.On("CreateLeaves", uint(1), addLeaves).Return(uint(0), errors.New("failed to create leaves data"))

		err := userService.CreateNewEmployee(cmID, addPersonal, addEmployment, addPayroll, addLeaves)

		assert.Error(t, err)
		assert.EqualError(t, err, "failed to create leaves data")
		mockCompany.AssertCalled(t, "GetCompanyID", cmID)
		mockAccountUtility.AssertExpectations(t)
		mockDataUser.AssertExpectations(t)
	})

	t.Run("fail to send email", func(t *testing.T) {
		mockAccountUtility.On("EmailValidator", addPersonal.Email).Return(nil)
		mockAccountUtility.On("PhoneNumberValidator", addPersonal.PhoneNumber).Return(nil)
		mockAccountUtility.On("GenderValidator", addPersonal.Gender).Return(nil)
		mockAccountUtility.On("ReligionValidator", addPersonal.Religion).Return(nil)
		mockAccountUtility.On("EmploymentStatusValidator", addEmployment.EmploymentStatus).Return(nil)
		mockAccountUtility.On("JobLevelValidator", addEmployment.JobLevel).Return(nil)
		mockCompany.On("GetCompanyID", cmID).Return(companyData, nil)

		addPersonal.Role = "employees" // Set the expected role value
		mockDataUser.On("CreatePersonal", companyData.ID, addPersonal).Return(uint(1), nil)
		mockDataUser.On("CreateEmployment", uint(1), addEmployment).Return(uint(2), nil)
		mockDataUser.On("CreatePayroll", uint(2), addPayroll).Return(nil)
		addLeaves.TotalLeave = 12 // Set the expected total leave value
		mockDataUser.On("CreateLeaves", uint(1), addLeaves).Return(uint(3), nil)
		mockAccountUtility.On("SendEmail", addPersonal.Email, mock.Anything, mock.Anything).Return(errors.New("failed to send email"))

		err := userService.CreateNewEmployee(cmID, addPersonal, addEmployment, addPayroll, addLeaves)

		assert.Error(t, err)
		assert.EqualError(t, err, "failed to send email")
		mockCompany.AssertCalled(t, "GetCompanyID", cmID)
		mockAccountUtility.AssertExpectations(t)
		mockDataUser.AssertExpectations(t)
	})
}
