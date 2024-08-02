package handler

import (
	"be-empower-hr/app/middlewares"
	users "be-empower-hr/features/Users"
	"be-empower-hr/utils/cloudinary"
	"be-empower-hr/utils/responses"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService       users.ServiceUserInterface
	cloudinaryUtility cloudinary.CloudinaryUtilityInterface
}

func New(us users.ServiceUserInterface, cu cloudinary.CloudinaryUtilityInterface) *UserHandler {
	return &UserHandler{
		userService:       us,
		cloudinaryUtility: cu,
	}
}

func (uh *UserHandler) RegisterAdmin(c echo.Context) error {
	newUser := UserRequest{}
	if errBind := c.Bind(&newUser); errBind != nil {
		log.Printf("register: Error binding data: %v", errBind)
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "error binding data: "+errBind.Error(), nil))
	}

	dataUser := users.PersonalDataEntity{
		Name:        newUser.Name,
		Email:       newUser.WorkEmail,
		Password:    newUser.Password,
		PhoneNumber: newUser.PhoneNumber,
		Role:        "admin",
	}

	empData := users.EmploymentDataEntity{
		Department:  newUser.Department,
		JobPosition: newUser.JobPosition,
	}

	_, _, errInsert := uh.userService.RegistrasiAccountAdmin(dataUser, newUser.Company, empData.Department, empData.JobPosition)
	if errInsert != nil {
		log.Printf("register: error registering user: %v", errInsert)
		if strings.Contains(errInsert.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "failed", "user registration failed: "+errInsert.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "user registration failed: "+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse(http.StatusCreated, "success", "user registration successful", nil))
}

func (uh *UserHandler) Login(c echo.Context) error {
	loginReq := LoginRequest{}
	if errBind := c.Bind(&loginReq); errBind != nil {
		log.Printf("login: Error binding data: %v", errBind)
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "failed", "error binding data: "+errBind.Error(), nil))
	}

	data, token, err := uh.userService.LoginAccount(loginReq.Email, loginReq.Password)
	if err != nil {
		log.Printf("login: User login failed: %v", err)
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
		log.Println("invalid user ID from token")
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "invalid token", nil))
	}

	profile, err := uh.userService.GetProfile(uint(userID))
	if err != nil {
		log.Printf("error getting user profile: %v", err)
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "error getting user profile", nil))
	}

	profileResponse := ProfileResponse{
		ProfilePicture: profile.ProfilePicture,
		Name:           profile.Name,
		Email:          profile.Email,
		PhoneNumber:    profile.PhoneNumber,
		PlaceBirthDate: profile.PlaceBirth,
		BirthDate:      profile.BirthDate,
		Gender:         profile.Gender,
		Religion:       profile.Religion,
		NIK:            profile.NIK,
		Address:        profile.Address,
		EmploymentData: make([]EmploymentDataResponse, len(profile.EmploymentData)),
	}

	for i, emp := range profile.EmploymentData {
		profileResponse.EmploymentData[i] = EmploymentDataResponse{
			EmploymentStatus: emp.EmploymentStatus,
			JoinDate:         emp.JoinDate,
			Department:       emp.Department,
			JobPosition:      emp.JobPosition,
			JobLevel:         emp.JobLevel,
			Schedule:         emp.Schedule,
			ApprovalLine:     emp.ApprovalLine,
		}
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "profile retrieved successfully", profileResponse))
}

// lihat detail profile employee
func (uh *UserHandler) GetProfileById(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "failed", "error converting id: "+errConv.Error(), nil))
	}

	profile, err := uh.userService.GetProfileById(uint(idConv))
	if err != nil {
		log.Printf("error getting user profile: %v", err)
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "error getting user profile", nil))
	}

	profileResponse := ProfileResponse{
		ProfilePicture: profile.ProfilePicture,
		Name:           profile.Name,
		Email:          profile.Email,
		PhoneNumber:    profile.PhoneNumber,
		PlaceBirthDate: profile.PlaceBirth,
		BirthDate:      profile.BirthDate,
		Gender:         profile.Gender,
		Religion:       profile.Religion,
		NIK:            profile.NIK,
		Address:        profile.Address,
		EmploymentData: make([]EmploymentDataResponse, len(profile.EmploymentData)),
	}

	for i, emp := range profile.EmploymentData {
		profileResponse.EmploymentData[i] = EmploymentDataResponse{
			EmploymentStatus: emp.EmploymentStatus,
			JoinDate:         emp.JoinDate,
			Department:       emp.Department,
			JobPosition:      emp.JobPosition,
			JobLevel:         emp.JobLevel,
			Schedule:         emp.Schedule,
			ApprovalLine:     emp.ApprovalLine,
		}
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "profile retrieved successfully", profileResponse))
}

