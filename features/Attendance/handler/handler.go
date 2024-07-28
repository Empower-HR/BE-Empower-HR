package handler

import (
	"be-empower-hr/app/middlewares"
	attendance "be-empower-hr/features/Attendance"
	"be-empower-hr/utils"
	"be-empower-hr/utils/responses"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AttHandler struct {
	srv attendance.AServices
}

func New(as attendance.AServices) attendance.AHandler {
	return &AttHandler{
		srv: as,
	}
}

// add Data Absen
func (ah *AttHandler) AddAttendance(c echo.Context) error {
	userID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "unauthorized", nil))
	}

	fmt.Println("user-id :", userID)
	AttRequest := AttRequest{}
	if errBind := c.Bind(&AttRequest); errBind != nil {
		log.Printf("Add Attendances: Error binding data: %v", errBind)
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "error binding data: "+errBind.Error(), nil))
	}

	dataAtt := attendance.Attandance{
		PersonalDataID: uint(userID),
		Clock_in:       AttRequest.Clock_in,
		Date:           AttRequest.Date,
		Long:           AttRequest.Long,
		Lat:            AttRequest.Lat,
		Notes:          AttRequest.Notes,
	}

	// panggil fungsi addAtt pada service
	err := ah.srv.AddAtt(dataAtt)
	if err != nil {
		log.Printf("Add Attandance: Error add Attendance: %v", err)
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "error add data: "+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "clock in successfully", nil))
}

// Update by id attendance
func (ah *AttHandler) UpdateAttendance(c echo.Context) error {
	attID := c.Param("attendance_id")

	attId, err := strconv.ParseUint(attID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "failed", "Invalid Attendance ID", nil))
	}
	userID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "unauthorized", nil))
	}

	AttRequest := AttUpdateReq{}
	if errBind := c.Bind(&AttRequest); errBind != nil {
		log.Printf("Add Attendances: Error binding data: %v", errBind)
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "error binding data: "+errBind.Error(), nil))
	}

	dataAtt := attendance.Attandance{
		PersonalDataID: uint(userID),
		Clock_out:      AttRequest.Clock_out,
		Status:         AttRequest.Status,
		Long:           AttRequest.Long,
		Lat:            AttRequest.Lat,
		Notes:          AttRequest.Notes,
	}

	// panggil fungsi updateAtt pada service
	err = ah.srv.UpdateAtt(uint(attId), dataAtt)
	if err != nil {
		log.Printf("Update Attandance: Error update Attendance: %v", err)
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "error update data: "+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "clock out successfully", nil))
}

// Delete Att by id attendance
func (ah *AttHandler) DeleteAttendance(c echo.Context) error {
	attID := c.Param("attendance_id")

	attId, err := strconv.ParseUint(attID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "failed", "Invalid Attendance ID", nil))
	}
	userID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "unauthorized", nil))
	}
	err = ah.srv.DeleteAttByID(uint(attId))
	if err != nil {
		log.Printf("Delete Attandance: Error delete Attendance: %v", err)
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "error delete data: "+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "Delete successfully", nil))
}

// Get by personal ID
func (ah *AttHandler) GetAttendancesHandler(c echo.Context) error {
	attID := c.Param("attendance_id")

	attId, err := strconv.ParseUint(attID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "failed", "Invalid Attendance ID", nil))
	}
	personalID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if personalID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "unauthorized", nil))
	}

	// Membaca parameter dari query string
	pageStr := c.QueryParam("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSizeStr := c.QueryParam("pageSize")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	// Create Pagination object
	pagination := utils.NewPagination(page, pageSize)

	// Use Pagination object to get offset and limit
	offset := pagination.Offset()
	limit := pagination.PageSize

	// Call the service to retrieve the records
	attendances, err := ah.srv.GetAttByPersonalID(uint(attId), limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	totalItems, _ := ah.srv.CountAllAtt()
	var response []AttResponse
	for _, att := range attendances {
		response = append(response, ToGetAttendanceResponse(att))
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

	meta := map[string]interface{}{
		"totalItems":   totalItems,
		"itemsPerPage": limit,
		"currentPage":  page,
		"totalPages":   totalPages,
	}
	// Return the retrieved records as JSON
	return c.JSON(http.StatusOK, responses.PaginatedJSONResponse(http.StatusOK, "success", "attendance records retrieved successfully", response, meta))
}

func (ah *AttHandler) GetAllAttendancesHandler(c echo.Context) error {
	personalID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if personalID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "unauthorized", nil))
	}
	fmt.Println("user_id:", personalID)

	// Membaca parameter dari query string
	pageStr := c.QueryParam("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	pageSizeStr := c.QueryParam("pageSize")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}
	var attendances []attendance.Attandance
	// Create Pagination object
	pagination := utils.NewPagination(page, pageSize)

	// Use Pagination object to get offset and limit
	offset := pagination.Offset()
	limit := pagination.PageSize

	// filter by date
	filterDate := c.QueryParam("date")
	if filterDate != "" {
		attendances, err = ah.srv.GetAllAttbyDate(filterDate, limit, offset)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	} else {
		attendances, err = ah.srv.GetAllAtt(limit, offset)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}
	totalItems, _ := ah.srv.CountAllAtt()
	var response []AttAllResponse
	for _, att := range attendances {
		response = append(response, ToGetAllAttendance(att))
	}
	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

	meta := map[string]interface{}{
		"totalItems":   totalItems,
		"itemsPerPage": limit,
		"currentPage":  page,
		"totalPages":   totalPages,
	}
	// Return the retrieved records as JSON
	return c.JSON(http.StatusOK, responses.PaginatedJSONResponse(http.StatusOK, "success", "attendance records retrieved successfully", response, meta))

}

func (ah *AttHandler) DownloadPdf(c echo.Context) error {
	personalID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if personalID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "unauthorized", nil))
	}
	err := ah.srv.DownloadAllAtt()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "Download failed", nil))
	}
	return c.File("./Attendance.pdf")
}
