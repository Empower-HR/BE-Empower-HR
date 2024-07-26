package handler

import (
	"be-empower-hr/app/middlewares"
	users "be-empower-hr/features/Users"
	"be-empower-hr/utils/responses"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService users.ServiceUserInterface
}

func New(us users.ServiceUserInterface) *UserHandler {
	return &UserHandler{
		userService: us,
	}
}

func (uh *UserHandler) RegisterAdmin(c echo.Context) error {
	newUser := UserRequest{}
	if errBind := c.Bind(&newUser); errBind != nil {
		log.Printf("Register: Error binding data: %v", errBind)
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "error binding data: "+errBind.Error(), nil))
	}

	dataUser := users.PersonalDataEntity{
		Name:        newUser.PersonalData.Name,
		Email:       newUser.PersonalData.Email,
		Password:    newUser.PersonalData.Password,
		PhoneNumber: newUser.PersonalData.PhoneNumber,
		Role:        "admin",
	}

	// Menggunakan data dari EmploymentData yang pertama, jika ada
	empData := newUser.PersonalData.EmploymentData[0]
	_, errInsert := uh.userService.RegistrasiAccountAdmin(dataUser, newUser.CompanyData.CompanyName, empData.Department, empData.JobPosition)
	if errInsert != nil {
		log.Printf("Register: Error registering user: %v", errInsert)
		if strings.Contains(errInsert.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "failed", "user registration failed: "+errInsert.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "user registration failed: "+errInsert.Error(), nil))
	}

	userResponse := UserResponse{
		PersonalData: newUser.PersonalData,
		CompanyName:  newUser.CompanyData.CompanyName,
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse(http.StatusCreated, "success", "user registration successful", userResponse))
}

// func (uh *UserHandler) RegisterEmployee(c echo.Context) error {

// }

func (uh *UserHandler) Login(c echo.Context) error {
	loginReq := LoginRequest{}
	if errBind := c.Bind(&loginReq); errBind != nil {
		log.Printf("Login: Error binding data: %v", errBind)
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "failed", "error binding data: "+errBind.Error(), nil))
	}

	data, token, err := uh.userService.LoginAccount(loginReq.Email, loginReq.Password)
	if err != nil {
		log.Printf("Login: User login failed: %v", err)
		if strings.Contains(err.Error(), "email atau password tidak sesuai") {
			return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "user login failed: "+err.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "user login failed: "+err.Error(), nil))
	}

	loginResponse := struct {
		ID       uint   `json:"id"`
		FullName string `json:"fullname"`
		Role     string `json:"role"`
		Token    string `json:"token"`
	}{
		ID:       data.PersonalDataID,
		FullName: data.Name,
		Role:     data.Role,
		Token:    token,
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "Login successfully", loginResponse))
}

func (uh *UserHandler) GetProfile(c echo.Context) error {
	userID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if userID == 0 {
		log.Println("Invalid user ID from token")
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "invalid token", nil))
	}

	data, err := uh.userService.GetProfile(uint(userID))
	if err != nil {
		log.Printf("Error getting user profile: %v", err)
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "error getting user profile", nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "profile retrieved successfully", data))
}

func (uh *UserHandler) UpdateProfileAdmins(c echo.Context) error {
	userID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "unauthorized", nil))
	}

	updatedUser := UserRequest{}
	if errBind := c.Bind(&updatedUser); errBind != nil {
		log.Printf("UpdateProfileAdmins: Error binding data: %v", errBind)
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "error binding data: "+errBind.Error(), nil))
	}

	dataUser := users.PersonalDataEntity{
		ProfilePicture: updatedUser.PersonalData.ProfilePicture,
		Name:           updatedUser.PersonalData.Name,
		Email:          updatedUser.PersonalData.Email,
		Password:       updatedUser.PersonalData.Password,
		PhoneNumber:    updatedUser.PersonalData.PhoneNumber,
		PlaceBirth:     updatedUser.PersonalData.PlaceBirth,
		BirthDate:      updatedUser.PersonalData.BirthDate,
		Gender:         updatedUser.PersonalData.Gender,
		Religion:       updatedUser.PersonalData.Religion,
		NIK:            updatedUser.PersonalData.NIK,
		Address:        updatedUser.PersonalData.Address,
		EmploymentData: []users.EmploymentDataEntity{
			{
				EmploymentStatus: updatedUser.PersonalData.EmploymentData[0].EmploymentStatus,
				JoinDate:         updatedUser.PersonalData.EmploymentData[0].JoinDate,
				Department:       updatedUser.PersonalData.EmploymentData[0].Department,
				JobPosition:      updatedUser.PersonalData.EmploymentData[0].JobPosition,
				JobLevel:         updatedUser.PersonalData.EmploymentData[0].JobLevel,
				Schedule:         updatedUser.PersonalData.EmploymentData[0].Schedule,
				ApprovalLine:     updatedUser.PersonalData.EmploymentData[0].ApprovalLine,
				Manager:          updatedUser.PersonalData.EmploymentData[0].Manager,
			},
		},
	}

	err := uh.userService.UpdateProfileAdmins(uint(userID), dataUser)
	if err != nil {
		log.Printf("UpdateProfileAdmins: Error updating profile: %v", err)
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "error updating profile: "+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "profile updated successfully", nil))
}

func (uh *UserHandler) DeleteAccountAdmin(c echo.Context) error {
	userID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if userID == 0 {
		log.Println("Invalid user ID from token")
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "invalid token", nil))
	}

	err := uh.userService.DeleteAccountAdmin(uint(userID))
	if err != nil {
		log.Printf("Error deleting admin account: %v", err)
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "error deleting admin account", nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "account deleted successfully", nil))
}

func (uh *UserHandler) UpdateProfileEmployees(c echo.Context) error {
	userID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "unauthorized", nil))
	}

	updatedUser := UserRequest{}
	if errBind := c.Bind(&updatedUser); errBind != nil {
		log.Printf("update profile employees: Error binding data: %v", errBind)
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "error binding data: "+errBind.Error(), nil))
	}

	dataUser := users.PersonalDataEntity{
		ProfilePicture: updatedUser.PersonalData.ProfilePicture,
		Name:           updatedUser.PersonalData.Name,
		Email:          updatedUser.PersonalData.Email,
		Password:       updatedUser.PersonalData.Password,
		PhoneNumber:    updatedUser.PersonalData.PhoneNumber,
		PlaceBirth:     updatedUser.PersonalData.PlaceBirth,
		BirthDate:      updatedUser.PersonalData.BirthDate,
		Gender:         updatedUser.PersonalData.Gender,
		Religion:       updatedUser.PersonalData.Religion,
		NIK:            updatedUser.PersonalData.NIK,
		Address:        updatedUser.PersonalData.Address,
	}

	err := uh.userService.UpdateProfileEmployees(uint(userID), dataUser)
	if err != nil {
		log.Printf("update profile employees: Error updating profile: %v", err)
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "error updating profile: "+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "profile updated successfully", nil))
}

func (uh *UserHandler) DeleteAccountEmployees(c echo.Context) error {
	userID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if userID == 0 {
		log.Println("Invalid user ID from token")
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "invalid token", nil))
	}

	err := uh.userService.DeleteAccountEmployee(uint(userID))
	if err != nil {
		log.Printf("Error deleting employees account: %v", err)
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "error deleting employees account", nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "account deleted successfully", nil))
}