func (uh *UserHandler) UpdateProfileAdmins(c echo.Context) error {
	userID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "unauthorized", nil))
	}

	updatedUser := UpdateAdminRequest{}
	if errBind := c.Bind(&updatedUser); errBind != nil {
		log.Printf("update profile admin: Error binding data: %v", errBind)
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "error binding data: "+errBind.Error(), nil))
	}

	// Handle profile picture upload to Cloudinary
	profilePictureFile, err := c.FormFile("profile_picture")
	if err == nil {
		src, err := profilePictureFile.Open()
		if err != nil {
			log.Printf("update profile admins: Error opening file: %v", err)
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "error opening file: "+err.Error(), nil))
		}
		defer src.Close()

		profilePictureURL, err := uh.cloudinaryUtility.UploadCloudinary(src, profilePictureFile.Filename)
		if err != nil {
			log.Printf("update profile admin: Error uploading to Cloudinary: %v", err)
			return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "error uploading to Cloudinary: "+err.Error(), nil))
		}
		updatedUser.ProfilePicture = profilePictureURL
	}

	dataUser := users.PersonalDataEntity{
		ProfilePicture: updatedUser.ProfilePicture,
		Name:           updatedUser.Name,
		Email:          updatedUser.Email,
		Password:       updatedUser.Password,
		PhoneNumber:    updatedUser.PhoneNumber,
		PlaceBirth:     updatedUser.PlaceBirth,
		BirthDate:      updatedUser.BirthDate,
		Gender:         updatedUser.Gender,
		Religion:       updatedUser.Religion,
		NIK:            updatedUser.NIK,
		Address:        updatedUser.Address,
	}

	err = uh.userService.UpdateProfileAdmins(uint(userID), dataUser)
	if err != nil {
		log.Printf("update profile admin: Error updating profile: %v", err)
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "error updating profile: "+err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse(http.StatusCreated, "success", "profile updated successfully", nil))
}

func (uh *UserHandler) UpdateProfileEmployment(c echo.Context) error {
	userID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "unauthorized", nil))
	}

	updatedUser := EmploymentData{}
	if errBind := c.Bind(&updatedUser); errBind != nil {
		log.Printf("update profile admin: Error binding data: %v", errBind)
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "error binding data: "+errBind.Error(), nil))
	}

	dataUser := users.EmploymentDataEntity{
		EmploymentStatus: updatedUser.EmploymentStatus,
		JoinDate:         updatedUser.JoinDate,
		Department:       updatedUser.Department,
		JobPosition:      updatedUser.JobPosition,
		JobLevel:         updatedUser.JobLevel,
		Schedule:         updatedUser.Schedule,
		ApprovalLine:     updatedUser.ApprovalLine,
	}

	err := uh.userService.UpdateProfileEmployments(uint(userID), dataUser)
	if err != nil {
		log.Printf("update profile admin: Error updating profile: %v", err)
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "error updating profile: "+err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse(http.StatusCreated, "success", "profile updated successfully", nil))
}

func (uh *UserHandler) UpdateProfileEmploymentByAdmin(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error convert id: " + errConv.Error(),
		})
	}

	updatedUser := EmploymentData{}
	if errBind := c.Bind(&updatedUser); errBind != nil {
		log.Printf("update profile admin: Error binding data: %v", errBind)
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "error binding data: "+errBind.Error(), nil))
	}

	dataUser := users.EmploymentDataEntity{
		EmploymentStatus: updatedUser.EmploymentStatus,
		JoinDate:         updatedUser.JoinDate,
		Department:       updatedUser.Department,
		JobPosition:      updatedUser.JobPosition,
		JobLevel:         updatedUser.JobLevel,
		Schedule:         updatedUser.Schedule,
		ApprovalLine:     updatedUser.ApprovalLine,
	}

	err := uh.userService.UpdateProfileEmployments(uint(idConv), dataUser)
	if err != nil {
		log.Printf("update profile admin: Error updating profile: %v", err)
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "error updating profile: "+err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse(http.StatusCreated, "success", "profile updated successfully", nil))
}

