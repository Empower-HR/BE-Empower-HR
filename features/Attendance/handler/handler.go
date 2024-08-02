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
	"time"

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
	AttRequest := AttRequest{}
	if errBind := c.Bind(&AttRequest); errBind != nil {
		log.Printf("Add Attendances: Error binding data: %v", errBind)
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "error binding data: "+errBind.Error(), nil))
	}
	format := "02-01-2006"
	date, err := time.Parse(format, AttRequest.Date)
    if err != nil {
        fmt.Println("Error parsing date:", err)
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "error binding data: "+err.Error(), nil))
    }

	dataAtt := attendance.Attandance{
		PersonalDataID: uint(userID),
		Clock_in:       AttRequest.Clock_in,
		Date:           date,
		Long:           AttRequest.Long,
		Lat:            AttRequest.Lat,
		Notes:          AttRequest.Notes,
	}

	// panggil fungsi addAtt pada service
	err = ah.srv.AddAtt(dataAtt)
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
	attID := c.Param("employee_id")

	attId, err := strconv.ParseUint(attID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "failed", "Invalid Emplpyee ID", nil))
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
	status := c.QueryParam("status")
	searchBox:= c.QueryParam("search")
	filterDate:= c.QueryParam("date")

	
	// Create Pagination object
	pagination := utils.NewPagination(page, pageSize)

	// Use Pagination object to get offset and limit
	offset := pagination.Offset()
	limit := pagination.PageSize

	var attDetail []attendance.AttendanceDetail
	var responseDetail []AttResponse
	var result interface{}
	var totalItems int64

	if status != "" { 
		attDetail, err = ah.srv.GetAttByPersonalIDandStatus(uint(attId), status, limit, offset)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		for _, att := range attDetail {
            responseDetail = append(responseDetail, ToGetAttendanceResponse(att))
        }
		totalItems, err = ah.srv.CountAllAttbyStatusandPerson(status, uint(attId))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		result = responseDetail
	} else if filterDate != "" {
		date, err := strconv.Atoi(filterDate)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		err = ah.srv.CheckingTheValueOfDate(date)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		attDetail, err = ah.srv.GetAllAttbyDateandPerson(int(date), limit, offset, uint(attId))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		for _, att := range attDetail {
            responseDetail = append(responseDetail, ToGetAttendanceResponse(att))
        }
		result = responseDetail
		totalItems, err = ah.srv.CountAllAttbyDateandPerson(int(date), uint(attId))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	} else{
		attDetail, err = ah.srv.GetAttByPersonalID(uint(attId), searchBox, limit, offset)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		for _, detail := range attDetail {
			responseDetail = append(responseDetail, ToGetAttendanceResponse(detail))
		}
		result = responseDetail
	}
	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

		meta := map[string]interface{}{
			"totalItems":   totalItems,
			"itemsPerPage": limit,
			"currentPage":  page,
			"totalPages":   totalPages,
		}

		
	// Return the retrieved records as JSON
	return c.JSON(http.StatusOK, responses.PaginatedJSONResponse(http.StatusOK, "success", "attendance records retrieved successfully", result, meta))
}

func (ah *AttHandler) GetAllAttendancesHandler(c echo.Context) error {
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
	var attDetail []attendance.AttendanceDetail
	// Create Pagination object
	pagination := utils.NewPagination(page, pageSize)

	// Use Pagination object to get offset and limit
	offset := pagination.Offset()
	limit := pagination.PageSize

	// filter by date
	filterDate := c.QueryParam("date")
	filterStatus := c.QueryParam("status")
	searchBox := c.QueryParam("search")
	var responseDetail []AttResponse
	var result interface{}
	var totalItems int64

	if filterDate != "" {
		date, err := strconv.Atoi(filterDate)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		err = ah.srv.CheckingTheValueOfDate(date)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		attDetail, err = ah.srv.GetAllAttbyDate(int(date), limit, offset)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		for _, att := range attDetail {
            responseDetail = append(responseDetail, ToGetAttendanceResponse(att))
        }
		result = responseDetail
		totalItems, err = ah.srv.CountAllAttbyDate(int(date))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	} else if filterStatus != "" {
		attDetail, err = ah.srv.GetAllAttbyStatus(filterStatus, limit, offset)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		for _, att := range attDetail {
            responseDetail = append(responseDetail, ToGetAttendanceResponse(att))
        }
		result = responseDetail
		totalItems, err = ah.srv.CountAllAttbyStatus(filterStatus)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}else{
		attDetail, err = ah.srv.GetAllAtt(searchBox, limit, offset)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		for _, detail := range attDetail {
			responseDetail = append(responseDetail, ToGetAttendanceResponse(detail))
		}
		result = responseDetail
		totalItems, err = ah.srv.CountAllAtt()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}
	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

		meta := map[string]interface{}{
			"totalItems":   totalItems,
			"itemsPerPage": limit,
			"currentPage":  page,
			"totalPages":   totalPages,
		}

		
	// Return the retrieved records as JSON
	return c.JSON(http.StatusOK, responses.PaginatedJSONResponse(http.StatusOK, "success", "attendance records retrieved successfully", result, meta))
}

func (ah *AttHandler) GetAttendancesbyID(c echo.Context) error {
	attID := c.Param("attendance_id")

	attId, err := strconv.ParseUint(attID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "failed", "Invalid Attendance ID", nil))
	}
	userID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "unauthorized", nil))
	}
	att, err := ah.srv.GetAttByIdAtt(uint(attId))
	if err != nil {
		log.Printf("Get Attandance: Error Get Attendance: %v", err)
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "error Get data: "+err.Error(), nil))
	}
	var responseDetail []AttResponse
	for _, detail := range att {
		responseDetail = append(responseDetail, ToGetAttendanceResponse(detail))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "Get Attendance successfully", responseDetail))
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
