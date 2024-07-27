package handler

import (
	companies "be-empower-hr/features/Companies"
	schedule "be-empower-hr/features/Schedule"
	"be-empower-hr/utils"
	"be-empower-hr/utils/responses"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ScheduleHandler struct {
	scheduleService  schedule.ServiceScheduleInterface
	companiesService companies.Service
}

func New(sc schedule.ServiceScheduleInterface, cs companies.Service) *ScheduleHandler {
	return &ScheduleHandler{
		scheduleService:  sc,
		companiesService: cs,
	}
}

func (sh *ScheduleHandler) CreateSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req ScheduleRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
		}

		company, err := sh.companiesService.GetCompany(req.Company)
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusNotFound, responses.JSONWebResponse(http.StatusNotFound, "not found", "Company id not found", nil))
		}

		parsedTime, err := utils.StringToDate(req.EffectiveDate)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input effective date"})
		}
		scheduleData := schedule.ScheduleDataEntity{
			CompanyID:     company.ID,
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

	company, err := sh.companiesService.GetCompany(req.Company)
	if err != nil {
		log.Print("Error", err.Error())
		return c.JSON(http.StatusNotFound, responses.JSONWebResponse(http.StatusNotFound, "not found", "Company id not found", nil))
	}

	parsedTime, err := utils.StringToDate(req.EffectiveDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input effective date"})
	}

	// Convert ScheduleRequest to ScheduleDataEntity
	scheduleEntity := schedule.ScheduleDataEntity{
		CompanyID:     company.ID,
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