func (uh *UserHandler) DeleteAccountAdmin(c echo.Context) error {
	userID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if userID == 0 {
		log.Println("invalid user ID from token")
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "invalid token", nil))
	}

	err := uh.userService.DeleteAccountAdmin(uint(userID))
	if err != nil {
		log.Printf("error deleting admin account: %v", err)
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "error deleting admin account", nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "account deleted successfully", nil))
}

func (uh *UserHandler) UpdateProfileEmployees(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
	return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "failed", "invalid ID", nil))
	}
	
	updatedUser := UpdateAdminRequest{}
	if errBind := c.Bind(&updatedUser); errBind != nil {
	log.Printf("update profile employees: Error binding data: %v", errBind)
	return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "error binding data: "+errBind.Error(), nil))
	}
	
	// Handle profile picture upload to Cloudinary
	profilePictureFile, err := c.FormFile("profile_picture")
	if err == nil {
	src, err := profilePictureFile.Open()
	if err != nil {
	log.Printf("update profile employees: Error opening file: %v", err)
	return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "error opening file: "+err.Error(), nil))
	}
	defer src.Close()
	
	profilePictureURL, err := uh.cloudinaryUtility.UploadCloudinary(src, profilePictureFile.Filename)
	if err != nil {
	log.Printf("update profile employees: Error uploading to Cloudinary: %v", err)
	return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "error uploading to Cloudinary: "+err.Error(), nil))
	}
	updatedUser.ProfilePicture = profilePictureURL
	}
	
	dataUser := users.PersonalDataEntity{
	ProfilePicture: updatedUser.ProfilePicture,
	Name:           updatedUser.Name,
	Email:          updatedUser.Email,
	Password:       updatedUser.Password,
	PhoneNumber:    updatedUser.PhoneNumber,
	PlaceBirth:     updatedUser.PlaceBirth,
	BirthDate:      updatedUser.BirthDate,
	Gender:         updatedUser.Gender,
	Religion:       updatedUser.Religion,
	NIK:            updatedUser.NIK,
	Address:        updatedUser.Address,
	}
	
	err = uh.userService.UpdateProfileEmployees(uint(id), dataUser)
	if err != nil {
	log.Printf("update profile employees: Error updating profile: %v", err)
	return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "error updating profile: "+err.Error(), nil))
	}
	
	return c.JSON(http.StatusCreated, responses.JSONWebResponse(http.StatusCreated, "success", "profile updated successfully", nil))
	}

func (uh *UserHandler) DeleteAccountEmployees(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error convert id: " + errConv.Error(),
		})
	}

	companyID, err := middlewares.NewMiddlewares().ExtractCompanyID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"status":  "failed",
			"message": "unauthorized: " + err.Error(),
		})
	}

	err = uh.userService.DeleteAccountEmployeeByAdmin(uint(idConv), companyID)
	if err != nil {
		log.Printf("error deleting employees account: %v", err)
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "error deleting employees account", nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "account deleted successfully", nil))
}

func (uh *UserHandler) GetAllAccount(c echo.Context) error {
	userID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "unauthorized", nil))
	}
	companyID, err := middlewares.NewMiddlewares().ExtractCompanyID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"status":  "failed",
			"message": "unauthorized: " + err.Error(),
		})
	}
	name := c.QueryParam("name")
	jobLevel := c.QueryParam("job_level")
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1
	}
	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil {
		pageSize = 10
	}

	allAccount, err := uh.userService.GetAllAccount(companyID, name, jobLevel, page, pageSize)
	if err != nil {
		log.Println("error fetching accounts:", err)
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "error", "Failed to fetch accounts", nil))
	}

	var allUserResponse []AllUsersResponse
	for _, v := range allAccount {
		for _, emp := range v.EmploymentData {
			allUserResponse = append(allUserResponse, AllUsersResponse{
				PersonalDataID:   emp.EmploymentDataID,
				Name:             v.Name,
				JobPosition:      emp.JobPosition,
				JobLevel:         emp.JobLevel,
				EmploymentStatus: emp.EmploymentStatus,
				JoinDate:         emp.JoinDate,
			})
		}
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "All accounts fetched successfully", allUserResponse))
}

