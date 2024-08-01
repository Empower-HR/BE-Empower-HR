package handler

import (
	"be-empower-hr/app/middlewares"
	leaves "be-empower-hr/features/Leaves"
	"be-empower-hr/utils/responses"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type LeavesHandler struct {
	leavesService leaves.ServiceLeavesInterface
}

func New(le leaves.ServiceLeavesInterface) *LeavesHandler {
	return &LeavesHandler{
		leavesService: le,
	}
}

func (lh *LeavesHandler) RequestLeave(c echo.Context) error {
	var leaveRequest LeaveRequest
	if err := c.Bind(&leaveRequest); err != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "failed", "Invalid request payload", nil))
	}

	userID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "Invalid token", nil))
	}

	leave := leaves.LeavesDataEntity{
		StartDate:      leaveRequest.StartDate,
		EndDate:        leaveRequest.EndDate,
		Reason:         leaveRequest.Reason,
		PersonalDataID: uint(userID),
	}

	if err := lh.leavesService.RequestLeave(uint(userID), leave); err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "error", "Failed to request leave", nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "Leave requested successfully", nil))
}

func (lh *LeavesHandler) UpdateLeaveStatus(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "error converting id: " + errConv.Error(),
		})
	}

	updatedLeaves := LeaveRequest{}
	if errBind := c.Bind(&updatedLeaves); errBind != nil {
		log.Printf("update leave status: Error binding data: %v", errBind)
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "error binding data: "+errBind.Error(), nil))
	}
	userID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "Invalid token", nil))
	}

	leaveUpdate := leaves.LeavesDataEntity{
		Status: updatedLeaves.Status,
		Reason: updatedLeaves.Reason,
	}

	err := lh.leavesService.UpdateLeaveStatus(uint(userID), uint(idConv), leaveUpdate)
	if err != nil {
		log.Printf("update leave status: Error updating status: %v", err)
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "failed", "error updating status: "+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "leave status updated successfully", nil))
}

func (lh *LeavesHandler) GetLeavesByID(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "Invalid leave ID", nil))
	}

	leaveEntity, err := lh.leavesService.GetLeavesByID(uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "error", "Failed to get leave details", nil))
	}

	leaveHistory := LeaveHistory{
		EmployeeID:   leaveEntity.PersonalDataID,
		PersonalName: leaveEntity.Name,
		JobPosition:  leaveEntity.JobPosition,
		StartDate:    leaveEntity.StartDate,
		EndDate:      leaveEntity.EndDate,
		Reason:       leaveEntity.Reason,
		Status:       leaveEntity.Status,
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "Leave details retrieved successfully", leaveHistory))
}

func (lh *LeavesHandler) ViewLeaveHistory(c echo.Context) error {
	userID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "Invalid token", nil))
	}

	companyID, err := middlewares.NewMiddlewares().ExtractCompanyID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "Invalid company ID", nil))
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1
	}
	pageSize, err := strconv.Atoi(c.QueryParam("pageSize"))
	if err != nil {
		pageSize = 10
	}
	status := c.QueryParam("status")
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")

	// Get leave history
	leaveEntities, err := lh.leavesService.ViewLeaveHistory(companyID, uint(userID), page, pageSize, status, startDate, endDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "error", "Failed to fetch leave history", nil))
	}

	// Get dashboard stats
	dashboardStats, err := lh.leavesService.Dashboard(uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "error", "Failed to fetch dashboard stats", nil))
	}

	var leaveHistories []LeaveHistory
	for _, leave := range leaveEntities {
		leaveHistories = append(leaveHistories, LeaveHistory{
			LeaveID:      leave.LeavesID,
			EmployeeID:   leave.PersonalDataID,
			PersonalName: leave.Name,
			JobPosition:  leave.JobPosition,
			StartDate:    leave.StartDate,
			EndDate:      leave.EndDate,
			Reason:       leave.Reason,
			Status:       leave.Status,
		})
	}

	response := LeaveHistoryResponse{
		Code:           http.StatusOK,
		Status:         "success",
		Message:        "Leave applications retrieved successfully",
		Data:           leaveHistories,
		TotalEmployees: int(dashboardStats.TotalUsers),
		TotalLeaves:    int(dashboardStats.LeavesPending),
	}

	return c.JSON(http.StatusOK, response)
}

func (lh *LeavesHandler) ViewLeaveHistoryEmployee(c echo.Context) error {
	userID := middlewares.NewMiddlewares().ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.JSONWebResponse(http.StatusUnauthorized, "failed", "Invalid token", nil))
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1
	}
	pageSize, err := strconv.Atoi(c.QueryParam("pageSize"))
	if err != nil {
		pageSize = 10
	}
	status := c.QueryParam("status")
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")

	leaveEntities, err := lh.leavesService.ViewLeaveHistoryEmployee(uint(userID), page, pageSize, status, startDate, endDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "error", "Failed to fetch leave history for employee", nil))
	}

	var leaveHistories []LeaveHistory
	for _, leave := range leaveEntities {
		leaveHistories = append(leaveHistories, LeaveHistory{
			LeaveID:      leave.LeavesID,
			EmployeeID:   leave.PersonalDataID,
			PersonalName: leave.Name,
			JobPosition:  leave.JobPosition,
			StartDate:    leave.StartDate,
			EndDate:      leave.EndDate,
			Reason:       leave.Reason,
			Status:       leave.Status,
		})
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "Leave history for employee retrieved successfully", leaveHistories))
}
