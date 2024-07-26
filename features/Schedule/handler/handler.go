package handler

import (
	schedule "be-empower-hr/features/Schedule"
	"be-empower-hr/utils/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type ScheduleHandler struct {
	scheduleService schedule.ServiceScheduleInterface
}

func New(sc schedule.ServiceScheduleInterface) *ScheduleHandler {
	return &ScheduleHandler{
		scheduleService: sc,
	}
}

func (sc *ScheduleHandler) CreateSchedule(c echo.Context) error {
	var req ScheduleRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	scheduleData := schedule.ScheduleDataEntity{
		Name:          req.Name,
		EffectiveDate: req.EffectiveDate,
		ScheduleIn:    req.ScheduleIn,
		ScheduleOut:   req.ScheduleOut,
		BreakStart:    req.BreakStart,
		BreakEnd:      req.BreakEnd,
		Days:          req.Days,
		Description:   req.Description,
	}

	id, err := sc.scheduleService.CreateSchedule(scheduleData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "Schedule created successfully", map[string]uint{"id": id}))
}

func (sh *ScheduleHandler) GetAllSchedule(c echo.Context) error {
	schedules, err := sh.scheduleService.GetAllSchedule()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "error", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "Schedules retrieved successfully", schedules))
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

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "Schedule retrieved successfully", schedule))
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

	// Convert ScheduleRequest to ScheduleDataEntity
	scheduleEntity := schedule.ScheduleDataEntity{
		Name:          req.Name,
		EffectiveDate: req.EffectiveDate,
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

	return c.JSON(http.StatusOK, responses.JSONWebResponse(http.StatusOK, "success", "Schedule updated successfully", nil))
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