// handler update employment employee
func (uh *UserHandler) UpdateEmploymentEmployee(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "failed", "error converting id: "+errConv.Error(), nil))
	}

	employeeID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if employeeID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "unauthorized", nil))
	}

	var input EmploymentData

	if errBind := c.Bind(&input); errBind != nil {
		log.Printf("update profile employees: Error binding data: %v", errBind)
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "error binding data: "+errBind.Error(), nil))
	}

	err := uh.userService.UpdateEmploymentEmployee(uint(idConv), uint(employeeID), ToModelEmploymentData(input))
	if err != nil {
		log.Printf("update profile employees: Error updating profile: %v", err)
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "error updating profile: "+err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse(http.StatusCreated, "success", "profile updated successfully", nil))
}

func (uh *UserHandler) CreateNewEmployee(c echo.Context) error {
	companyID, err := middlewares.NewMiddlewares().ExtractCompanyID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"status":  "failed",
			"message": "unauthorized: " + err.Error(),
		})
	}
	var newEmployeeRequeste NewEmployeeRequest
	if errNewEmployeReq := c.Bind(&newEmployeeRequeste); errNewEmployeReq != nil {
		log.Printf("update profile employees: Error binding data: %v", errNewEmployeReq)
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "error binding data: "+errNewEmployeReq.Error(), nil))
	}

	err = uh.userService.CreateNewEmployee(
		companyID,
		ToModelPersonalData(newEmployeeRequeste.PersonalData),
		ToModelEmploymentData(newEmployeeRequeste.EmploymentData),
		ToModelPayroll(newEmployeeRequeste.Payroll),
		ToModelLeaves(newEmployeeRequeste.Leaves),
	)

	if err != nil {
		log.Printf("update profile employees: Error create employee data: %v", err)
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "error create employee data: "+err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse(http.StatusCreated, "success", "Create new employee successfully", nil))
}
func (uh *UserHandler) DasboardAdmin(c echo.Context) error {
	userID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if userID == 0 {
		log.Println("invalid user ID from token")
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "invalid token", nil))
	}

	companyID, err := middlewares.NewMiddlewares().ExtractCompanyID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "unauthorized: "+err.Error(), nil))
	}

	dashboardData, err := uh.userService.Dashboard(companyID)
	if err != nil {
		log.Printf("error fetching dashboard data: %v", err)
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "error", "failed to fetch dashboard data", nil))
	}

	responseData := DashboardStatsResponses{
		TotalUsers:               dashboardData.TotalUsers,
		MalePercentage:           dashboardData.MalePercentage,
		FemalePercentage:         dashboardData.FemalePercentage,
		ContractUsersPercentage:  dashboardData.ContractUsersPercentage,
		PermanentUsersPercentage: dashboardData.PermanentUsersPercentage,
		PayrollRecords:           dashboardData.PayrollRecords,
		LeavesRecords:            dashboardData.LeavesPending,
		PersonalDataName:         dashboardData.PersonalDataNames,
		AttendanceRecords:        dashboardData.AttendanceHadir,
		CurrentDate:              dashboardData.CurrentDate,
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "Dashboard data retrieved successfully", responseData))
}

func getCompanyIDFromUserID(userID int) (uint, error) {
	return 1, nil
}

func (uh *UserHandler) DashboardEmployees(c echo.Context) error {
	userID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if userID == 0 {
		log.Println("invalid user ID from token")
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "invalid token", nil))
	}

	companyID, err := getCompanyIDFromUserID(userID)
	if err != nil {
		log.Printf("Error retrieving company ID: %v", err)
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "error", "Failed to retrieve company ID", nil))
	}

	dashboardData, err := uh.userService.Dashboard(companyID)
	if err != nil {
		log.Printf("error fetching dashboard data: %v", err)
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "error", "failed to fetch dashboard data", nil))
	}

	responseData := DashboardStatsResponses{
		PersonalDataName:         dashboardData.PersonalDataNames,
		MalePercentage:           dashboardData.MalePercentage,
		FemalePercentage:         dashboardData.FemalePercentage,
		ContractUsersPercentage:  dashboardData.ContractUsersPercentage,
		PermanentUsersPercentage: dashboardData.PermanentUsersPercentage,
		CurrentDate:              dashboardData.CurrentDate,
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "Dashboard data retrieved successfully", responseData))
}
