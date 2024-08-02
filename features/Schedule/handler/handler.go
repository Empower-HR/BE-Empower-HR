package handler

import (
	schedule "be-empower-hr/features/Schedule"
	"be-empower-hr/utils"
	"be-empower-hr/utils/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ScheduleHandler struct {
	scheduleService schedule.ServiceScheduleInterface
}

func New(sc schedule.ServiceScheduleInterface) *ScheduleHandler {
	return &ScheduleHandler{
		scheduleService: sc,
	}
}

func (sh *ScheduleHandler) CreateSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req ScheduleRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
		}

		parsedTime, err := utils.StringToDate(req.EffectiveDate)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input effective date"})
		}
		scheduleData := schedule.ScheduleDataEntity{
			CompanyID:     req.Company,
			Name:          req.Name,
			EffectiveDate: parsedTime,
			ScheduleIn:    req.ScheduleIn,
			ScheduleOut:   req.ScheduleOut,
			BreakStart:    req.BreakStart,
			BreakEnd:      req.BreakEnd,
			Days:          req.Days,
			Description:   req.Description,
		}

		id, err := sh.scheduleService.CreateSchedule(scheduleData)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "error", err.Error(), nil))
		}

		schedule, err := sh.scheduleService.GetScheduleById(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, responses.JSONWebResponse(http.StatusNotFound, "error", "Schedule not found", nil))
		}

		return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "Schedule created successfully", schedule))
	}
}

func (sh *ScheduleHandler) GetAllSchedule(c echo.Context) error {
	schedules, err := sh.scheduleService.GetAllSchedule()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "error", err.Error(), nil))
	}
	var responseDetail []ScheduleDataResponse
	for _, detail := range schedules {
		days := detail.Days
		responseDetail = append(responseDetail, ToGetAllSchedule(detail, days))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "Schedules retrieved successfully", responseDetail))
}

// GetScheduleById retrieves a schedule by ID.
func (sh *ScheduleHandler) GetScheduleById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "Invalid schedule ID", nil))
	}

	schedule, err := sh.scheduleService.GetScheduleById(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, responses.JSONWebResponse(http.StatusNotFound, "error", "Schedule not found", nil))
	}
	var responseDetail []ScheduleDataResponse

	days := schedule.Days
	responseDetail = append(responseDetail, ToGetAllSchedule(*schedule, days))

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "Schedule retrieved successfully", responseDetail))
}

// UpdateSchedule updates an existing schedule.
func (sh *ScheduleHandler) UpdateSchedule(c echo.Context) error {
	var req ScheduleRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "Invalid input", nil))
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "Invalid schedule ID", nil))
	}

	parsedTime, err := utils.StringToDate(req.EffectiveDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input effective date"})
	}

	// Convert ScheduleRequest to ScheduleDataEntity
	scheduleEntity := schedule.ScheduleDataEntity{
		CompanyID:     req.Company,
		Name:          req.Name,
		EffectiveDate: parsedTime,
		ScheduleIn:    req.ScheduleIn,
		ScheduleOut:   req.ScheduleOut,
		BreakStart:    req.BreakStart,
		BreakEnd:      req.BreakEnd,
		Days:          req.Days,
		Description:   req.Description,
	}

	err = sh.scheduleService.UpdateSchedule(uint(id), scheduleEntity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "error", err.Error(), nil))
	}

	schedule, err := sh.scheduleService.GetScheduleById(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, responses.JSONWebResponse(http.StatusNotFound, "error", "Schedule not found", nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "Schedule updated successfully", schedule))
}

// DeleteSchedule deletes a schedule by ID.
func (sh *ScheduleHandler) DeleteSchedule(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "Invalid schedule ID", nil))
	}

	err = sh.scheduleService.DeleteSchedule(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "Schedule deleted successfully", nil))
}
