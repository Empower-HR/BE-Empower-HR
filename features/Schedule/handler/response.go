package handler

import (
	schedule "be-empower-hr/features/Schedule"
	"time"
)

type ScheduleDataResponse struct {
	ID    		  uint      `json:"id"`
	CompanyID     uint      `json:"company_id"`
	Name          string    `json:"name"`
	EffectiveDate time.Time `json:"effective_date"`
	ScheduleIn    string    `json:"schedule_in"`
	ScheduleOut   string    `json:"schedule_out"`
	BreakStart    string    `json:"break_start"`
	BreakEnd      string    `json:"break_end"`
	Days          string       `json:"repeat_until"`
	Description   string    `json:"description"`
}


func ToGetAllSchedule(schedule schedule.ScheduleDataEntity, days string) ScheduleDataResponse {
	days = days + " days"
	return ScheduleDataResponse{
		ID:       schedule.ScheduleId,
		CompanyID: schedule.CompanyID,
		Name: schedule.Name,
		EffectiveDate: schedule.EffectiveDate,
		ScheduleIn: schedule.ScheduleIn,
		ScheduleOut: schedule.ScheduleOut,
		BreakStart: schedule.BreakStart,
		BreakEnd: schedule.BreakEnd,
		Days: days,
		Description: schedule.Description,
	}
}